package db

import (
	"context"
	"fmt"
	"os"
	"time"

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

func (w *mongoDBWorker) initConfigFile(param *InitParam) (string, error) {
	client, err := w.connect()
	if err != nil {
		return "", err
	}

	infos, err := w.getDatabaseCollectionInfoWithClient(client)
	if err != nil {
		return "", err
	}

	toGenColls := parseCollectionInputFromArgs(param.TargetColls)
	schemaTreeChan := make(chan *cf.SchemaTree)

	for db, colls := range toGenColls {
		for _, coll := range colls {
			if _, ok := infos[db][coll]; !ok {
				fmt.Fprintf(os.Stderr, "skip: database %s collection %s: not exist\n", db, coll)
				continue
			}
			fmt.Printf("generate: database %s collection %s", db, coll)
			cursor, err := w.getCursor(client, db, coll)
			if err != nil {
				fmt.Fprintf(os.Stderr, "skip: database %s collection %s: error: %s\n", db, coll, err)
				continue
			}
			go w.extractSchema(cursor, schemaTreeChan)
		}
	}

	for tree := range schemaTreeChan {
		fmt.Println(tree)
	}

	// save to file

	// return path, error
	return "", nil
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
			fmt.Printf("skip database %s :%s", db, err)
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
	coll := client.Database(dbName).Collection(collName)
	return coll.Find(context.TODO(), bson.M{})
}

func (*mongoDBWorker) extractSchema(cursor *mongo.Cursor, out chan<- *cf.SchemaTree) {
	ctx := context.TODO()
	defer cursor.Close(ctx)

	treeChan := make(chan *cf.SchemaTree)
	go func() {
		for cursor.Next(ctx) {
			var tree map[string]interface{}
			if err := cursor.Decode(&tree); err != nil {
				continue
			}
			treeChan <- cf.ParseSchemaTree(tree)
		}

		close(treeChan)
	}()

	gatheredTree := cf.NewSchemaTree()

	for tree := range treeChan {
		if newTree, err := cf.MergeSchemaTree(gatheredTree, tree); err == nil {
			gatheredTree = newTree
		}
	}

	out <- gatheredTree

}
