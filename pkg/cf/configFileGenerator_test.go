package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFileGeneratorBegin(t *testing.T) {

	doc1 := &map[string]interface{}{"name": "john"}
	doc2 := &map[string]interface{}{"name": "jane"}
	doc3 := &map[string]interface{}{"name": "johnson", "age": 15}
	doc4 := &map[string]interface{}{"name": "jeff", "age": "25"}

	cases := []struct {
		givenFileType string
		givenDocs     []*map[string]interface{}
		expected      *SSConfig
	}{
		{"yaml", []*map[string]interface{}{}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll"}}}}},
		{"yaml", []*map[string]interface{}{doc1}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}, {F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc3, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}, {Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc3, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}, {Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc3, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}, {Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2, doc3, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}, {F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}, {Item: Item{Type: Type{Type: SS_String}}}}}}}}}}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("ConfigFileGenerator with %d items", len(tc.givenDocs)), func(t *testing.T) {
			t.Parallel()
			givenDBName := "DB"
			givenCollName := "Coll"
			givenGenerator := NewConfigFileGenerator(tc.givenFileType)

			givenCallback, givenOnFinish := givenGenerator.Begin(givenDBName, givenCollName)

			for _, doc := range tc.givenDocs {
				givenCallback(*doc)
			}

			givenOnFinish()

			go func() {
				for range givenGenerator.OutChan {
					givenGenerator.Done()
				}
			}()

			givenGenerator.Wait()
			givenGenerator.Close()

			// Assert tree
			assert.Equal(t, tc.expected, givenGenerator.SSConfig)

		})
	}

}

// {"yaml", []*map[string]interface{}{}, 0},
// {"yaml", []*map[string]interface{}{doc1}, 1},
// {"yaml", []*map[string]interface{}{doc2}, 1},
// {"yaml", []*map[string]interface{}{doc3}, 1},
// {"yaml", []*map[string]interface{}{doc4}, 1},
// {"yaml", []*map[string]interface{}{doc1, doc2}, 2},
// {"yaml", []*map[string]interface{}{doc1, doc3}, 2},
// {"yaml", []*map[string]interface{}{doc1, doc4}, 2},
// {"yaml", []*map[string]interface{}{doc2, doc3}, 2},
// {"yaml", []*map[string]interface{}{doc2, doc4}, 2},
// {"yaml", []*map[string]interface{}{doc3, doc4}, 2},
// {"yaml", []*map[string]interface{}{doc1, doc2, doc3}, 3},
// {"yaml", []*map[string]interface{}{doc1, doc2, doc4}, 3},
// {"yaml", []*map[string]interface{}{doc1, doc3, doc4}, 3},
// {"yaml", []*map[string]interface{}{doc2, doc3, doc4}, 3},
// {"yaml", []*map[string]interface{}{doc1, doc2, doc3, doc4}, 4},
