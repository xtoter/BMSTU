package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"reflect"
)

var stack []ast.Node

func replace(file *ast.File) {
	was := false
	ast.Inspect(file, func(node ast.Node) bool {

		if node, ok := node.(*ast.Ident); ok {
			if node.Name == "FUNCNAME" {
				if reflect.TypeOf(node.Name) != reflect.TypeOf("") {
					fmt.Println("error")
				}
				replacestring := "(global)"
				for i := len(stack) - 1; i >= 0; i-- {
					if k, ok := stack[i].(*ast.FuncDecl); ok {
						replacestring = k.Name.String()
					}
				}
				if !was && replacestring == "(global)" {
					was = true
				}
				node.Name = replacestring
			}
		}

		if node == nil {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, node)
		}
		return true
	})
	if !was {
		err := errors.New("Not found var")
		if err != nil {
			fmt.Println("An error occurred:", err)
			os.Exit(1)
		}

	}
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	fset := token.NewFileSet()
	if file, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments); err == nil {
		replace(file)

		if format.Node(os.Stdout, fset, file) != nil {
			fmt.Printf("Formatter error: %v\n", err)
		}
	} else {
		fmt.Printf("Errors in %s\n", os.Args[1])
	}
}
