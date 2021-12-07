package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDBWorker struct {
	cntStr string
}

func createMongoDBWorker(connectionString string) *mongoDBWorker {
	return &mongoDBWorker{cntStr: connectionString}
}

func (w *mongoDBWorker) ping() error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}

func (w *mongoDBWorker) ps(dbNameFilter string) (databaseCollectionInfo, error) {
	info := make(databaseCollectionInfo)

	client, err := w.connect()
	if err != nil {
		return info, err
	}

	dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		return info, err
	}

	for _, db := range dbs {

		if dbNameFilter != "" && db != dbNameFilter {
			continue
		}

		colls, err := client.Database(db).ListCollectionNames(context.TODO(), bson.D{})
		if err != nil {
			fmt.Printf("skip database %s :%s", db, err)
			continue
		}
		info[db] = colls
	}
	return info, nil
}

func (w *mongoDBWorker) connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(w.cntStr))
}
