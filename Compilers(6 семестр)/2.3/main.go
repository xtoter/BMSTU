package main

import (
	"2.3/lexer"
	"2.3/parser"
	"fmt"
)

func main() {
	tokens := lexer.GetTokens()
	fmt.Println(tokens)
	parser.TopDownParse(tokens).Print("", nil)

}
