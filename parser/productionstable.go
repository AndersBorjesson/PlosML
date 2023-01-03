
package parser

import(
    "ploshml/ast"
)

type (
	//TODO: change type and variable names to be consistent with other tables
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index int
		NumSymbols int
		ReduceFunc func([]interface{}) (interface{}, error)
	}
)

var productionsTable = ProdTab {
	ProdTabEntry{
		String: `G0 : EndExpr ;`,
		Id: "G0",
		NTType: 0,
		Index: 0,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.G00(X[0])
		},
	},
	ProdTabEntry{
		String: `EndExpr : StructuralExpr ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 1,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr0(X[0])
		},
	},
	ProdTabEntry{
		String: `EndExpr : key_component Alias key_component ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 2,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr1(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `EndExpr : key_component Attribute key_component ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 3,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr2(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `EndExpr : key_component Attribute KeyExpr ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 4,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr3(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `EndExpr : Call key_component Alias key_component ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 5,
		NumSymbols: 4,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr4(X[0],X[1],X[2],X[3])
		},
	},
	ProdTabEntry{
		String: `EndExpr : NormalExpr ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 6,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr5(X[0])
		},
	},
	ProdTabEntry{
		String: `EndExpr : PragmaExpr ;`,
		Id: "EndExpr",
		NTType: 4,
		Index: 7,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.EndExpr6(X[0])
		},
	},
	ProdTabEntry{
		String: `StructuralExpr : Structural key_component ;`,
		Id: "StructuralExpr",
		NTType: 14,
		Index: 8,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.StructuralExpr0(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : KeyExpr Dualsidedkey key_component ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 9,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr0(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : key_component Dualsidedkey key_component ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 10,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr1(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : KeyExpr Dualsidedkey KeyExpr ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 11,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr2(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : key_component Dualsidedkey KeyExpr ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 12,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr3(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : NormalExpr Dualsidedkey KeyExpr ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 13,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr4(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `NormalExpr : NormalExpr Dualsidedkey key_component ;`,
		Id: "NormalExpr",
		NTType: 6,
		Index: 14,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.NormalExpr5(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `KeyExpr : key_component Separator key_component ;`,
		Id: "KeyExpr",
		NTType: 5,
		Index: 15,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.KeyExpr0(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `KeyExpr : KeyExpr Separator key_component ;`,
		Id: "KeyExpr",
		NTType: 5,
		Index: 16,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.KeyExpr1(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `KeyExpr : key_component Property key_component ;`,
		Id: "KeyExpr",
		NTType: 5,
		Index: 17,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.KeyExpr2(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `KeyExpr : KeyExpr Property key_component ;`,
		Id: "KeyExpr",
		NTType: 5,
		Index: 18,
		NumSymbols: 3,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.KeyExpr3(X[0],X[1],X[2])
		},
	},
	ProdTabEntry{
		String: `PragmaExpr : Pragma key_component ;`,
		Id: "PragmaExpr",
		NTType: 9,
		Index: 19,
		NumSymbols: 2,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.PragmaExpr0(X[0],X[1])
		},
	},
	ProdTabEntry{
		String: `Separator : , ;`,
		Id: "Separator",
		NTType: 12,
		Index: 20,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Separator0(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : triggers ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 21,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey0(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : handles ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 22,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey1(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : owns ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 23,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey2(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : relates ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 24,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey3(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : requires ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 25,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey4(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : aggregates ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 26,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey5(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : implements ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 27,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey6(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : exhibits ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 28,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey7(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : generalizes ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 29,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey8(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : -> ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 30,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey9(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : >- ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 31,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey10(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : -< ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 32,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey11(X[0])
		},
	},
	ProdTabEntry{
		String: `Dualsidedkey : <- ;`,
		Id: "Dualsidedkey",
		NTType: 3,
		Index: 33,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Dualsidedkey12(X[0])
		},
	},
	ProdTabEntry{
		String: `Structural : zoomin ;`,
		Id: "Structural",
		NTType: 13,
		Index: 34,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Structural0(X[0])
		},
	},
	ProdTabEntry{
		String: `Structural : subroutine ;`,
		Id: "Structural",
		NTType: 13,
		Index: 35,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Structural1(X[0])
		},
	},
	ProdTabEntry{
		String: `Structural : class ;`,
		Id: "Structural",
		NTType: 13,
		Index: 36,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Structural2(X[0])
		},
	},
	ProdTabEntry{
		String: `Pstart : { ;`,
		Id: "Pstart",
		NTType: 11,
		Index: 37,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Pstart0(X[0])
		},
	},
	ProdTabEntry{
		String: `Pend : } ;`,
		Id: "Pend",
		NTType: 7,
		Index: 38,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Pend0(X[0])
		},
	},
	ProdTabEntry{
		String: `Property : :: ;`,
		Id: "Property",
		NTType: 10,
		Index: 39,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Property0(X[0])
		},
	},
	ProdTabEntry{
		String: `Alias : as ;`,
		Id: "Alias",
		NTType: 0,
		Index: 40,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Alias0(X[0])
		},
	},
	ProdTabEntry{
		String: `Attribute : is ;`,
		Id: "Attribute",
		NTType: 1,
		Index: 41,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Attribute0(X[0])
		},
	},
	ProdTabEntry{
		String: `Pragma : #pragma ;`,
		Id: "Pragma",
		NTType: 8,
		Index: 42,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Pragma0(X[0])
		},
	},
	ProdTabEntry{
		String: `Call : call ;`,
		Id: "Call",
		NTType: 2,
		Index: 43,
		NumSymbols: 1,
		ReduceFunc: func(X []interface{}) (interface{}, error) {
            return ast.Call0(X[0])
		},
	},
	
}
