package miraiex

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

const ApiBase = "https://api.miraiex.com"

// BuildURL creates an API URL.
func BuildURL(path string, args ...interface{}) string {
	if args == nil {
		return ApiBase + path
	}
	return ApiBase + fmt.Sprintf(path, args...)
}

// Client interacts with MiraiEx.
type Client struct {
	apiKey    string
	clientID  string
	secretKey string
	logger *logrus.Logger
}

// NewClient creates a new client.
func NewClient(logger *logrus.Logger) *Client {
	c := &Client{}
	c.logger = logger
	return c
}

// SetAuthentication sets access tokens for MiraiEx, allowing use of private APIs
func (c *Client) SetAuthentication(apiKey, clientID, secretKey string) *Client {
	c.apiKey = apiKey
	c.clientID = clientID
	c.secretKey = secretKey
	return c
}