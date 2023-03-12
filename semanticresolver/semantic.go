package semanticresolver

import (
	"fmt"
	"os"
	"ploshml/language_parser/ast"
	"ploshml/preprocessor"
	"strings"

	"gonum.org/v1/gonum/graph"
)

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
		os.Exit(1)

	}
}

func (s *Stack) derive_structurals(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	(*s.StrucInd2String)[0] = "BaseContent"
	errs := []string{}
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

						(*s.Structurals)[name] = append((*s.Structurals)[name], tmp2)
					}
				} else {

					errstring := fmt.Sprintf("%d , %d : Ownership or indentation error \n", (*pp)[i].Filename, (*pp)[i].Lineno)
					errs = append(errs, errstring)
				}

			}
		} else {
			ind := (*s.StrucInd2String)[(*pp)[i].IndentID]
			(*s.Structurals)[ind] = append((*s.Structurals)[ind], (*yield)[i])
		}

	}
	if len(errs) > 0 {
		for _, l1 := range errs {
			fmt.Print(l1)
		}
		os.Exit(1)
	}
}
func (s *Stack) derive_structurals_old(istart int, yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	i_max := len(*yield)
	errs := []string{}
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
				errstring := fmt.Sprintf("%d , %d : Ownership or indentation error \n", (*pp)[i].Filename, (*pp)[i].Lineno)
				errs = append(errs, errstring)
			}

		} else {
			name := s.Name

			if name != "" {

				if (*pp)[i].IndentID == s.OwnerID {
					(*pp)[i].Used = true
					(*s.Structurals)[name] = append((*s.Structurals)[name], (*yield)[i])
				} else if (*pp)[i].IndentID != s.OwnerID {

					break
				} else {
					errstring := fmt.Sprintf("%d , %d : Syntax error \n", (*pp)[i].Filename, (*pp)[i].Lineno)
					errs = append(errs, errstring)
				}
			}

		}
	}
	if len(errs) > 0 {
		for _, l1 := range errs {
			fmt.Print(l1)
		}
		os.Exit(1)
	}
}

func (s *Stack) FixResidual(yield *[]ast.ParserOut, pp *[]preprocessor.LineType) {
	expected := -1
	errs := []string{}
	i_max := len(*yield)
	for i := 0; i < i_max; i++ {
		// fmt.Println((*pp)[i], expected)
		if (*pp)[i].Used == false {
			if expected == -1 {
				expected = (*pp)[i].OwnerID
			} else {
				if (*pp)[i].IndentID != expected {
					errstring := fmt.Sprintf("%d , %d : Unexpected indentation \n", (*pp)[i].Filename, (*pp)[i].Lineno)
					errs = append(errs, errstring)

				}
			}
		}
	}
	for i := 0; i < i_max; i++ {
		if (*pp)[i].Used == false {
			(*s.Structurals)["BaseContent"] = append((*s.Structurals)["BaseContent"], (*yield)[i])
		}
	}
	if len(errs) > 0 {
		for _, l1 := range errs {
			fmt.Print(l1)
		}
		os.Exit(1)
	}
}

func (s *Stack) ExpandAll() {
	s.ExpandStructure("BaseContent")
	// for _, l1 := range *s.Structurals {
	// 	if l1[0].Typ == "structural" {
	// 		fmt.Println(l1[0].Typ, l1[0].Structural.Name)
	// 		// if l1[0].Structural.Operator == "zoomin" {
	// 		// 	s.ExpandStructure(l1[0].Structural.Name)
	// 		// }
	// 	}
	// }
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

func Resolve(sum_yield *[]ast.ParserOut, filecontent *[]preprocessor.LineType) graph.Directed {
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
	s.indent_check(0, sum_yield, filecontent)
	s.derive_structurals(0, sum_yield, filecontent)

	//s.FixResidual(&sum_yield, &filecontent)
	s.ExpandAll()
	derive_types(*s.Relations, s.Items)
	graph_representation := makegraph(s.Relations, s.Items)
	fmt.Println("Semantic analysis successful")
	return graph_representation
}
