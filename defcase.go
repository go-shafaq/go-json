package json

import "github.com/go-shafaq/go-json/internal/defcase"

type DefCase = defcase.DefCase

func SetDefCase(dcase DefCase) {
	defcase.SetDefCase(dcase)
}
