package json

import "github.com/go-shafaq/go-json/internal/defcase"

func DefCase(f func(tag string) func(pkgPath, name string) string) {
	defcase.Fn = f("json")
}
