package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"ploshml/language_parser"
	"ploshml/preprocessor"
	"ploshml/productions"
	"ploshml/semanticresolver"
	"strings"
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

func check_file_exists(filename string) (bool, error) {

	_, err := os.Stat(filename)
	if err == nil {
		//fmt.Printf("File %s exists", filename)
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		fmt.Printf("File/path %s does not exist\n", filename)
		return false, err
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

type Input struct {
	Content  string
	LineNo   int
	Indent   int
	FileName string
}

func main() {
	//
	// ASW()
	// os.Exit(2)
	in_flags := parse_flags()
	// fmt.Println(in_flags)
	// os.Exit(2)

	// about()

	filename := *in_flags.infile

	// Preprocessing
	preproc := preprocessor.Preprocessor{Filename: filename}
	preproc.Textstack = &[]preprocessor.LineType{}
	preproc.Errorstack = &[]preprocessor.ErrorType{}
	preproc.Do()
	preproc.IndentCheck()
	preproc.CleanEmpty()
	preproc.Check_ownership()
	filecontent := *preproc.Textstack
	fmt.Println("Preprocessing successful")

	// Language parsing
	sum_yield, _ := language_parser.DoParse(filecontent)

	graph_representation := semanticresolver.Resolve(&sum_yield, &filecontent)
	// // Language AST handling
	// var s Stack
	// s.Expected_indent = 0
	// s.OwnerID = 0
	// s.Entry_indent = -1
	// s.Structurals = &map[string][]ast.ParserOut{}
	// s.StrucInd2String = &map[int]string{}
	// s.Relations = &[]Relation{}
	// s.Items = &[]Item{}
	// s.OwnerList = []string{}
	// s.Name = "BaseContent"
	// s.indent_check(0, &sum_yield, &filecontent)
	// s.derive_structurals(0, &sum_yield, &filecontent)

	// //s.FixResidual(&sum_yield, &filecontent)
	// s.ExpandAll()
	// derive_types(*s.Relations, s.Items)
	// graph_representation := makegraph(s.Relations, s.Items)
	// Presentation

	if *in_flags.makegv {
		fmt.Println("Starting production of structural image")
		productions.Makegv(graph_representation, *in_flags.outpath, *in_flags.projectname)
		fmt.Println("Structural image production successful")
	}
	productions.Produce(graph_representation, *in_flags.outpath, *in_flags.projectname)
	// ResolveStructure(graph_representation)
}
