package cf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSSConfig(t *testing.T) {
	t.Parallel()
	t.Run("parse a config file from invalid file", func(t *testing.T) {
		cases := []struct {
			filePath string
			errorMsg string
		}{
			{"./../../test/config", "a given path is not a file: ./../../test/config"},
			{"./../../test/config/invalid", "error: not support file type (support only .json .toml .yaml)"},
			{"./../../test/config/not-exist", "stat ./../../test/config/not-exist: no such file or directory"},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetSSConfig from %s expected error", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotSSConfig, gotErr := NewConfigFileReader(tc.filePath).GetSSConfig()
				assert.Nil(t, gotSSConfig)

				assert.Equal(t, tc.errorMsg, gotErr.Error())
			})
		}
	})

	t.Run("parse a config file from init command", func(t *testing.T) {
		cases := []struct {
			filePath string
		}{
			{"./../../test/config/json/00_from_init.ss.json"},
			{"./../../test/config/yaml/00_from_init.ss.yaml"},
			{"./../../test/config/toml/00_from_init.ss.toml"},
		}
		expectedData := &SSConfig{
			Databases: []Database{
				{
					D_name: "sample_training",
					Collection: Collection{
						C_name: "trips",
						Fields: []Field{
							{F_name: "_id", Constraints: []Constraint{{Item: Item{Type: Type{Type: "objectID"}}}}},
							{F_name: "bikeid", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "birth year", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "integer"}}},
								{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "end station id", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "end station location", Constraints: []Constraint{{Item: Item{Type: Type{Type: "object", P_ElementType: []interface{}{map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"element_type": []interface{}{map[string]interface{}{"type": "double"}}, "type": "array"}}, "f_name": "coordinates"}, map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"type": "string"}}, "f_name": "type"}}}}}}},
							{F_name: "end station name", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "start station id", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "start station location", Constraints: []Constraint{{Item: Item{Type: Type{Type: "object", P_ElementType: []interface{}{map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"element_type": []interface{}{map[string]interface{}{"type": "double"}}, "type": "array"}}, "f_name": "coordinates"}, map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"type": "string"}}, "f_name": "type"}}}}}}},
							{F_name: "start station name", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "start time", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "stop time", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "tripduration", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "usertype", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}}}}},
				{
					D_name: "sample_training",
					Collection: Collection{
						C_name: "posts",
						Fields: []Field{
							{F_name: "_id", Constraints: []Constraint{{Item: Item{Type: Type{Type: "objectID"}}}}},
							{F_name: "author", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "body", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "comments", Constraints: []Constraint{{Item: Item{Type: Type{Type: "array", P_ElementType: []interface{}{map[string]interface{}{"element_type": []interface{}{map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"type": "string"}}, "f_name": "author"}, map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"type": "string"}}, "f_name": "body"}, map[string]interface{}{"constraints": []interface{}{map[string]interface{}{"type": "string"}}, "f_name": "email"}}, "type": "object"}}}}}}},
							{F_name: "date", Constraints: []Constraint{{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "permalink", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "tags", Constraints: []Constraint{{Item: Item{Type: Type{Type: "array", P_ElementType: []interface{}{map[string]interface{}{"type": "string"}}}}}}},
							{F_name: "title", Constraints: []Constraint{{Item: Item{Type: Type{Type: "string"}}}}}}}}}}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetSSConfig from %s (from init command)", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotSSConfig, gotErr := NewConfigFileReader(tc.filePath).GetSSConfig()
				assert.Nil(t, gotErr)
				assert.Equal(t, expectedData.String(), gotSSConfig.String())
			})
		}
	})

	t.Run("parse a config file from 01_simple", func(t *testing.T) {
		cases := []struct {
			filePath string
		}{
			{"./../../test/config/json/01_simple.ss.json"},
			{"./../../test/config/yaml/01_simple.ss.yaml"},
			{"./../../test/config/toml/01_simple.ss.toml"},
		}
		expectedData := &SSConfig{
			Databases: []Database{
				{
					D_name: "school",
					Collection: Collection{
						C_name: "student",
						Count:  30,
						Fields: []Field{
							{F_name: "s_id", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "objectId"}}}}},
							{F_name: "name", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "string", P_Prefix: "a", P_Suffix: "m", P_Length: 5}}}}},
							{F_name: "sex", Omit_weight: 5, Constraints: []Constraint{
								{Weight: 2, Item: Item{Value: Value{Value: "M"}}},
								{Weight: 3, Item: Item{Value: Value{Value: "F"}}}}},
							{F_name: "year",
								Constraints: []Constraint{
									{Item: Item{Enum: Enum{Enum: []interface{}{"freshman", "sophomore", "junior", "senior"}}}}},
								Sets: []Set{
									{At: []int{1, 2, 3}, Item: Item{Value: Value{Value: "super senior"}}},
									{At: []int{5}, Item: Item{Type: Type{Type: "integer", P_Min: 5, P_Max: 8}}}}}}}}},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetSSConfig from %s (from 01_simple)", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotSSConfig, gotErr := NewConfigFileReader(tc.filePath).GetSSConfig()
				assert.Nil(t, gotErr)
				assert.Equal(t, expectedData.String(), gotSSConfig.String())
			})
		}
	})

	t.Run("parse a config file from 02_two_coll", func(t *testing.T) {
		cases := []struct {
			filePath string
		}{
			{"./../../test/config/json/02_two_coll.ss.json"},
			{"./../../test/config/yaml/02_two_coll.ss.yaml"},
			{"./../../test/config/toml/02_two_coll.ss.toml"},
		}
		expectedData := &SSConfig{
			Databases: []Database{
				{
					D_name: "school",
					Collection: Collection{
						C_name: "student",
						Count:  30,
						Fields: []Field{
							{F_name: "s_id", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "objectId"}}}}},
							{F_name: "name", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "string"}}}}},
						},
					},
				},
				{
					D_name: "school",
					Collection: Collection{
						C_name: "teacher",
						Count:  15,
						Fields: []Field{
							{F_name: "t_id", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "objectId"}}}}},
							{F_name: "name", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "age", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "integer", P_Min: 30}}}}},
						},
					},
				},
			},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(fmt.Sprintf("GetSSConfig from %s (from 01_simple)", tc.filePath), func(t *testing.T) {
				t.Parallel()
				gotSSConfig, gotErr := NewConfigFileReader(tc.filePath).GetSSConfig()
				assert.Nil(t, gotErr)
				assert.Equal(t, expectedData.String(), gotSSConfig.String())
			})
		}
	})
}

// TODO: Fix Alias
// func TestGetSSConfigWithAlias(t *testing.T) {

// 	filePath := "./../../test/config/toml/03_simple_alias.ss.toml"

// 	expectedData := &SSConfig{
// 		Databases: []Database{
// 			{
// 				D_name: "school",
// 				Collection: Collection{
// 					C_name: "student",
// 					Count:  30,
// 					Fields: []Field{
// 						{F_name: "s_id", Constraints: []Constraint{
// 							{Item: Item{Type: Type{Type: "objectId"}}}}},
// 						{F_name: "name", Constraints: []Constraint{
// 							{Item: Item{Type: Type{Type: "string"}}}}},
// 						{F_name: "class", Constraints: []Constraint{
// 							{Item: Item{Type: Type{Type: "string", P_Length: 10}}}}},
// 						{F_name: "elective_class", Constraints: []Constraint{
// 							{Item: Item{Type: Type{Type: "string", P_Length: 10}}}}},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	gotSSConfig, gotErr := NewConfigFileReader(filePath).GetSSConfig()
// 	assert.Nil(t, gotErr)
// 	assert.Equal(t, expectedData.String(), gotSSConfig.String())
// }

// "./../../test/config/project/01_configSpec_simple.ss.toml"
// "./../../test/config/project/02_configSpec_array.ss.toml"
// "./../../test/config/project/03_configSpec_object.ss.toml"
// "./../../test/config/project/04_configSpec_alias.ss.toml"
// "./../../test/config/project/05_configSpec_embedded.ss.toml"
// "./../../test/config/project/06_configSpec_refs.ss.toml"
