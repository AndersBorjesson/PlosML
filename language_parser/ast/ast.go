// Generated by GoGLL.
package ast

import (
	"errors"
	"fmt"
	"ploshml/language_parser/token"
)

type KeyItem struct {
	Name      string
	Attribute []string
}
type Outdata struct {
	Lhs      []KeyItem
	Rhs      []KeyItem
	Operator string
}

type ParserOut struct {
	Typ        string
	Normal     []Outdata
	Structural StructuralOut
	Call       CallOut
	Owner      []string
}

type StructuralOut struct {
	Operator string
	Name     string
}

type CallOut struct {
	Caller string
	Calle  string
}

// G0 : EndExpr ;
func G00(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.G00 is unimplemented")
	return nil, nil
}

// EndExpr : StructuralExpr ;
func EndExpr0(p0 interface{}) (interface{}, error) {
	var yield ParserOut
	yield.Typ = "structural"
	yield.Structural = p0.(StructuralOut)
	return yield, nil
}

// EndExpr : key_component Alias key_component ;
func EndExpr1(p0, p1, p2 interface{}) (interface{}, error) {
	fmt.Println("ast.EndExpr1 is unimplemented")
	return nil, nil
}

// EndExpr : key_component Attribute key_component ;
func EndExpr2(p0, p1, p2 interface{}) (interface{}, error) {
	fmt.Println("ast.EndExpr2 is unimplemented")
	return nil, nil
}

// EndExpr : key_component Attribute KeyExpr ;
func EndExpr3(p0, p1, p2 interface{}) (interface{}, error) {
	fmt.Println("ast.EndExpr3 is unimplemented")
	return nil, nil
}

// EndExpr : Call key_component Alias key_component ;
func EndExpr4(p0, p1, p2, p3 interface{}) (interface{}, error) {
	fmt.Println("ast.EndExpr4 is unimplemented")
	fmt.Println(p1.(*token.Token).LiteralString())
	Call := CallOut{Calle: p1.(*token.Token).LiteralString(),
		Caller: p3.(*token.Token).LiteralString()}
	yield := ParserOut{Typ: "call",
		Call: Call}
	// os.Exit(3)
	return yield, nil
}

// EndExpr : NormalExpr ;
func EndExpr5(p0 interface{}) (interface{}, error) {
	var yield ParserOut
	yield.Typ = "normal"
	yield.Normal = p0.([]Outdata)
	return yield, nil
}

// EndExpr : PragmaExpr ;
func EndExpr6(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.EndExpr6 is unimplemented")
	return nil, nil
}

// StructuralExpr : Structural key_component ;
func StructuralExpr0(p0, p1 interface{}) (interface{}, error) {

	var yield StructuralOut
	yield.Operator = p0.(string)
	yield.Name = p1.(*token.Token).LiteralString()
	return yield, nil
}

// NormalExpr : KeyExpr Dualsidedkey key_component ;
func NormalExpr0(p0, p1, p2 interface{}) (interface{}, error) {
	var yield []Outdata
	var tmp Outdata
	tmp.Lhs = p0.([]KeyItem)
	tmp.Operator = (p1.(string))
	var tmp2 KeyItem
	tmp2.Name = p2.(*token.Token).LiteralString()
	tmp.Rhs = append(tmp.Rhs, tmp2)
	yield = append(yield, tmp)
	return yield, nil
}

// NormalExpr : key_component Dualsidedkey key_component ;
func NormalExpr1(p0, p1, p2 interface{}) (interface{}, error) {
	var yield []Outdata
	var tmp Outdata
	var tmp2 KeyItem
	tmp2.Name = p0.(*token.Token).LiteralString()
	tmp.Lhs = append(tmp.Lhs, tmp2)
	tmp.Operator = (p1.(string))
	var tmp3 KeyItem
	tmp3.Name = p2.(*token.Token).LiteralString()
	tmp.Rhs = append(tmp.Rhs, tmp3)
	yield = append(yield, tmp)
	return yield, nil
}

// NormalExpr : KeyExpr Dualsidedkey KeyExpr ;
func NormalExpr2(p0, p1, p2 interface{}) (interface{}, error) {
	var yield []Outdata
	var tmp Outdata
	tmp.Lhs = p0.([]KeyItem)
	tmp.Operator = (p1.(string))
	tmp.Rhs = p2.([]KeyItem)
	yield = append(yield, tmp)
	return yield, nil
}

// NormalExpr : key_component Dualsidedkey KeyExpr ;
func NormalExpr3(p0, p1, p2 interface{}) (interface{}, error) {
	var yield []Outdata
	var tmp Outdata
	tmp.Rhs = p2.([]KeyItem)
	tmp.Operator = (p1.(string))
	var tmp3 KeyItem
	tmp3.Name = p0.(*token.Token).LiteralString()
	tmp.Lhs = append(tmp.Lhs, tmp3)
	yield = append(yield, tmp)
	return yield, nil
}

// NormalExpr : NormalExpr Dualsidedkey KeyExpr ;
func NormalExpr4(p0, p1, p2 interface{}) (interface{}, error) {
	yield := p0.([]Outdata)
	var tmp Outdata
	tmp.Rhs = p2.([]KeyItem)
	tmp.Operator = (p1.(string))
	tmp.Lhs = yield[len(yield)-1].Rhs
	yield = append(yield, tmp)
	return yield, nil
}

// NormalExpr : NormalExpr Dualsidedkey key_component ;
func NormalExpr5(p0, p1, p2 interface{}) (interface{}, error) {

	yield := p0.([]Outdata)
	var tmp Outdata
	tmp.Lhs = yield[len(yield)-1].Rhs
	tmp.Operator = (p1.(string))
	var tmp3 KeyItem
	tmp3.Name = p2.(*token.Token).LiteralString()
	tmp.Rhs = append(tmp.Rhs, tmp3)
	yield = append(yield, tmp)
	return yield, nil
}

func transform(p interface{}) KeyItem {
	var tmp KeyItem
	tmp.Name = p.(*token.Token).LiteralString()
	return tmp
}

// KeyExpr : key_component Separator key_component ;
func KeyExpr0(p0, p1, p2 interface{}) (interface{}, error) {
	yield := []KeyItem{transform(p0), transform(p2)}

	return yield, nil
}

// KeyExpr : KeyExpr Separator key_component ;
func KeyExpr1(p0, p1, p2 interface{}) (interface{}, error) {
	yield := p0.([]KeyItem)
	yield = append(yield, transform(p2))
	return yield, nil
}

// KeyExpr : key_component Property key_component ;
func KeyExpr2(p0, p1, p2 interface{}) (interface{}, error) {
	fmt.Println("ast.KeyExpr2 is unimplemented")
	return nil, nil
}

// KeyExpr : KeyExpr Property key_component ;
func KeyExpr3(p0, p1, p2 interface{}) (interface{}, error) {
	fmt.Println("ast.KeyExpr3 is unimplemented")
	return nil, nil
}

// PragmaExpr : Pragma key_component ;
func PragmaExpr0(p0, p1 interface{}) (interface{}, error) {
	fmt.Println("ast.PragmaExpr0 is unimplemented")
	return nil, nil
}

// Separator : , ;
func Separator0(p0 interface{}) (interface{}, error) {
	return nil, nil
}

// Dualsidedkey : triggers ;
func Dualsidedkey0(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : handles ;
func Dualsidedkey1(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : owns ;
func Dualsidedkey2(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : relates ;
func Dualsidedkey3(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : requires ;
func Dualsidedkey4(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : aggregates ;
func Dualsidedkey5(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : implements ;
func Dualsidedkey6(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : exhibits ;
func Dualsidedkey7(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : generalizes ;
func Dualsidedkey8(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : -> ;
func Dualsidedkey9(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : >- ;
func Dualsidedkey10(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : -< ;
func Dualsidedkey11(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Dualsidedkey : <- ;
func Dualsidedkey12(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Structural : zoomin ;
func Structural0(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Structural : subroutine ;
func Structural1(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Structural : class ;
func Structural2(p0 interface{}) (interface{}, error) {
	operator := (p0.(*token.Token).LiteralString())
	return operator, nil
}

// Pstart : { ;
func Pstart0(p0 interface{}) (interface{}, error) {
	err := errors.New("ast.Pstart0  ({) is unimplemented")
	return nil, err
}

// Pend : } ;
func Pend0(p0 interface{}) (interface{}, error) {
	err := errors.New("ast.Pend0  (}) is unimplemented")
	return nil, err
}

// Property : :: ;
func Property0(p0 interface{}) (interface{}, error) {
	err := errors.New("Property : :: is unimplemented")
	return nil, err
}

// Alias : as ;
func Alias0(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.Alias0 is unimplemented")
	return nil, nil
}

// Attribute : is ;
func Attribute0(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.Attribute0 is unimplemented")
	return nil, nil
}

// Pragma : #pragma ;
func Pragma0(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.Pragma0 is unimplemented")
	return nil, nil
}

// Call : call ;
func Call0(p0 interface{}) (interface{}, error) {
	fmt.Println("ast.Call0 is unimplemented")
	return nil, nil
}
