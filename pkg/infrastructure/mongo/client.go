package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	eq           = "$eq"
	set          = "$set"
)

type Client struct {
	mc *mongo.Client
}

func NewClient(conn string) *Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	return &Client{
		mc: client,
	}
}

func (c *Client) InsertOne(db string, coll string, v interface{}) error {
	col := c.mc.Database(db).Collection(coll)
	f, err := bson.Marshal(v)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, err = col.InsertOne(context.Background(), f); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c *Client) Update(db string, coll string, f bson.D, v interface{}) error {
	col := c.mc.Database(db).Collection(coll)
	u := bson.D{{set, v}}
	var updatedDocument bson.M
	err := col.FindOneAndUpdate(context.Background(), f, u).Decode(&updatedDocument)
	fmt.Println(err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (c *Client) FindOne(db string, coll string, q bson.M, v interface{}) error {
	col := c.mc.Database(db).Collection(coll)
	if err := col.FindOne(context.Background(), q).Decode(v); err != nil {
		return err
	}
	return nil
}

func (c *Client) All(db string, coll string, f bson.M) ([]bson.M, error) {
	col := c.mc.Database(db).Collection(coll)
	cursor, err := col.Find(context.Background(), f, nil)
	if err != nil {
		return nil, err
	}
	var list []bson.M
	if err = cursor.All(context.Background(), &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *Client) Latest(db string, coll string) (bson.M, error) {
	col := c.mc.Database(db).Collection(coll)
	opt := options.Find()
	opt.SetSort(bson.D{{"_id", -1}})
	sortCursor, err := col.Find(context.Background(), bson.D{}, opt)
	if err != nil {
		return nil, err
	}
	var itemsSorted []bson.M
	if err = sortCursor.All(context.Background(), &itemsSorted); err != nil {
		return nil, err
	}
	item := itemsSorted[0]
	return item, nil
}

func (c *Client) Count(db string, coll string, f bson.M) (int64, error) {
	col := c.mc.Database(db).Collection(coll)
	count, err := col.CountDocuments(context.Background(), f)
	if err != nil {
		return 0, err
	}
	return count, nil
}
