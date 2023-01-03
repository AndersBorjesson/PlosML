package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"ploshml/ast"
	"ploshml/lexer"
	"ploshml/parser"
	"strings"
	// "github.com/dominikbraun/graph"
	// "github.com/dominikbraun/graph/draw"
	// "github.com/goccy/go-graphviz"
)

type Item struct {
	Alias      string
	Name       string
	FullName   string
	Attributes []string
	Typ        string
}

type Relation struct {
	From string
	To   string
	Typ  string
}

func add_item(r *[]Item, item ast.KeyItem) {
	var tmp Item
	tmp.Name = item.Name
	tmp.Attributes = item.Attribute
	*r = append(*r, tmp)
}
func add_relation(r *[]Relation, lhs, rhs ast.KeyItem, operator string) {
	var tmp Relation
	forward := true
	realoperator := operator
	switch operator {
	case "<-":
		forward = false
		realoperator = "->"
	case "-<":
		forward = false
		realoperator = ">-"
	default:
		forward = true
		realoperator = operator
	}
	if forward {
		tmp.From = lhs.Name
		tmp.To = rhs.Name

	} else {
		tmp.From = rhs.Name
		tmp.To = lhs.Name
	}
	tmp.Typ = realoperator
	*r = append(*r, tmp)
}
func derive_structure(raw []ast.ParserOut) ([]Relation, []Item) {
	var relations []Relation
	var items []Item
	for _, y := range raw {
		if y.Typ == "normal" {
			for _, ly := range y.Normal {
				for _, lhs := range ly.Lhs {
					add_item(&items, lhs)
					for _, rhs := range ly.Rhs {
						add_item(&items, rhs)
						add_relation(&relations, lhs, rhs, ly.Operator)
					}
				}

			}

		}
	}
	return relations, items
}
func check_file_exists(filename string) (bool, error) {

	_, err := os.Stat(filename)
	if err == nil {
		//fmt.Printf("File %s exists", filename)
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File %s does not exist", filename)
		return false, nil
	}
	return false, err
}
func ReadFile(filename string) ([]string, error) {
	existance, err := check_file_exists(filename)
	valid := existance && (err == nil)
	if valid != true {
		return nil, err
	}
	f, _ := os.Open(filename)
	defer f.Close()

	raw, _ := io.ReadAll(f)
	strslice := strings.Split(string(raw), "\n")
	return strslice, nil
}
func item_relations() (map[string]string, map[string]string) {
	from := map[string]string{
		"triggers":    "P",
		"handles":     "O",
		"relates":     "any",
		"requires":    "P",
		"aggregates":  "any",
		"implements":  "any",
		"exhibits":    "P",
		"generalizes": "any",
		"->":          "P",
		">-":          "O",
		"-<":          "P",
		"<-":          "O"}
	to := map[string]string{
		"triggers":    "P",
		"handles":     "P",
		"relates":     "any",
		"requires":    "O",
		"aggregates":  "any",
		"implements":  "any",
		"exhibits":    "O",
		"generalizes": "any",
		"->":          "O",
		">-":          "P",
		"-<":          "O",
		"<-":          "P"}
	return from, to
}
func derive_types(relations []Relation, items *[]Item) {
	itemtypes := make(map[string][]string)
	fromtype, totype := item_relations()
	for _, l1 := range relations {
		itemtypes[l1.From] = append(itemtypes[l1.From], fromtype[l1.Typ])
		itemtypes[l1.To] = append(itemtypes[l1.To], totype[l1.Typ])

		// fmt.Println(l1.Typ)

	}
	unknowns := false
	for l1, val := range itemtypes {
		unknown := allSameStrings(val) == false
		if unknown {
			fmt.Printf("Type of item %s is ambiguous", l1)
			unknowns = true
		}
	}
	if unknowns {
		os.Exit(2)
	}
	for l1, _ := range *items {

		(*items)[l1].Typ = itemtypes[(*items)[l1].Name][0]
	}

}
func allSameStrings(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
func main() {
	filename := "ploshml_input_test.opa"
	filecontent, _ := ReadFile(filename)
	for _, l1 := range filecontent {
		fmt.Println(l1)
	}

	var sum_yield []ast.ParserOut
	var errors []error
	for _, l1 := range filecontent {
		lex := lexer.New([]rune(l1))

		yield, err := parser.New(lex).Parse()
		fmt.Println(yield)
		if err != nil {
			errors = append(errors, err)
		} else {
			sum_yield = append(sum_yield, yield.(ast.ParserOut))
		}
	}
	// fmt.Println(sum_yield)
	relations, items := derive_structure(sum_yield)
	derive_types(relations, &items)
	makegv(relations, items)
	fmt.Println(items)
	//// makegraphviz()
}
