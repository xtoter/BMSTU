package main

import "github.com/skorobogatov/input"

type Tag int

const (
	ERROR  Tag = 1 << iota // Неправильная лексема
	NUMBER                 // Целое число
	VAR                    // Имя переменной
	PLUS                   // Знак +
	MINUS                  // Знак -
	MUL                    // Знак *
	DIV                    // Знак /
	LPAREN                 // Левая круглая скобка
	RPAREN                 // Правая круглая скобка
)

type Lexem struct {
	Tag
	Image string
}

func lexer(expr string, lexems chan Lexem) {
	//...
}
func main() {
	var lexems chan Lexem
	a := input.Gets()
	input.Scanf("\n")
	lexer(a, lexems)
}
