package cf

import (
	"log"
	"sync"
)

type ConfigFileGenerator struct {
	FileType string

	OutChan   chan []byte
	outChanWG *sync.WaitGroup

	treeChan   chan *SchemaTree
	treeChanWG *sync.WaitGroup
}

func NewConfigFileGenerator(fileType string) *ConfigFileGenerator {
	configGen := &ConfigFileGenerator{
		treeChan: make(chan *SchemaTree),
		OutChan:  make(chan []byte),

		FileType:   fileType,
		treeChanWG: &sync.WaitGroup{},
		outChanWG:  &sync.WaitGroup{},
	}

	go func() {
		for tree := range configGen.treeChan {

			enc := NewEncoder(configGen.FileType)
			if err := enc.Encode(tree.ToSSConfig()); err != nil {
				log.Fatal(err)
			}
			configGen.OutChan <- enc.Buf.Bytes()
		}
	}()

	return configGen
}

func (configGen *ConfigFileGenerator) Wait() {
	configGen.treeChanWG.Wait()
	configGen.outChanWG.Wait()
}

func (configGen *ConfigFileGenerator) Done() {
	configGen.outChanWG.Done()
}

func (configGen *ConfigFileGenerator) Begin(dbName, collName string) (extractFunc func(map[string]interface{}), onFinishFunc func()) {
	configGen.treeChanWG.Add(1)
	docChan := make(chan *SchemaTree)
	gatheredTree := newSchemaTree(dbName, collName)

	extractFunc = func(doc map[string]interface{}) {
		docChan <- parseSchemaTree(dbName, collName, doc)
	}
	onFinishFunc = func() {
		close(docChan)
	}

	go func() {
		for tree := range docChan {
			if newTree, err := mergeSchemaTree(gatheredTree, tree); err == nil {
				gatheredTree = newTree
			}
		}

		configGen.treeChan <- gatheredTree
		configGen.outChanWG.Add(1)
		configGen.treeChanWG.Done()
	}()

	return extractFunc, onFinishFunc
}
