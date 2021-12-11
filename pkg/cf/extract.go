package cf

import (
	"sync"
)

type SchemaExtracter struct {
	TreeChan chan *SchemaTree

	TreeChanWG *sync.WaitGroup
}

func NewSchemaExtracter() *SchemaExtracter {
	return &SchemaExtracter{
		TreeChan:   make(chan *SchemaTree),
		TreeChanWG: &sync.WaitGroup{},
	}
}

func (se *SchemaExtracter) GetExtractSchemaFunc(dbName, collName string) (extractFunc func(map[string]interface{}), onFinishFunc func()) {
	se.TreeChanWG.Add(1)
	docChan := make(chan *SchemaTree)
	gatheredTree := NewSchemaTree(dbName, collName)

	extractFunc = func(doc map[string]interface{}) {
		docChan <- ParseSchemaTree(dbName, collName, doc)
	}
	onFinishFunc = func() {
		close(docChan)
	}

	go func() {
		for tree := range docChan {
			if newTree, err := MergeSchemaTree(gatheredTree, tree); err == nil {
				gatheredTree = newTree
			}
		}
		se.TreeChan <- gatheredTree
		// se.TreeChanWG.Done()
	}()

	return extractFunc, onFinishFunc
}
