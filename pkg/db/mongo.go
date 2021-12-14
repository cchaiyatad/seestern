package db

import (
	"context"
	"fmt"
	"time"

	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/cf"
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
	info, err := w.getDatabaseCollectionInfo()
	if err != nil || dbNameFilter == "" {
		return info, err
	}

	specificDBInfo := make(databaseCollectionInfo)
	if colls, ok := info[dbNameFilter]; ok {
		specificDBInfo[dbNameFilter] = colls
	}

	return specificDBInfo, nil
}

func (w *mongoDBWorker) initConfigFile(param *InitParam, configGenerator *cf.ConfigFileGenerator) error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	infos, err := w.getDatabaseCollectionInfoWithClient(client)
	if err != nil {
		return err
	}

	toGenColls := parseCollectionInputFromArgs(param.TargetColls)

	for db, colls := range toGenColls {
		for _, coll := range colls {
			if _, ok := infos[db][coll]; !ok {
				log.Logf(log.Warning, "%s\n", &ErrSkipCreateConfigfile{db, coll, "not exist"})
				continue
			}
			cursor, err := w.getCursor(client, db, coll)
			if err != nil {
				log.Logf(log.Warning, "%s\n", &ErrSkipCreateConfigfile{db, coll, err.Error()})
				continue
			}
			log.Logf(log.Info, "generate: database %s collection %s\n", db, coll)

			callBack, onFinish := configGenerator.Begin(db, coll)
			go w.iterateByCursor(cursor, db, coll, callBack, onFinish)
		}
	}
	return nil
}

func (w *mongoDBWorker) insert() {
	fmt.Println("insert")
	panic("Not implemented")
}

func (w *mongoDBWorker) drop() {
	fmt.Println("drop")
	panic("Not implemented")
}

func (w *mongoDBWorker) connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, options.Client().ApplyURI(w.cntStr))
}

func (w *mongoDBWorker) getDatabaseCollectionInfo() (databaseCollectionInfo, error) {
	client, err := w.connect()
	if err != nil {
		return nil, err
	}

	return w.getDatabaseCollectionInfoWithClient(client)
}

func (*mongoDBWorker) getDatabaseCollectionInfoWithClient(client *mongo.Client) (databaseCollectionInfo, error) {
	info := make(databaseCollectionInfo)

	dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	for _, db := range dbs {
		colls, err := client.Database(db).ListCollectionNames(context.TODO(), bson.D{})
		if err != nil {
			log.Logf(log.Warning, "skip database %s :%s\n", db, err)
			continue
		}

		info[db] = make(map[string]struct{})
		for _, coll := range colls {
			info[db][coll] = struct{}{}
		}
	}
	return info, nil
}

func (*mongoDBWorker) getCursor(client *mongo.Client, dbName string, collName string) (*mongo.Cursor, error) {
	if client == nil {
		return nil, ErrClientIsNil
	}
	coll := client.Database(dbName).Collection(collName)
	return coll.Find(context.TODO(), bson.M{})
}

func (*mongoDBWorker) iterateByCursor(cursor *mongo.Cursor, dbName string, collName string, callBack func(map[string]interface{}), onFinish func()) {
	if cursor == nil {
		return
	}

	ctx := context.TODO()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc map[string]interface{}
		if err := cursor.Decode(&doc); err != nil {
			continue
		}
		callBack(doc)
	}
	onFinish()
}
