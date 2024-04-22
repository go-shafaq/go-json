# go-json

forked from github.com/goccy/go-json

Added future:

    casting names automatically if needed
#
    no more tagging every single field to specify its name

## Basic Usage

### Installation

```bash
go get github.com/go-shafaq/go-json
```


### Code

Here is a minimal example.

```go
package main

import (
	"fmt"
	"github.com/go-shafaq/defcase"
	json "github.com/go-shafaq/go-json"
)

func main() {
	// Get Public Default case
	dcase := defcase.Get()
	// Set a case for "json" tag and specific package
	// "*" means any package
	dcase.SetCase("json", "*", defcase.Snak_case)
	
	// Set a Default_Case to library
	json.SetDefCase(dcase)
	// Marshal a struct (which has no tags)
	bytes, _ := json.Marshal(Item{Id: 3, Name: "laptop", CategoryId: 11})

	fmt.Println(string(bytes))
}

type Item struct {
	Id         int `casting:"no more tags"`
	Name       string
	CategoryId int
}
```

### Terminal

Printed result

```json
{"id":3,"name":"laptop","category_id":11}
```

Pretty json

```json
{
  "id": 3,
  "name": "laptop",
  "category_id": 11
}
```
