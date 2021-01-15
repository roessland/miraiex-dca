package miraiex

import (
	"fmt"
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
	ApiKey    string
	ClientID  string
	SecretKey string
}

// NewClient creates a new client.
func NewClient() *Client {
	return &Client{}
}