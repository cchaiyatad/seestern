# SEESTERN SYNTAX

A reference for configuration file. This syntax reference examples were written in TOML format.

## Database (Database)

This object represent the database.

### field: d_name

- **Description**: Use to assign database name
- **Value type**: string
- **Require**: yes
- **Example**:
  
``` toml
d_name = "school"
```

### field: collection

- **Description**: Use to represent a collection inside database
- **Value type**: Database.Collection
- **Require**: yes
- **Example**:
  
``` toml
[[databases]]
d_name = "school"

[databases.collection]
c_name = "student"
count = 30
```

## Collection (Database.Collection)

This object represent the collection inside database.

### field: c_name

- **Description**: Use to assign collection name
- **Value type**: string
- **Require**: yes
- **Example**:
  
``` toml
c_name = "student"
```

### field: count

- **Description**: Use to assign number of document to be generate for the collection
- **Value type**: integer
- **Require**: yes
- **Example**:
  
``` toml
count = 30
```

### field: fields

- **Description**: Use to assign a field that will be in generated document
- **Value type**: \[Database.Collection.Field\]
- **Require**: yes
- **Example**:
  
``` toml
fields = [{f_name = "name", constraint = [ {type="string"} ]}]
```

## Field (Database.Collection.Field)

This object represent a field of object. A document is also an object.

### field: f_name

- **Description**: Use to assign a name of the field
- **Value type**: string
- **Require**: yes
- **Example**:

``` toml
f_name = "name"
```

### field: omit_weight

- **Description**: Use to assign a possibility whether the field should be omited or not. 1 is always omit, 0 is always show.
- **Value type**: double (between 0.0 to 1.0)
- **Require**: no (default = 0.0)
- **Example**:

``` toml
omit_weight = 0.5
# which mean if there are ten documents, five of them will not have this field
```

### field: constraint

- **Description**: Use to assign constraint(s) of the field. if more than one Constraint are given, the possibilty that the field constraint will be decide base on weight field (see Constraint below)
- **Value type**: \[Database.Collection.Field.Constraint\]
- **Require**: no
- **Example**:

``` toml
# example 1
# this field will always be string
constraint = [{type = "string"}]

# example 2
# this field will have 50% chance to be string and 50% chance to be integer
constraint = [{type = "string"}, {type = "integer"}]
```

### field: set

- **Description**: Use to specific fields in the object to have specific value. Index of the object is start at 0
- **Value type**: \[Database.Collection.Field.Set\]
- **Require**: no
- **Example**:

``` toml
# since index start at 0
# the field in second, third and fouth object will have value equal to "alice"
# the field in first object and other will be (random) string
set = [{value = "alice", at = [1, 2, 3]}]
constraint = [{type = "string"}]
```

If both **constraint** and **set** are not given and not be omited, the value of this field will be **null**

## Constraint (Database.Collection.Field.Constraint)

This object represent the constraint of the generated value.

In this object, there is one optional field, **weight**, and one required field which can be either **value**, **enum** or **type**.

### field: weight

- **Description**: the possibility ratio that this constraint will be picked to generate
- **Value type**: int
- **Require**: no (default 1)
- **Example**:

``` toml
weight = 2

# for example
# if we want 75% of generate documents/objectß to has a sex = "M" 
# and the other will be "F" it can be written like this
[[database.collection.fields]]
f_name = "sex" 

[[database.collection.fields.constraint]]
value = "M"
weight = 3


[[database.collection.fields.constraint]]
value = "F"
```

### field: value

- **Description**: see **Database.Collection.Field.Value**
- **Value type**: Database.Collection.Field.Value
- **Require**: no, if enum or type was already assigned
- **Example**: see **Database.Collection.Field.Value**

### field: enum

- **Description**: see **Database.Collection.Field.Enum**
- **Value type**: Database.Collection.Field.Enum
- **Require**: no, if value or type was already assigned
- **Example**: see **Database.Collection.Field.Enum**

### field: type

- **Description**: see **Database.Collection.Field.Type**
- **Value type**: Database.Collection.Field.Type
- **Require**: no, if value or enum was already assigned
- **Example**: see **Database.Collection.Field.Type**

### example

Constraint with Value

``` toml
[[databases.collection.fields.constraints]]
value = 3.14
```

Constraint with Enum

``` toml
[[databases.collection.fields.constraints]]
enum = ["freshman", "sophomore", "junior", "senior"]
weight = 2
```

Constraint with Type

``` toml
[[databases.collection.fields.constraints]]
type = "boolean"
```

## Set (Database.Collection.Field.Set)

This object represent the constraint of the generated value **at specific index**.

In this object, there is two required fields, **at**, and either **value**, **enum** or **type**.

### field: at

- **Description**: Use to specific the order that will have a specific value that is define by set. Index of the object is start at 0
- **Value type**: [int]
- **Require**: yes
- **Example**:

``` toml
# since index start at 0
# the field in second, third and fouth object will have that is specific in Set
# the field in first object and other will be (random) string
at = [1, 2, 3]
```

### field: value

- **Description**: see **Database.Collection.Field.Value**
- **Value type**: Database.Collection.Field.Value
- **Require**: no, if enum or type was already assigned
- **Example**: see **Database.Collection.Field.Value**

### field: enum

- **Description**: see **Database.Collection.Field.Enum**
- **Value type**: Database.Collection.Field.Enum
- **Require**: no, if value or type was already assigned
- **Example**: see **Database.Collection.Field.Enum**

### field: type

- **Description**: see **Database.Collection.Field.Type**
- **Value type**: Database.Collection.Field.Type
- **Require**: no, if value or enum was already assigned
- **Example**: see **Database.Collection.Field.Type**

### example

Set with Value

``` toml
[[databases.collection.fields.sets]]
value = 3.14
at = [0, 1]
```

Set with Enum

``` toml
[[databases.collection.fields.sets]]
enum = ["freshman", "sophomore", "junior", "senior"]
at = [20]
```

Set with Type

``` toml
[[databases.collection.fields.sets]]
type = "boolean"
at = [3, 5, 7]
```

## Value (Database.Collection.Field.Value)

This object is used to assign a specific value to generate.

### field: value

- **Description**:  Use to assign a field that will be in generated document
- **Value type**: any
- **Require**: yes
- **Example**:

``` toml
# simple
value = 420

# for array
value = ["test", "array"]

# for object
[[databases.collection.fields]]
f_name = "name"

[[databases.collection.fields.constraints]]
constraints = {value = {first_name = "John", last_name = "Doe"}}

# or just
[[databases.collection.fields]]
f_name = "name"
constraints = [{value = {first_name = "John", last_name = "Doe"}}]
```

## Enum (Database.Collection.Field.Enum)

This object is used to assign specific values that will be randomly picked to generate.

### field: enum

- **Description**: Used to assign specific values that will be randomly picked to generate.
- **Value type**: [any]
- **Require**: yes
- **Example**:

``` toml
# simple
enum = ["freshman", "sophomore", "junior", "senior"]

# for array
enum = [["test", "array", "one"], ["test", "array", "two"]]


# for object
# this will generate either  "name": {"first_name": "John", "last_name": "Doe" } 
# or "name": {"first_name": "Jane", "last_name": "Doe" }
[[databases.collection.fields]]
f_name = "name"

[[databases.collection.fields.constraints]]
enum = [{first_name = "John", last_name = "Doe"}, {first_name = "Jane", last_name = "Doe"}]
```

## Type (Database.Collection.Field.Type)

This object is used to assign specific type of value that will be generate. Each type will have it own properties. Also, if this object is assigned inside the **Constraint (Database.Collection.Field.Constraint)**, it will have one additional field ,**ref**.

### field: type

- **Description**: type of data that will be generated
- **Value type**: string (only "string", "integer", "double", "boolean", "null", "objectId", "array", "object")
- **Require**: yes
- **Example**:

``` toml
type = "string"
```

### field: ref

- **Description**: use to make this field to be a foreign key to other collection
- **Value type**: string (in the format "database.collection.document_field\[.document_field_1.document_field_2...\])
- **Require**: no
- **Example**:

``` toml
# reference to database school collection student field name
ref = "school.student.name"
```

if the value in **type** field is string, int, double, array or object. the **Type** will has additional field.

### when type = "string"

#### field: prefix

- **Description**: prefix of the generated value
- **Value type**: string
- **Require**: no
- **Example**:

``` toml
prefix = "Al"
```

#### field: suffix

- **Description**: suffix of the generated value
- **Value type**: string
- **Require**: no
- **Example**:

``` toml
prefix = "est"
```

#### field: length

- **Description**: maximum length of generated string
- **Value type**: int (default = 20)
- **Require**: no
- **Example**:

``` toml
length = 15
```

### when type = "integer"

#### field: min

- **Description**: minimum value of the generated value (inclusive)
- **Value type**: int (default 0)
- **Require**: no
- **Example**:

``` toml
min = -4 # ([4, max))
```

#### field: max

- **Description**: maximum value of the generated value (exclusive)
- **Value type**: int (if not given or less than min field value, this value will equal to min + 100)
- **Require**: no
- **Example**:

``` toml
max = 100 # ([min, 100))
```

### when type = "double"

#### field: min

- **Description**: minimum value of the generated value (inclusive)
- **Value type**: double (default 0)
- **Require**: no
- **Example**:

``` toml
min = -4.5 # ([4.5, max))
```

#### field: max

- **Description**: maximum value of the generated value (exclusive)
- **Value type**: double (if not given or less than min field value, this value will equal to min + 100)
- **Require**: no
- **Example**:

``` toml
max = 37.35 # ([min, 37.35))
```

### when type = "array"

#### field: element_type

- **Description**: a possible constraint of member in generated array
- **Value type**: \[Database.Collection.Field.Constraint\]
- **Require**: yes
- **Example**:

``` toml
# member in generated array can be a value 5 or a random string
element_type = [ {value = 5, weight = 2}, {type = "string"}]
```

#### field: set

- **Description**: a value of member in generated array
- **Value type**: \[Database.Collection.Field.Set\]
- **Require**: yes
- **Example**:

``` toml
# member in generated array index equal to 1 or 3 will has value = [3.14, "test", "array"] 
set = [{value = [3.14, "test", "array"], at = [1, 3]}]
```

#### field: min_item

- **Description**: minimum number of members of the generated value (inclusive)
- **Value type**: int (default 0)
- **Require**: no
- **Example**:

``` toml
min_item = 3
```

#### field: max_item

- **Description**: maximum number of members of the generated value (inclusive)
- **Value type**: int (default 10)
- **Require**: no
- **Example**:

``` toml
max_item = 30

# if one want a exact number of member,
# it can be done by assign min = max value
# Ex. the following array will have exact 2 members
min_item = 2
max_item = 2
```

### when type = "object"

#### field: fields

- **Description**: a field of the generated object
- **Value type**: \[Database.Collection.Field\]
- **Require**: yes
- **Example**:

``` toml
fields = [
    {f_name = "class_name", constraint = [{type = "string"}]}, 
    {f_name = "instructor", constraint = [{type = "string"}]}, 
]
# the above value will represent the following object
# { 
#    "class_name": string,
#    "instructor": string
# } 
```

## Alias (Alias)

work like a variable to reduce the number of duplicate value. **Only support in toml**.

### example

if there is a configuration file that was written like this

``` toml
[[collection]]
c_name = "student"
count = 10

[[database.collection.fields]] 
f_name = "first_name" 
constraint = [{type = "string", length = 12}]

[[database.collection.fields]] 
f_name = "last_name" 
constraint = [{type = "string", length = 12}]
```

the `[{type = "string", length = 12}]` was written two time so to reduce the duplicate code, with Alias it can be written like this

``` toml
[[collection]]
c_name = "student"
count = 10

[[database.collection.fields]]
f_name = "first_name"
constraint = "#{{string_12_constraint}}"

[[database.collection.fields]]
f_name = "last_name"
constraint = "#{{string_12_constraint}}"

[[alias]]
key = "string_12_constraint"
value = [{type = "string", length = 12}]
```

when using the alias, we have to assign its key and value before, then the `#{{key_name}}` will be interpret to the value that it represent. 

### field: Key

- **Description**: a key (name) of alias
- **Value type**: string
- **Require**: yes
- **Example**:

``` toml
key = "class_constraint"
```

### field: Value

- **Description**: a value that the it represent
- **Value type**: any
- **Require**: yes
- **Example**:

``` toml
value = [{type=”string”, length=12}]
```
