package mongoapi

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func (d *DB) SetDatabase(database string) *DB {
	d.Database = database
	return d

}
func (d *DB) SetCollection(collection string) *DB {
	d.Collection = collection
	return d
}

// InsertOne 插入一个文档
func (d *DB) insertOneResult(doc interface{}) (*mongo.InsertOneResult, error) {
	return d.Client.Database(d.Database).Collection(d.Collection).InsertOne(context.Background(), doc)
}
func (d *DB) InsertOne(doc interface{}) error {
	_, err := d.insertOneResult(doc)
	return err
}

// FindOne 查询单个文档
func (d *DB) FindOne(filter interface{}) *mongo.SingleResult {
	return d.Client.Database(d.Database).Collection(d.Collection).FindOne(context.Background(), filter)
}

func (d *DB) Find(filter interface{}) (*mongo.Cursor, error) {
	return d.Client.Database(d.Database).Collection(d.Collection).Find(context.Background(), filter)
}
func (d *DB) UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {
	return d.Client.Database(d.Database).Collection(d.Collection).UpdateOne(context.Background(), filter, update)
}

func (d *DB) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return d.Client.Database(d.Database).Collection(d.Collection).DeleteOne(context.Background(), filter)
}
