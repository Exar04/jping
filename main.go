package main

import (
	"fmt"
	"jping/lexer"
	"jping/parser"
)

type lak struct {
	Name string
	Age  int
}

func main() {
	data := `{
		"Name":"cats are cute",
		"Age": 123423,
		"IsAlive": true,
		"ScoreInt" : [1,2,3],
		"ScoreBool" : [false, false, true],
		"Scorestr" : ["ma", "sco", "sec"],
		}

	 `
	lex := lexer.New(data)
	par := parser.New(lex)

	vals := par.ParseJson()

	for k, v := range vals {
		fmt.Printf("%v : %v\n", k, v)
	}

	// var m lak
	// m.Name = "yash"
	// m.Age = 2

	// val := reflect.ValueOf(m)
	// typ := reflect.TypeOf(m)

	// for i := 0; i < val.NumField(); i++ {
	// 	field := typ.Field(i).Name
	// 	value := val.Field(i).Interface()
	// 	fmt.Printf("Field: %s, Value: %v\n", field, value)
	// }
}

// tok := lex.NextToken()
// for tok.Type != token.EOF {
// 	fmt.Println(tok.Type, "  ", tok.Literal)
// 	tok = lex.NextToken()
// }
