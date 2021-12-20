package cf

import (
	"testing"
)

func TestGen(t *testing.T) {
	ssconfig := &SSConfig{
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

						// {F_name: "gpa", Constraints: []Constraint{
						// 	{Item: Item{Type: Type{Type: "double"}}}}},
					}}}},
	}
	ssconfig.Gen()

}
