package storage

import (
	"encoding/json"
	"github.com/roessland/miraiex-dca/miraiex-dca/models"
	"github.com/rs/xid"
	"go.etcd.io/bbolt"
	"log"
	"time"
)

var ordersBucketName = []byte("orders")

type Repo struct {
	db *bbolt.DB
}

func NewRepo(dbPath string) (*Repo, error) {
	var repo Repo

	db, err := bbolt.Open(dbPath, 0666, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	repo.db = db

	// Create buckets
	repo.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("orders"))
		return err
	})

	return &repo, nil
}

func (repo *Repo) Close() {
	_ = repo.db.Close()
}

func (repo *Repo) CreateOrder(o *models.Order) (string, error) {
	var orderId string
	err := repo.db.Update(func(tx *bbolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket(ordersBucketName)
		if b == nil {
			panic("bucket does not exist")
		}

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id := xid.New()
		o.ID = id.String()

		// Marshal user data into bytes.
		buf, err := json.Marshal(o)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		orderId = id.String()
		return b.Put(id[:], buf)
	})
	return orderId, err
}

func (repo *Repo) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	err := repo.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(ordersBucketName)
		if b == nil {
			panic("bucket does not exist")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var order models.Order
			err := json.Unmarshal(v, &order)
			if err != nil {
				log.Print("Cannot unmarshal as order:", string(v))
			}
			orders = append(orders, order)
		}
		return nil
	})
	return orders, err
}
