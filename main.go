package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"ploshml/ast"
	"ploshml/lexer"
	"ploshml/parser"
	"ploshml/preprocessor"
	"strings"
	// "github.com/dominikbraun/graph"
	// "github.com/dominikbraun/graph/draw"
	// "github.com/goccy/go-graphviz"
)

func about() {
	fmt.Println("PLoSH ML ver 0.1")
	fmt.Println("\n\n\n")
	fmt.Println("'Plosh' :  Combination of splash and plop")
	fmt.Println("'Plosh' :  The practice of masturbation while in a swimming pool.")
	fmt.Println("'Plosh' :  So subtextually gay that it is basically canon but the straight people don not see it")
	fmt.Println("\n\n\n")
	fmt.Println("'PLoSH ML' :  The Peasant Logic Short Hand Modelling Language.")

}

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

type Input struct {
	Content  string
	LineNo   int
	Indent   int
	FileName string
}

type Stack struct {
	Structurals     *map[string][]ast.ParserOut
	Entry_indent    int
	Expected_indent int
	IAdd            int
	Inds            []int
	Used            []bool
	Name            string
	Relations       *[]Relation
	Items           *[]Item
}

func (s *Stack) derive_structurals(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	i_max := len(*yield)
	for i := istart; i < i_max; i++ {
		fmt.Println(s.Expected_indent, i, istart, i_max, (*yield)[i].Typ)
		if (*yield)[i].Typ == "structural" {

			if (*pp)[i].PossibleOwner {
				var tmp Stack
				name := (*yield)[i].Structural.Name
				(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				fmt.Println("NAME : ", name, (*pp)[i].Owner, s.Expected_indent)
				tmp.Structurals = s.Structurals
				tmp.Expected_indent = (*pp)[i+1].IndentID
				tmp.Entry_indent = (*pp)[i].Owner
				tmp.Name = name
				(*pp)[i].Used = true
				tmp.derive_structurals(i+1, yield, pp)
			} else {
				fmt.Println("Report error here, should not happen")
			}

		} else {
			name := s.Name

			if name != "" {
				if (*pp)[i].IndentID == s.Expected_indent {
					(*pp)[i].Used = true
					(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				} else if (*pp)[i].Owner != s.Expected_indent {
					break
				} else {
					fmt.Println("Handle error here, this should not happen unless syntax error")
					fmt.Println(name)
					fmt.Println((*pp)[i].Owner, s.Expected_indent)
					fmt.Println((*pp)[i].PossibleOwner)
					os.Exit(2)
				}
			}

		}
	}
}

func (s *Stack) FixResidual(yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	expected := -1
	i_max := len(*yield)
	for i := 0; i < i_max; i++ {

		if (*pp)[i].Used == false {
			if expected == -1 {
				expected = (*pp)[i].Owner
			} else {
				if (*pp)[i].Owner != expected {
					fmt.Println("Error here, unexpected indent")
				}
			}
		}
	}
	for i := 0; i < i_max; i++ {
		if (*pp)[i].Used == false {
			(*s.Structurals)["BaseContent"] = append((*s.Structurals)["BaseContent"], (*yield)[i])
		}
	}
}

func (s *Stack) ExpandAll() {
	s.ExpandStructure("BaseContent")
	for _, l1 := range *s.Structurals {
		if l1[0].Typ == "structural" {
			fmt.Println(l1[0].Typ, l1[0].Structural.Name)
			if l1[0].Structural.Operator == "zoomin" {
				s.ExpandStructure(l1[0].Structural.Name)
			}
		}
	}
}
func (s *Stack) ExpandStructure(entrykey string) {

	for _, y := range (*s.Structurals)[entrykey] {
		if y.Typ == "normal" {
			for _, ly := range y.Normal {
				for _, lhs := range ly.Lhs {
					add_item(s.Items, lhs)

					for _, rhs := range ly.Rhs {
						add_item(s.Items, rhs)
						fmt.Println(entrykey, lhs, rhs)
						add_relation(s.Relations, lhs, rhs, ly.Operator)
					}
				}

			}

		} else if y.Typ == "call" {
			tmp := Stack{Structurals: s.Structurals,
				Relations: s.Relations,
				Items:     s.Items}

			tmp.ExpandStructure(y.Call.Calle)
		}
	}

}
func main() {
	// build_graph()
	// os.Exit(2)
	if false {
		about()
	}
	filename := "ploshml_input_test.opa"
	preproc := preprocessor.Preprocessor{Filename: filename}
	preproc.Textstack = &[]preprocessor.LineType{}
	preproc.Errorstack = &[]preprocessor.ErrorType{}
	preproc.Do()
	preproc.IndentCheck()
	preproc.CleanEmpty()
	preproc.Check_ownership()

	// filecontent, _ := ReadFile(filename)
	filecontent := *preproc.Textstack
	for _, l1 := range filecontent {
		fmt.Println(l1)
	}

	var sum_yield []ast.ParserOut
	var errors []error
	for _, l1 := range filecontent {
		lex := lexer.New([]rune(l1.Text))

		yield, err := parser.New(lex).Parse()
		fmt.Println(yield)
		if err != nil {
			errors = append(errors, err)
		} else {
			sum_yield = append(sum_yield, yield.(ast.ParserOut))
		}
	}
	// fmt.Println(sum_yield)
	var s Stack
	s.Expected_indent = 0
	s.Entry_indent = -1
	s.Structurals = &map[string][]ast.ParserOut{}
	s.Relations = &[]Relation{}
	s.Items = &[]Item{}
	s.derive_structurals(0, &sum_yield, &filecontent)
	s.FixResidual(&sum_yield, &filecontent)
	s.ExpandAll()

	// os.Exit(2)
	// derive_structurals(sum_yield, filecontent)
	// relations, items := derive_structure(sum_yield)
	derive_types(*s.Relations, s.Items)
	makegv(*s.Relations, *s.Items)
	// fmt.Println(items)
	//// makegraphviz()
}
