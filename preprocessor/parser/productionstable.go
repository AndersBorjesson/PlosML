package parser

import (
	"ploshml/preprocessor/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]interface{}) (interface{}, error)
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String:     `G0 : Expr ;`,
		Id:         "G0",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.G00(X[0])
		},
	},
	ProdTabEntry{
		String:     `Expr : Start filename Stop ;`,
		Id:         "Expr",
		NTType:     0,
		Index:      1,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Expr0(X[0], X[1], X[2])
		},
	},
	ProdTabEntry{
		String:     `Start : \include{ ;`,
		Id:         "Start",
		NTType:     1,
		Index:      2,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Start0(X[0])
		},
	},
	ProdTabEntry{
		String:     `Stop : } ;`,
		Id:         "Stop",
		NTType:     2,
		Index:      3,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
			return ast.Stop0(X[0])
		},
	},
}
