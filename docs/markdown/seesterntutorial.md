# SEESTERN TUTORIAL

A short instruction to install and use program.

## Installation

1. install go version 1.7 or later
2. run `make install`

## Tutorial

Before generate the test data, a tool need configuration file that can either generate by the tool or write by hand. This tutorial will cover both case, but for deep understand of syntax please see seesternsyntax.

### use tool to create configuration file

In this section, you will learn how to use **ps** to list a collection inside database, **init** to create configuration file and **generate** command to generate data.

#### 1. Set up database

We will use [MongoDB with sample dataset](https://hub.docker.com/r/weshigbee/docker-mongo-sample-datasets) docker.
Execute `$ docker run -it --rm -p 27017:27017 weshigbee/docker-mongo-sample-datasets`. now the connection string for this database is `mongodb://localhost:27017`

#### 2. list database and collection with ps command

Now if we want to list database and its collections inside.
Execute `$ seestern ps -s mongodb://localhost:27017`.
The output should look like this

``` text
[    Info]: 2022/03/13 15:25:22 list collections form mongodb://localhost:27017 connection string
database: test
 1 : restaurants
database: admin
 1 : system.version
database: config
 1 : system.sessions
database: local
 1 : startup_log
```

The first line is a log that was printed on standard error and the rest was printed on standard output.
In this MongoDB instance, there are four databases but we only interested on `test` database. We can also, list collections in specific database by execute `$ seestern ps -s mongodb://localhost:27017 -d test` now an output should look like this

``` text
[    Info]: 2022/03/13 15:33:58 list collections form mongodb://localhost:27017 connection string
database: test
 1 : restaurants
```

#### 3. generate configuration file with init command

After we know the name of database and collection that we want to generate a configuration file, we then execute `seestern init -s mongodb://localhost:27017 -c "test.restaurants" -o "."`.

After executing, the result will look like this

``` text
[    Info]: 2022/03/13 15:42:27 init with mongodb://localhost:27017 connection string
[    Info]: 2022/03/13 15:42:27 generate: database test collection restaurants
[    Info]: 2022/03/13 15:42:29 config file is saved to 1647160949.ss.yaml
```

The `1647160949.ss.yaml` is configuration file that contain schema of data inside restaurants collection. The file should look like this

``` yaml
databases:
    - d_name: test
      collection:
        c_name: restaurants
        fields:
            - f_name: _id
              constraints:
                - type: objectID
            - f_name: address
              constraints:
                - type: object
                  fields:
                    - f_name: building
                      constraints:
                        - type: string
                    - f_name: coord
                      constraints:
                        - type: array
                          element_type:
                            - type: double
                    - f_name: street
                      constraints:
                        - type: string
                    - f_name: zipcode
                      constraints:
                        - type: string
                - type: object
                  fields:
                    - f_name: building
                      constraints:
                        - type: string
                    - f_name: coord
                      constraints:
                        - type: array
                    - f_name: street
                      constraints:
                        - type: string
                    - f_name: zipcode
                      constraints:
                        - type: string
            - f_name: borough
              constraints:
                - type: string
            - f_name: cuisine
              constraints:
                - type: string
            - f_name: grades
              constraints:
                - type: array
                  element_type:
                    - type: object
                      fields:
                        - f_name: date
                          constraints:
                            - type: integer
                        - f_name: grade
                          constraints:
                            - type: string
                        - f_name: score
                          constraints:
                            - type: integer
                - type: array
                  element_type:
                    - type: object
                      fields:
                        - f_name: date
                          constraints:
                            - type: integer
                        - f_name: grade
                          constraints:
                            - type: string
                - type: array
            - f_name: name
              constraints:
                - type: string
            - f_name: restaurant_id
              constraints:
                - type: string

```

Our tool analyzed and found that the document inside `restaurants` collection has following fields.

- **_id** as id
- **address** as object that has field
  - **building** as string
  - **coord** as array contains (can be empty) double
  - **street** as string
  - **zipcode** as string
- **borough** as string
- **cuisine** as string
- **grades** as array contains (can be empty) objects that has field
  - **date** as integer
  - **grade** as string
  - **score** as integer (can be left off)
- **name** as string
- **restaurant_id** as string

One thing to noted that, at this version, the tool doesn't support datetime type so the date is convert to integer.

#### 4. generate data with generate command

After we got configuration file (`1647160949.ss.yaml`), we can execute `$ seestern generate -f 1647160949.ss.yaml -v` to see the output. Which will look like this

``` text
[    Info]: 2022/03/13 16:02:26 generate with configuration file 1647160949.ss.yaml
[ Warning]: 2022/03/13 16:02:26 can not generate database test collection restaurants with reason count have to be more than zero got: 0 (db: test, coll: restaurants)
// database test collection restaurants
null
```

That's right, we can not generate yet. We have to specific field `count` which will determine how many documents we will generate first. So, edit configuration file like this

``` yaml
databases:
    - d_name: test
      collection:
        c_name: restaurants
        count: 3 # add this line
        fields:
....
```

Then if we generate with the same command an output should look like this

``` text
[    Info]: 2022/03/13 16:09:24 generate with configuration file 1647160949.ss.yaml
// database test collection restaurants
[
    {
        "_id": "622db71c1b325ba6db97c3ea",
        "address": {
            "building": "0Sc0Rqv13BqKi",
            "coord": [],
            "street": "CjD",
            "zipcode": "e"
        },
        "borough": "u1pH",
        "cuisine": "lalyiBMhsazbyUQj6u",
        "grades": [],
        "name": "QP73VsqMOiKS4NtG",
        "restaurant_id": "RPsm"
    },
    {
        "_id": "622db71c1b325ba6db97c3eb",
        "address": {
            "building": "",
            "coord": [
                96.74254955040003,
                3.8138881021617816,
                43.45470891544519
            ],
            "street": "RXx1miAaOnSGb",
            "zipcode": "cfnyE2MoJi"
        },
        "borough": "XDJ0ERhMP09XFPmT",
        "cuisine": "8eV8DY0Ge",
        "grades": [],
        "name": "UdRDphTnhS2p2abt",
        "restaurant_id": "Qa5zY9vkbg8qUfboB"
    },
    {
        "_id": "622db71c1b325ba6db97c3ec",
        "address": {
            "building": "XSq8nBR6p5h",
            "coord": [
                0.24571618383621482,
                12.282523956489666
            ],
            "street": "ifYx6Cd9AARj6k",
            "zipcode": "0dSfHgwqOmeNQvwp"
        },
        "borough": "2KPsVKBoCM",
        "cuisine": "boj06dZNP1pXe4IVtr",
        "grades": [
            {
                "date": 68,
                "grade": "9x64"
            },
            {
                "date": 78,
                "grade": "BKHPvpEudL"
            },
        ],
        "name": "PncQVI0",
        "restaurant_id": "3zBn2SM4w7t"
    }
]
```

Congratulation! we have finished generate test data.

### create configuration file by hand

Now, let's try to create a configuration file that got the same data above. We will create as `.toml` because it is easier to read.

#### 1. init file

Let's name the configuration file `config.ss.toml` then open it with your favorite editor.
Add the following line

``` toml
[[databases]]
d_name = "test"

# collection: student
[databases.collection]
c_name = "restaurants"
count = 3
```

this is equavalent to the generated configuration

``` yaml
databases:
    - d_name: test
      collection:
        c_name: restaurants
        count: 3 
```

**d_name** and **c_name** is used to specific the name of database and collection respectively.

If we try to execute `$ seestern generate -f config.ss.toml -v`, it will has output like this

``` text
// database test collection restaurants
[
    {},
    {},
    {}
]
```

#### 2. add simple data type fields

Let's recap that our data will have field like this

- **_id** as id
- **address** as object that has field
  - **building** as string
  - **coord** as array contains (can be empty) double
  - **street** as string
  - **zipcode** as string
- **borough** as string
- **cuisine** as string
- **grades** as array contains (can be empty) objects that has field
  - **date** as integer
  - **grade** as string
  - **score** as integer (can be left off)
- **name** as string
- **restaurant_id** as string

So call simple data type for this tool is `id`, `string`, `double` and `integer`. To add a field we need a name of field and its data type.
For **__id** field, the configuration file can be edited like this

``` toml
....
count = 3

[[databases.collection.fields]]
constraints = [{type = "objectID"}]
f_name = "_id"
```

Actually the data type is not `id` but `objectID` so we need to specific it "objectID" in **type** inside **constraints** field, If we try to execute `$ seestern generate -f config.ss.toml -v`, it will has output like this

``` text
[
    {
        "_id": "622dbb2c806512c9e86306b3"
    },
    {
        "_id": "622dbb2c806512c9e86306b4"
    },
    {
        "_id": "622dbb2c806512c9e86306b5"
    }
]
```

Other than **constraints** a field can also has **enum** and **value**, inside **constraints** can also has multiple **type**, please see **seesternsyntax** for more information

Now come back to finish what we start, we can simple add to generate field `borough`, `cuisine`, `name` and `restaurant_id` that have type `string`

``` toml
....
f_name = "_id"

[[databases.collection.fields]]
constraints = [{type = "string"}]
f_name = "borough"

[[databases.collection.fields]]
constraints = [{type = "string"}]
f_name = "cuisine"

[[databases.collection.fields]]
constraints = [{type = "string"}]
f_name = "name"

[[databases.collection.fields]]
constraints = [{type = "string"}]
f_name = "restaurant_id"
```

And try to execute generate command will has output like this.

``` text
[
    {
        "_id": "622dbc6c13b2bcf89b071e4c",
        "borough": "cpdQ07NOr0o",
        "cuisine": "pWnIT5Btl2YjVn",
        "name": "V0wMWmYzu86XJ1X73l5",
        "restaurant_id": "9LVyFWDGTMEQ3"
    },
    {
        "_id": "622dbc6c13b2bcf89b071e4d",
        "borough": "szpE0",
        "cuisine": "aq08pDEO0y0PC5f",
        "name": "MGUp9vnbuf2Vrh0m",
        "restaurant_id": "iKwEWAkj8L0WGlkZHa0"
    },
    {
        "_id": "622dbc6c13b2bcf89b071e4e",
        "borough": "4w9SyrGGSETixUF4TA",
        "cuisine": "",
        "name": "S2z4J1xeZGe2MAfT6",
        "restaurant_id": "il1MtG"
    }
]
```

#### 3. add composite data type fields

Start with **address** field that has schema following schema

- **address** as object that has field
  - **building** as string
  - **coord** as array contains (can be empty) double
  - **street** as string
  - **zipcode** as string

We will start with field **building**, **street** and **zipcode** inside **address** first, add the following code

``` toml
....
f_name = "restaurant_id"

[[databases.collection.fields]]
constraints = [
  {type = "object", fields = [
    {f_name = "building", constraints = [{type = "string"}]},
    {f_name = "street", constraints = [{type = "string"}]},
    {f_name = "zipcode", constraints = [{type = "string"}]},
  ]},
]
f_name = "address"
```

And try to execute generate command will has output like this (show only one document).

``` text
[
    {
        "_id": "622dbf8049d01692302da698",
        "address": {
            "building": "g",
            "street": "Gh76a2srxt3y9",
            "zipcode": ""
        },
        "borough": "IOI7cXeM0F",
        "cuisine": "UbNFiOdk",
        "name": "zuKlX2ft6tdy",
        "restaurant_id": "3Tvq0ixb820aP3"
    }
    ...
]
```

Next, for **coord** field that is `array` with `double`. this can be easily done by add following code

``` toml
....
f_name = "restaurant_id"

[[databases.collection.fields]]
constraints = [
  {type = "object", fields = [
    {f_name = "building", constraints = [{type = "string"}]},
    {f_name = "street", constraints = [{type = "string"}]},
    {f_name = "zipcode", constraints = [{type = "string"}]},
    {f_name = "coord", constraints = [{type = "array", element_type = [{type = "double"}]}]},
  ]},
]
f_name = "address"
```

And try to execute generate command will has output like this (show only one document).

``` text
[
   {
        "_id": "622dc04c48fd124c2953c8b9",
        "address": {
            "building": "xlCxeE95RwrUfRqVGsw",
            "coord": [
                33.59234768839571
            ],
            "street": "9H",
            "zipcode": "LtPl"
        },
        "borough": "i22M8NMQqSZp",
        "cuisine": "evBm",
        "name": "eR",
        "restaurant_id": "6j80pt3Jx"
    }
    ...
]
```

For the last field, **grades**, is array of objects, it can be done by add following code.

``` toml
....
f_name = "address"

[[databases.collection.fields]]
f_name = "grades"

[[databases.collection.fields.constraints]]
element_type = [
{type = "object", fields = [
        {f_name = "date", constraints = [{type = "integer"}]},
        {f_name = "grade", constraints = [{type = "string"}]},
        {f_name = "score", constraints = [{type = "integer"}]}
    ]}
]
type = "array"
```

And try to execute generate command will has output like this (show only one document).

``` text
[
    {
        "_id": "622dc2cf1337e6b491005107",
        "address": {
            "building": "AoeOB2",
            "coord": [
                35.57844767799763,
                34.51697929110399,
                24.920308188431363
            ],
            "street": "1m",
            "zipcode": ""
        },
        "borough": "hmiwurC",
        "cuisine": "MSJJIzyk",
        "grades": [
            {
                "date": 77,
                "grade": "y3drMWVH1o74lcvh",
                "score": 36
            }
        ],
        "name": "YlUFvG7IXOzSi",
        "restaurant_id": "PRzNK"
    }
    ...
]
```

We nearly done the work, but one last thing is that **grades.score** can be omited, this can be done by add following code

``` toml
...
[[databases.collection.fields.constraints]]
element_type = [
{type = "object", fields = [
        {f_name = "date", constraints = [{type = "integer"}]},
        {f_name = "grade", constraints = [{type = "string"}]},
        {f_name = "score", constraints = [{type = "integer"}], omit_weight = 0.5} # add omit_weight
    ]}
]
type = "array"
```

When add `omit_weight` field, the value has to be between 0.0-1.0. It is represent the percentage of how likely that this field will be omitted. 1.0 is it always be omitted and 0.0 is not omitted at all (default). So, after edited the code and execute generate command the output for **grades** field will look like this

``` text
"grades": [
            {
                "date": 10,
                "grade": "nYmrBR0W"
            },
            {
                "date": 77,
                "grade": "PdL4ZWAJOt"
            },
            {
                "date": 61,
                "grade": "qNGCmu"
            },
            {
                "date": 21,
                "grade": "0gmANnL"
            },
            {
                "date": 39,
                "grade": "vAnMNdAN",
                "score": 43
            },
            {
                "date": 38,
                "grade": "qEHRB9ycu6WwQ"
            },
            {
                "date": 95,
                "grade": "HcrgMmkXppZFy"
            }
        ]
```

Look at the result, some has **score** field and some don't, so we already finish and got the same configuration file that is the same with the generated file.

Hooray! now we already finished the tutorial, please see `seesternsyntax` for more information about each fields and their arrtibutes.
