package main

import (
	"fmt"
	"github.com/go-shafaq/go-json"
)

func main() {
	marshal, err := json.Marshal(u)

	fmt.Println(string(marshal), err)

	uzr := new(User)
	err = json.Unmarshal(marshal, uzr)
	if err != nil {
		panic(err)
	}

	fmt.Println(uzr)
}

var u = &User{13, 56}

type User struct {
	Id       int `json:"i_D"`
	ParentId int
}

func init() {
	json.DefCase(func(tag string) func(pkgPath string, name string) string {
		return func(pkgPath string, name string) string {
			println(pkgPath, name)
			return name + "_works"
		}
	})
}
