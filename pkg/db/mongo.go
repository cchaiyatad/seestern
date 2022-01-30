package db

import (
	"sync"

	"github.com/cchaiyatad/seestern/internal/log"
	"github.com/cchaiyatad/seestern/pkg/cf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoDBController struct {
	cntStr string

	client     *mongo.Client
	initClient sync.Once
}

func createMongoDBController(connectionString string) *mongoDBController {
	return &mongoDBController{cntStr: connectionString}
}

func (w *mongoDBController) connect() (*mongo.Client, error) {
	var err error
	w.initClient.Do(func() {
		ctx, cancel := getCtxForConnect()
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

func (w *mongoDBController) ping() error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	ctx, cancel := getCtxForTransaction()
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}

func (w *mongoDBController) PS(dbNameFilter string) (NameRecord, error) {
	record, err := w.getNameRecord()
	if err != nil || dbNameFilter == "" {
		return record, err
	}

	specificDBInfo := make(NameRecord)
	if colls, ok := record[dbNameFilter]; ok {
		specificDBInfo[dbNameFilter] = colls
	}

	return specificDBInfo, nil
}

func (w *mongoDBController) InitConfigFile(targetColls []string, configGenerator *cf.ConfigFileGenerator) error {
	records, err := w.getNameRecord()
	if err != nil {
		return err
	}

	collsToGen := parseCollectionInputFromArgs(targetColls)

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

func (w *mongoDBController) Insert(dbName, collName string, documents ...interface{}) error {
	client, err := w.connect()
	if err != nil {
		return err
	}

	coll := client.Database(dbName).Collection(collName)

	ctx, cancel := getCtxForTransaction()
	defer cancel()
	_, err = coll.InsertMany(ctx, documents)
	return err
}

func (w *mongoDBController) Drop(dbName, collName string) error {
	client, err := w.connect()
	if err != nil {
		return err
	}
	ctx, cancel := getCtxForTransaction()
	defer cancel()
	return client.Database(dbName).Collection(collName).Drop(ctx)
}

func (w *mongoDBController) getNameRecord() (NameRecord, error) {
	client, err := w.connect()
	if err != nil {
		return nil, err
	}

	record := make(NameRecord)

	ctx, cancel := getCtxForTransaction()
	defer cancel()
	dbs, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	for _, db := range dbs {
		ctx, cancel := getCtxForTransaction()
		defer cancel()
		colls, err := client.Database(db).ListCollectionNames(ctx, bson.D{})
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

func (w *mongoDBController) getCursor(dbName string, collName string) (*mongo.Cursor, error) {
	client, err := w.connect()
	if err != nil {
		return nil, err
	}
	coll := client.Database(dbName).Collection(collName)
	ctx, cancel := getCtxForTransaction()
	defer cancel()
	return coll.Find(ctx, bson.M{})
}

func (*mongoDBController) iterateByCursor(cursor *mongo.Cursor, dbName string, collName string, callBack func(map[string]interface{}), onFinish func()) {
	if cursor == nil {
		return
	}
	ctx, cancel := getCtxForTransaction()
	defer cancel()
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
