//Package redis is a simple encapsulation of gopkg.in/redis.v3.
package redis

import (
	"strings"

	"gopkg.in/redis.v3"
)

// GetNamespace returns a string from args separated by colons.
func GetNamespace(args ...string) string {
	return strings.Join(args, ":")
}

type Client struct {
	*redis.Client
}

// NewRedisClient returns a redis Client instance with specified redis server address.
func NewRedisClient(addr string) *Client {
	return &Client{
		redis.NewClient(
			&redis.Options{Addr: addr},
		),
	}
}

// IsRedisOnline detects if redis server is online.
//
// If the redis server is not online, it will return a false and an error.
func (redis *Client) IsRedisOnline() (bool, error) {
	err := redis.Ping().Err()
	if err != nil {
		return false, err
	}

	return true, nil
}
