package cf

import (
	"sync"

	"github.com/cchaiyatad/seestern/internal/log"
)

type ConfigFileGenerator struct {
	FileType string

	OutChan   chan struct{}
	outChanWG *sync.WaitGroup
	SSConfig  *SSConfig

	treeChan   chan *SchemaTree
	treeChanWG *sync.WaitGroup
}

func NewConfigFileGenerator(fileType string) *ConfigFileGenerator {
	configGen := &ConfigFileGenerator{
		treeChan: make(chan *SchemaTree),
		OutChan:  make(chan struct{}),
		SSConfig: &SSConfig{},

		FileType:   fileType,
		treeChanWG: &sync.WaitGroup{},
		outChanWG:  &sync.WaitGroup{},
	}

	go func() {
		for tree := range configGen.treeChan {
			ssConfig := tree.ToSSConfig()
			configGen.SSConfig.Databases = append(configGen.SSConfig.Databases, ssConfig.Databases...)
			configGen.OutChan <- struct{}{}
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

func (configGen *ConfigFileGenerator) Close() {
	close(configGen.OutChan)
}

func (configGen *ConfigFileGenerator) Bytes() ([]byte, error) {
	enc := NewEncoder(configGen.FileType)
	if err := enc.Encode(configGen.SSConfig); err != nil {
		log.Log(log.Warning, err)
		return nil, err
	}
	return enc.Buf.Bytes(), nil
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

		configGen.outChanWG.Add(1)
		configGen.treeChan <- gatheredTree
		configGen.treeChanWG.Done()
	}()

	return extractFunc, onFinishFunc
}
