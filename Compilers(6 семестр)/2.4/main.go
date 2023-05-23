package main

import (
	"2.4/first"
	"2.4/syntaxAnalizer"
)

func main() {
	input := `E  ( T {("+","-") T} )
	T  ( F {("*","/") F} )
	F  ( "n", "-" F,
	     "(" E ")" )`

	parser := syntaxAnalizer.NewParser(input)
	tree := parser.Parse()
	first.GetFirst(tree)
}
