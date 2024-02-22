package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"time"
)

// Client 定义一个MongoDB客户端结构体
type Client struct {
	uri                   string
	poolMaxSize           int
	defaultContext        context.Context
	debug                 bool
	Client                *mongo.Client
	defaultDBName         string
	HTTPClient            *http.Client
	defaultCollectionName string
	timeout               time.Duration
	credential            *options.Credential
}
type ClientOptionFunc func(client *Client)

func (f ClientOptionFunc) Apply(client *Client) {
	f(client)
}

func WithURI(uri string) ClientOptionFunc {
	return func(client *Client) {
		client.uri = uri
	}
}
func WithPoolMaxSize(poolMaxSize int) ClientOptionFunc {
	return func(client *Client) {
		client.poolMaxSize = poolMaxSize
	}
}
func WithClient(client *mongo.Client) ClientOptionFunc {
	return func(c *Client) {
		c.Client = client
	}

}
func WithDatabase(database string) ClientOptionFunc {
	return func(c *Client) {
		c.defaultDBName = database
	}
}
func WithCollection(collection string) ClientOptionFunc {
	return func(c *Client) {
		c.defaultCollectionName = collection
	}
}
func WithDefaultContext(ctx context.Context) ClientOptionFunc {
	return func(c *Client) {
		c.defaultContext = ctx
	}
}
func WithDebug() ClientOptionFunc {
	return func(c *Client) {
		c.debug = true
	}
}
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}
func WithCredential(credential *options.Credential) ClientOptionFunc {
	return func(c *Client) {
		c.credential = credential
	}
}
func WithTimeoutSecond(timeout int) ClientOptionFunc {
	return func(c *Client) {
		c.timeout = time.Duration(timeout) * time.Second
	}
}
