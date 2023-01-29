package language_parser

import (
	"fmt"
	"ploshml/language_parser/ast"
	"ploshml/language_parser/lexer"
	"ploshml/language_parser/parser"
	"ploshml/preprocessor"
)

func DoParse(filecontent []preprocessor.LineType) ([]ast.ParserOut, []error) {
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
	return sum_yield, errors
}
