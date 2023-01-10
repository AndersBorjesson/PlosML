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
}

func (s *Stack) derive_structurals(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	i_max := len(*yield)
	for i := istart; i < i_max; i++ {
		fmt.Println(i, istart, i_max)
		if (*yield)[i].Typ == "structural" {
			if (*pp)[i].PossibleOwner {
				var tmp Stack
				name := (*yield)[i].Structural.Name
				(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				tmp.Structurals = s.Structurals
				tmp.Expected_indent = (*pp)[i+1].Owner

				tmp.Name = name
				(*pp)[i].Used = true
				tmp.derive_structurals(i+1, yield, pp)
			} else {
				fmt.Println("Report error here, should not happen")
			}

		} else {
			name := s.Name

			if name != "" {
				if (*pp)[i].Owner == s.Expected_indent {
					(*pp)[i].Used = true
					(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
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
func derive_structurals_bc(yield []ast.ParserOut, pp []preprocessor.LineType) {
	structurals := make(map[string][]ast.ParserOut)
	//var s Stack
	//s.Structurals = &structurals
	// s.derive_structurals(sum_yield, filecontent)
	// var outyield []ast.ParserOut
	var inds []int
	used := []bool{}
	for i, l1 := range yield {
		used = append(used, false)
		if l1.Typ == "structural" {
			if pp[i].PossibleOwner {
				inds = append(inds, i)
			} else {
				fmt.Println("Structural with wrong indentation, raise error here")
			}
		}
	}
	for _, l1 := range inds {
		broken := true
		iadd := 0

		name := yield[l1].Structural.Name
		expected_label := pp[l1+1].Owner
		structurals[name] = append(structurals[name], yield[l1])
		used[l1+iadd] = true
		for broken {
			iadd += 1
			cond1 := pp[l1+iadd].Owner == expected_label
			cond2 := used[l1+iadd] == false
			if cond1 && cond2 {
				used[l1+iadd] = true
				structurals[name] = append(structurals[name], yield[l1+iadd])
			} else {
				broken = false
				fmt.Println(l1, name, used)
			}
		}

	}
	// ind := -1
	// for _, l1 := range inds {
	// 	if
	// }

}
func ExpandStructure(sin map[string][]ast.ParserOut) {
	// yield := []ast.Outdata{}
	for _, l1 := range sin["BaseContent"] {
		if l1.Typ == "normal" {
			for _, l2 := range l1.Normal {
				fmt.Println(l2)
			}
		}
	}
}
func main() {
	build_graph()

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
	s.Structurals = &map[string][]ast.ParserOut{}
	s.derive_structurals(0, &sum_yield, &filecontent)
	s.FixResidual(&sum_yield, &filecontent)
	ExpandStructure(*s.Structurals)
	os.Exit(2)
	// derive_structurals(sum_yield, filecontent)
	relations, items := derive_structure(sum_yield)
	derive_types(relations, &items)
	makegv(relations, items)
	fmt.Println(items)
	//// makegraphviz()
}
