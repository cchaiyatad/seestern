package db

import (
	"context"
	"sync"
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

	client     *mongo.Client
	initClient sync.Once
}

func createMongoDBWorker(connectionString string) *mongoDBWorker {
	return &mongoDBWorker{cntStr: connectionString}
}

func (w *mongoDBWorker) connect() (*mongo.Client, error) {
	var err error
	w.initClient.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		w.client, err = mongo.Connect(ctx, options.Client().ApplyURI(w.cntStr))
	})

	if err != nil {
		return nil, err
	}

	if w.client == nil {
		return nil, ErrClientIsNil
	}

	return w.client, nil
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

func (w *mongoDBWorker) ps(dbNameFilter string) (nameRecord, error) {
	record, err := w.getNameRecord()
	if err != nil || dbNameFilter == "" {
		return record, err
	}

	specificDBInfo := make(nameRecord)
	if colls, ok := record[dbNameFilter]; ok {
		specificDBInfo[dbNameFilter] = colls
	}

	return specificDBInfo, nil
}

func (w *mongoDBWorker) initConfigFile(param *InitParam, configGenerator *cf.ConfigFileGenerator) error {
	records, err := w.getNameRecord()
	if err != nil {
		return err
	}

	collsToGen := parseCollectionInputFromArgs(param.TargetColls)

	for db, colls := range collsToGen {
		for _, coll := range colls {
			if _, ok := records[db][coll]; !ok {
				log.Logf(log.Warning, "%s\n", &ErrSkipCreateConfigfile{db, coll, "not exist"})
				continue
			}
			cursor, err := w.getCursor(db, coll)
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

func (w *mongoDBWorker) insert(dbName, collName string, documents ...interface{}) error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	coll := client.Database(dbName).Collection(collName)
	_, err = coll.InsertMany(context.TODO(), documents)
	return err
}

func (w *mongoDBWorker) drop(dbName, collName string) error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	return client.Database(dbName).Collection(collName).Drop(context.TODO())
}

func (w *mongoDBWorker) getNameRecord() (nameRecord, error) {
	client, err := w.connect()
	if err != nil {
		return nil, err
	}

	record := make(nameRecord)

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

		record[db] = make(map[string]struct{})
		for _, coll := range colls {
			record[db][coll] = struct{}{}
		}
	}
	return record, nil
}

func (w *mongoDBWorker) getCursor(dbName string, collName string) (*mongo.Cursor, error) {
	client, err := w.connect()
	if err != nil {
		return nil, err
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
