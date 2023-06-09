package main

import (
	"2.4/first"
	"2.4/syntaxAnalizer"
)

func main() {
	input := `E  ( T {("+","-") T} )
	T  ( F {("*","/") F} )
	F  ( "n", "-" F,
	     "(" E ")" )
    A ( "x" )
    B ( A )
    C ( B, D )
    D ( E )
    E ( "y" )`
	parser := syntaxAnalizer.NewParser(input)

	tree := parser.Parse()
	//fmt.Println(tree.ToString())
	first.GetFirst(tree)
}
