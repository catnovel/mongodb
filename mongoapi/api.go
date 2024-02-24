package mongoapi

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client     *mongo.Client
	Database   string
	Collection string
}

func (d *DB) getCTX() context.Context {
	return context.Background()
}

func (d *DB) SetDatabase(database string) *DB {
	d.Database = database
	return d
}
func (d *DB) SetCollection(collection string) *DB {
	d.Collection = collection
	return d
}

func (d *DB) GetCollection() *mongo.Collection {
	dbs := d.GetDatabase()
	if dbs == nil || d.Collection == "" {
		return nil
	}
	return dbs.Collection(d.Collection)
}
func (d *DB) GetDatabase() *mongo.Database {
	if d.Database == "" {
		return nil
	}
	return d.Client.Database(d.Database)
}
func (d *DB) getEmptyError() error {
	return fmt.Errorf(fmt.Sprintf("database:%s collection:%s is nil", d.Database, d.Collection))
}
func (d *DB) InsertOne(doc interface{}) error {
	if col := d.GetCollection(); col != nil {
		insertOneId, err := col.InsertOne(d.getCTX(), doc)
		if err != nil {
			return err
		}
		if insertOneId.InsertedID == nil {
			return fmt.Errorf("inserted id is nil")
		}
		return err

	}
	return d.getEmptyError()
}

// FindOne 查询单个文档
func (d *DB) FindOne(filter interface{}) *mongo.SingleResult {
	if col := d.GetCollection(); col != nil {
		return col.FindOne(d.getCTX(), filter)
	}
	return nil
}
func (d *DB) FindOneAndResult(filter interface{}, result interface{}) error {
	return d.FindOne(filter).Decode(result)
}

func (d *DB) Find(filter interface{}) (*mongo.Cursor, error) {
	if col := d.GetCollection(); col != nil {
		return col.Find(d.getCTX(), filter)
	}
	return nil, d.getEmptyError()
}
func (d *DB) FindAndResult(filter interface{}, result interface{}) error {
	cur, err := d.Find(filter)
	if err != nil {
		return err
	}
	return cur.All(d.getCTX(), result)
}

func (d *DB) UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {
	if col := d.GetCollection(); col != nil {
		return col.UpdateOne(d.getCTX(), filter, update)
	}
	return nil, d.getEmptyError()
}

func (d *DB) UpdateMany(filter, update interface{}) (*mongo.UpdateResult, error) {
	if col := d.GetCollection(); col != nil {
		return col.UpdateMany(d.getCTX(), filter, update)
	}
	return nil, d.getEmptyError()
}

func (d *DB) UpdateAndInsertOne(filter, update interface{}) error {
	if col := d.GetCollection(); col != nil {
		_, err := col.UpdateOne(d.getCTX(), filter, update, options.Update().SetUpsert(true))
		return err
	}
	return d.getEmptyError()
}
func (d *DB) DeleteMany(filter interface{}) error {
	if col := d.GetCollection(); col != nil {
		_, err := col.DeleteMany(d.getCTX(), filter)
		return err
	}
	return d.getEmptyError()
}

func (d *DB) DeleteOne(filter interface{}) error {
	_, err := d.Client.Database(d.Database).Collection(d.Collection).DeleteOne(d.getCTX(), filter)
	return err
}

func (d *DB) CreateIndex(unique bool, keys map[string]interface{}) error {
	if col := d.GetCollection(); col != nil {
		_, err := col.Indexes().CreateOne(d.getCTX(), mongo.IndexModel{Keys: keys, Options: options.Index().SetUnique(unique)})
		return err
	}
	return d.getEmptyError()
}
func (d *DB) CreateManyIndex(models []map[string]interface{}) []error {
	var errs []error
	for _, model := range models {
		if col := d.GetCollection(); col != nil {
			_, err := col.Indexes().CreateOne(d.getCTX(), mongo.IndexModel{Keys: model})
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, d.getEmptyError())
		}
	}
	return errs
}
