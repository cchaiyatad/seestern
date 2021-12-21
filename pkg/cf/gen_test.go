package cf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenWithError(t *testing.T) {
	t.Run("SSConfig that doesn't has databases", func(t *testing.T) {
		givenDB := "school"
		givenColl := "student"
		givenConfig := &SSConfig{Databases: []Database{}}
		gotResult := givenConfig.Gen()

		expectedResult := make(result)
		assert.Equal(t, expectedResult, gotResult)

		expected := "database or collection name doesn't exist in result"
		gotDocuments, gotErr := gotResult.GetDocuments(givenDB, givenColl)
		assert.Nil(t, gotDocuments)
		assert.Equal(t, expected, gotErr.Error())
	})

	t.Run("SSConfig that doesn't has count", func(t *testing.T) {
		givenDB := "school"
		givenColl := "student"
		givenConfig := &SSConfig{
			Databases: []Database{
				{
					D_name: "school",
					Collection: Collection{
						C_name: "student",
						Fields: []Field{
							{F_name: "name", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "age", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "integer"}}}}},
						}}}},
		}
		gotResult := givenConfig.Gen()

		expected := "count have to be more than zero got: 0 (db: school, coll: student)"
		gotDocuments, gotErr := gotResult.GetDocuments(givenDB, givenColl)
		assert.Nil(t, gotDocuments)

		assert.Equal(t, expected, gotErr.Error())
	})
}

func TestGen(t *testing.T) {
	t.Run("SSConfig with simple constraint", func(t *testing.T) {
		givenDB := "school"
		givenColl := "student"
		givenConfig := &SSConfig{
			Databases: []Database{
				{
					D_name: "school",
					Collection: Collection{
						C_name: "student",
						Count:  10,
						Fields: []Field{
							{F_name: "name", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "string"}}}}},
							{F_name: "age", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "integer"}}}}},
							{F_name: "isHonor", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "boolean"}}}}},
							{F_name: "invalid", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "null"}}}}},
							{F_name: "gpa", Constraints: []Constraint{
								{Item: Item{Type: Type{Type: "double"}}}}},
						}}}},
		}
		gotResult := givenConfig.Gen()

		_, gotErr := gotResult.GetDocuments(givenDB, givenColl)
		assert.Nil(t, gotErr)
	})
}
