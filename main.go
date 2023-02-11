package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"ploshml/language_parser"
	"ploshml/language_parser/ast"
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

func add_item(r *[]Item, item ast.KeyItem, name string) {
	var tmp Item
	tmp.Name = name + item.Name
	tmp.Attributes = item.Attribute
	*r = append(*r, tmp)
}
func add_relation(r *[]Relation, lhs, rhs ast.KeyItem, operator, name string) {
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
		tmp.From = name + lhs.Name
		tmp.To = name + rhs.Name

	} else {
		tmp.From = name + rhs.Name
		tmp.To = name + lhs.Name
	}
	tmp.Typ = realoperator
	*r = append(*r, tmp)
}

// func derive_structure(raw []ast.ParserOut) ([]Relation, []Item) {
// 	var relations []Relation
// 	var items []Item
// 	for _, y := range raw {
// 		if y.Typ == "normal" {
// 			for _, ly := range y.Normal {
// 				for _, lhs := range ly.Lhs {
// 					add_item(&items, lhs)
// 					for _, rhs := range ly.Rhs {
// 						add_item(&items, rhs)
// 						add_relation(&relations, lhs, rhs, ly.Operator)
// 					}
// 				}

// 			}

//			}
//		}
//		return relations, items
//	}
func check_file_exists(filename string) (bool, error) {

	_, err := os.Stat(filename)
	if err == nil {
		//fmt.Printf("File %s exists", filename)
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File %s does not exist\n", filename)
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
			fmt.Printf("Type of item %s is ambiguous\n", l1)
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
	StrucInd2String *map[int]string
	Entry_indent    int
	Expected_indent int
	IAdd            int
	Inds            []int
	Used            []bool
	Name            string
	Relations       *[]Relation
	Items           *[]Item
	OwnerList       []string
	OwnerID         int
}

func (s *Stack) indent_check(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	i_max := len(*yield)
	errs := []string{}
	for i := istart; i < i_max; i++ {
		if (*pp)[i].PossibleOwner == true {
			if (*yield)[i].Typ != "structural" {
				err := fmt.Sprintf("Indentation error in file %s  around line %d : %s", (*pp)[i].Filename, (*pp)[i].Lineno, (*pp)[i].Text)
				errs = append(errs, err)
			}
		}
		if (*pp)[i].PossibleOwner == false {
			if (*yield)[i].Typ == "structural" {
				err := fmt.Sprintf("Indentation error in file %s  around line %d : %s", (*pp)[i].Filename, (*pp)[i].Lineno, (*pp)[i].Text)
				errs = append(errs, err)
			}
		}
	}
	if len(errs) > 0 {
		for _, l1 := range errs {
			fmt.Println(l1)
		}

	}
}

func (s *Stack) derive_structurals(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	(*s.StrucInd2String)[0] = "BaseContent"
	i_max := len(*yield)
	for i := istart; i < i_max; i++ {
		if (*yield)[i].Typ == "structural" {

			if (*pp)[i].PossibleOwner {
				if (*s.StrucInd2String)[(*pp)[i].OwnerID] == "" {
					(*s.StrucInd2String)[(*pp)[i].OwnerID] = (*yield)[i].Structural.Name
					if (*yield)[i].Structural.Operator == "zoomin" {
						tmp := ast.CallOut{
							Caller: (*yield)[i].Structural.Name,
							Calle:  (*yield)[i].Structural.Name}

						tmp2 := ast.ParserOut{Typ: "call",
							Call: tmp,
						}
						id := (*pp)[i].IndentID
						name := (*s.StrucInd2String)[id]
						fmt.Println("The name is ", name)
						(*s.Structurals)[name] = append((*s.Structurals)[name], tmp2)
					}
				} else {
					fmt.Println("Error here")

				}

			}
		} else {
			ind := (*s.StrucInd2String)[(*pp)[i].IndentID]
			(*s.Structurals)[ind] = append((*s.Structurals)[ind], (*yield)[i])
		}
		fmt.Println((*pp)[i].PossibleOwner, (*pp)[i].OwnerID, (*pp)[i].Owner, (*pp)[i].IndentID, (*pp)[i].Text)
	}
}
func (s *Stack) derive_structurals_old(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	i_max := len(*yield)
	for i := istart; i < i_max; i++ {
		// fmt.Println(s.Expected_indent, i, istart, i_max, (*yield)[i].Typ)
		if (*yield)[i].Typ == "structural" {

			if (*pp)[i].PossibleOwner {
				// (*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				if (*yield)[i].Structural.Operator == "zoomin" {
					tmp := ast.CallOut{
						Caller: (*yield)[i].Structural.Name,
						Calle:  (*yield)[i].Structural.Name}

					tmp2 := ast.ParserOut{Typ: "call",
						Call: tmp,
					}
					name := s.Name
					// fmt.Println("The name is ", name)
					(*s.Structurals)[name] = append((*s.Structurals)[name], tmp2)
				}
				var tmp Stack
				name := (*yield)[i].Structural.Name
				(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				// fmt.Println("NAME : ", name, (*pp)[i].Owner, s.Expected_indent)
				tmp.Structurals = s.Structurals
				tmp.Expected_indent = (*pp)[i+1].IndentID
				tmp.Entry_indent = (*pp)[i].Owner
				tmp.OwnerID = (*pp)[i].OwnerID
				tmp.Name = name
				(*pp)[i].Used = true
				tmp.derive_structurals(i+1, yield, pp)
			} else {
				fmt.Println("Report error here, should not happen")
			}

		} else {
			name := s.Name

			if name != "" {
				fmt.Println("HERE", (*pp)[i])
				if (*pp)[i].IndentID == s.OwnerID {
					(*pp)[i].Used = true
					(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				} else if (*pp)[i].IndentID != s.OwnerID {
					fmt.Println("Breaking ", (*pp)[i], s.OwnerID)
					break
				} else {
					fmt.Println("Handle error here, this should not happen unless syntax error")
					fmt.Println(name)
					fmt.Println((*pp)[i].Owner, s.Expected_indent, (*pp)[i].IndentID)
					fmt.Println((*pp)[i].PossibleOwner)
					fmt.Println((*pp)[i])
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
		// fmt.Println((*pp)[i], expected)
		if (*pp)[i].Used == false {
			if expected == -1 {
				expected = (*pp)[i].OwnerID
			} else {
				if (*pp)[i].IndentID != expected {
					fmt.Println("Error here, unexpected indent")
					fmt.Println((*pp)[i-1], expected)
					fmt.Println((*pp)[i])
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
			// if l1[0].Structural.Operator == "zoomin" {
			// 	s.ExpandStructure(l1[0].Structural.Name)
			// }
		}
	}
}
func (s *Stack) ExpandStructure(entrykey string) {
	name := strings.Join(s.OwnerList, ".")
	if name != "" {
		name = name + "."
	}

	for _, y := range (*s.Structurals)[entrykey] {
		if y.Typ == "normal" {
			for _, ly := range y.Normal {
				for _, lhs := range ly.Lhs {
					add_item(s.Items, lhs, name)

					for _, rhs := range ly.Rhs {
						add_item(s.Items, rhs, name)
						// fmt.Println(entrykey, lhs, rhs, name)
						add_relation(s.Relations, lhs, rhs, ly.Operator, name)
					}
				}

			}

		} else if y.Typ == "call" {
			tmp := Stack{Structurals: s.Structurals,
				Relations: s.Relations,
				Items:     s.Items,
				OwnerList: append(s.OwnerList, y.Call.Caller)}

			tmp.ExpandStructure(y.Call.Calle)
		}
	}

}

// func do_parse(filecontent []preprocessor.LineType) ([]ast.ParserOut, []error) {
// 	var sum_yield []ast.ParserOut
// 	var errors []error
// 	for _, l1 := range filecontent {
// 		lex := lexer.New([]rune(l1.Text))

// 		yield, err := parser.New(lex).Parse()
// 		fmt.Println(yield)
// 		if err != nil {
// 			errors = append(errors, err)
// 		} else {
// 			sum_yield = append(sum_yield, yield.(ast.ParserOut))
// 		}
// 	}
// 	return sum_yield, errors
// }

func main() {
	//
	ASW()
	os.Exit(2)
	in_flags := parse_flags()
	// fmt.Println(in_flags)
	// os.Exit(2)
	if false {
		about()
	}
	filename := in_flags.infile
	// Preprocessing
	preproc := preprocessor.Preprocessor{Filename: filename}
	preproc.Textstack = &[]preprocessor.LineType{}
	preproc.Errorstack = &[]preprocessor.ErrorType{}
	preproc.Do()
	preproc.IndentCheck()
	preproc.CleanEmpty()
	preproc.Check_ownership()
	filecontent := *preproc.Textstack

	// Language parsing
	sum_yield, _ := language_parser.DoParse(filecontent)

	// Language AST handling
	var s Stack
	s.Expected_indent = 0
	s.OwnerID = 0
	s.Entry_indent = -1
	s.Structurals = &map[string][]ast.ParserOut{}
	s.StrucInd2String = &map[int]string{}
	s.Relations = &[]Relation{}
	s.Items = &[]Item{}
	s.OwnerList = []string{}
	s.Name = "BaseContent"
	s.indent_check(0, &sum_yield, &filecontent)
	s.derive_structurals(0, &sum_yield, &filecontent)

	//s.FixResidual(&sum_yield, &filecontent)
	s.ExpandAll()
	derive_types(*s.Relations, s.Items)
	graph_representation := makegraph(s.Relations, s.Items)
	// Presentation

	if in_flags.makegv {
		makegv(graph_representation)
	}
	// ResolveStructure(graph_representation)
}
