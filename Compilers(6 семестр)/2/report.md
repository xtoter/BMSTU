% Лабораторная работа № 2.1. Синтаксические деревья
% 28 февраля 2023 г.
% Терюха Михаил, ИУ9-62Б

# Цель работы
Целью данной работы является изучение представления синтаксических деревьев в памяти компилятора и приобретение навыков преобразования синтаксических деревьев.
# Индивидуальный вариант
Заменить вхождения глобальной константы FUNCNAME на имя неанонимной функции, в которой константы упоминается, либо на "(global)", если константа упоминается в глобальном контексте. В исходной программе константа FUNCNAME должна быть объявлена в глобальной области видимости типа string (в противном случае следует выдать ошибку), значение константы может быть произвольным, в новой программе она должна отсутстовать.

# Реализация

Демонстрационная программа:

```go
package main

import "fmt"

var FUNCNAME string = "main"

func testfunc() {
	fmt.Println(FUNCNAME)
}
func main() {
	testfunc()
	fmt.Println(FUNCNAME)
}
```

Программа, осуществляющая преобразование синтаксического дерева:

```go
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
```

# Тестирование

Результат трансформации демонстрационной программы:

```go
package main

import "fmt"

var (global) string = "main"

func testfunc() {
	fmt.Println(testfunc)
}
func main() {
	testfunc()
	fmt.Println(main)
}
```


# Вывод
Я ознакомился с AST деревьями в golang, приобрел навыки преобразования синтаксических деревьев и реализовал индивидуальный вариант задания.
