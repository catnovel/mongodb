package mongodb

import (
	"context"
	"github.com/catnovel/mongodb/mongoapi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(opt ...ClientOptionFunc) (*Client, error) {
	defaultOpt := []ClientOptionFunc{
		WithTimeoutSecond(30),
		WithPoolMaxSize(200),
		WithDefaultContext(context.Background()),
		WithURI("mongodb://localhost:27017"),
	}
	return newMongoDBClient(append(defaultOpt, opt...)...)
}

func newMongoDBClient(opt ...ClientOptionFunc) (*Client, error) {
	client := &Client{}
	for _, o := range opt {
		o.Apply(client)
	}
	var err error
	clientOptions := options.Client().ApplyURI(client.uri)
	if client.poolMaxSize > 0 {
		clientOptions.SetMaxPoolSize(uint64(client.poolMaxSize))
	}
	if client.HTTPClient != nil {
		clientOptions.SetHTTPClient(client.HTTPClient)
	}
	if client.credential != nil {
		clientOptions.SetAuth(*client.credential)
	}
	if client.timeout > 0 {
		clientOptions.SetTimeout(client.timeout)
	}
	client.Client, err = mongo.Connect(client.defaultContext, clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Client.Ping(client.defaultContext, nil); err != nil {
		return nil, err
	}
	return client, nil
}

// Disconnect 关闭MongoDB客户端连接
func (m *Client) Disconnect() error {
	return m.Client.Disconnect(m.defaultContext)
}

func (m *Client) NewDB() *mongoapi.DB {
	db := &mongoapi.DB{Client: m.Client, Database: m.defaultDBName, Collection: m.defaultCollectionName}
	if m.defaultDBName != "" {
		db.Database = m.defaultDBName
	}
	if m.defaultCollectionName != "" {
		db.Collection = m.defaultCollectionName
	}
	return db
}
