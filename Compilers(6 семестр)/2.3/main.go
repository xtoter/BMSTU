package main

import (
	"2.3/lexer"
	"fmt"
)

func main() {
	tokens := lexer.GetTokens()
	fmt.Println(tokens)
}
