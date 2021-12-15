package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFileGeneratorBeginSimple(t *testing.T) {

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
		t.Run(fmt.Sprintf("ConfigFileGenerator with %v items", tc.givenDocs), func(t *testing.T) {
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

			assert.Equal(t, tc.expected, givenGenerator.SSConfig)

		})
	}
}

func TestConfigFileGeneratorBeginArray(t *testing.T) {
	doc1 := &map[string]interface{}{"name": []string{"john"}}
	doc2 := &map[string]interface{}{"name": []string{"john", "johnson"}}
	doc3 := &map[string]interface{}{"name": []float64{3.14}}

	doc4 := &map[string]interface{}{"name": []interface{}{"john", 125, true, 1156}}

	cases := []struct {
		givenFileType string
		givenDocs     []*map[string]interface{}
		expected      *SSConfig
	}{
		{"yaml", []*map[string]interface{}{doc1}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_Double}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}, {Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_Double}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}, {Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_Double}}}}}}}}}}}}}},

		{"yaml", []*map[string]interface{}{doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}, Item{Type: Type{Type: SS_Integer}}, Item{Type: Type{Type: SS_Boolean}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}}}}}, {Item: Item{Type: Type{Type: SS_Array, ElementType: []interface{}{Item{Type: Type{Type: SS_String}}, Item{Type: Type{Type: SS_Integer}}, Item{Type: Type{Type: SS_Boolean}}}}}}}}}}}}}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("ConfigFileGenerator with %v items", tc.givenDocs), func(t *testing.T) {
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

			assert.Equal(t, tc.expected, givenGenerator.SSConfig)

		})
	}
}

func TestConfigFileGeneratorBeginObject(t *testing.T) {
	doc1 := &map[string]interface{}{"name": map[string]interface{}{"fname": "john"}}
	doc2 := &map[string]interface{}{"name": map[string]interface{}{"fname": "johnson"}}
	doc3 := &map[string]interface{}{"name": map[string]interface{}{"lname": "tomson"}}
	doc4 := &map[string]interface{}{"name": map[string]interface{}{"fname": "tom", "age": 25}}

	cases := []struct {
		givenFileType string
		givenDocs     []*map[string]interface{}
		expected      *SSConfig
	}{
		{"yaml", []*map[string]interface{}{doc1}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "lname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}, Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc2}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}, {Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "lname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc1, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}, {Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}, Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc3}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}, {Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "lname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc2, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}, {Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}, Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
		{"yaml", []*map[string]interface{}{doc3, doc4}, &SSConfig{Databases: []Database{{D_name: "DB", Collection: Collection{C_name: "Coll", Fields: []Field{{F_name: "name", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "lname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}, {Item: Item{Type: Type{Type: SS_Object, ElementType: []interface{}{Field{F_name: "age", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_Integer}}}}}, Field{F_name: "fname", Constraints: []Constraint{{Item: Item{Type: Type{Type: SS_String}}}}}}}}}}}}}}}}},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("ConfigFileGenerator with %v items", tc.givenDocs), func(t *testing.T) {
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

			assert.Equal(t, tc.expected, givenGenerator.SSConfig)
		})
	}
}
