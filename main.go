package main

import (
	"fmt"
	"jping/lexer"
	"jping/token"
)

type lak struct {
	Name string
}

func main() {
	data := `{"Name":"cats are cute"}`
	lex := lexer.New(data)

	tok := lex.NextToken()
	for tok.Type != token.EOF {
		fmt.Println(tok.Type, "  ", tok.Literal)
		tok = lex.NextToken()
	}
}
