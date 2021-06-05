package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Tag int

const (
	ERROR  = 1 << iota // Неправильная лексема
	NUMBER             // Целое число
	VAR                // Имя переменной
	PLUS               // Знак +
	MINUS              // Знак -
	MUL                // Знак *
	DIV                // Знак /
	LPAREN             // Левая круглая скобка
	RPAREN             // Правая круглая скобка
	COMMA              // Знак ,
	EQUAL              // Знак =
	BEGIN
	END
)

type Vertex struct {
	exection    string
	lefts       map[string]bool
	rights      map[string]bool
	help        int
	used        bool
	visited     bool
	connections *helplify
}
type Lexem struct {
	Tag
	Image string
}
type helplify struct {
	val  int
	next *helplify
}

func getnumber(expr string, i int) (string, int) {
	n := len(expr)
	for j := i; j < n; j++ {
		if !(expr[j] >= '0' && expr[j] <= '9') {
			return expr[i:j], j
		}
	}
	return expr[i:n], n
}
func getvar(expr string, i int) (string, int) {
	n := len(expr)
	for j := i; j < n; j++ {
		if !((expr[j] >= '0' && expr[j] <= '9') || (expr[j] >= 'A' && expr[j] <= 'Z') || (expr[j] >= 'a' && expr[j] <= 'z')) {
			return expr[i:j], j
		}
	}
	return expr[i:n], n
}
func getexpression(expr string, i int) (Lexem, int) {
	var cur Lexem
	if i == len(expr) {
		cur.Tag = END
		cur.Image = ""
		return cur, i + 1
	}
	for expr[i] == ' ' {
		i++
	}
	if expr[i] == '(' {
		cur.Tag = LPAREN
		cur.Image = "("
		return cur, i + 1
	}
	if expr[i] == ')' {
		cur.Tag = RPAREN
		cur.Image = ")"
		return cur, i + 1
	}
	if expr[i] == '+' {
		cur.Tag = PLUS
		cur.Image = "+"
		return cur, i + 1
	}
	if expr[i] == '-' {
		cur.Tag = MINUS
		cur.Image = "-"
		return cur, i + 1
	}
	if expr[i] == ',' {
		cur.Tag = COMMA
		cur.Image = ","
		return cur, i + 1
	}
	if expr[i] == '=' {
		cur.Tag = EQUAL
		cur.Image = "="
		return cur, i + 1
	}
	if expr[i] == '*' {
		cur.Tag = MUL
		cur.Image = "*"
		return cur, i + 1
	}
	if expr[i] == '/' {
		cur.Tag = DIV
		cur.Image = "/"
		return cur, i + 1
	}
	if expr[i] <= '9' && expr[i] >= '0' {
		temp1, temp2 := getnumber(expr, i)
		cur.Tag = NUMBER
		cur.Image = temp1
		return cur, temp2
	}
	if (expr[i] >= 'A' && expr[i] <= 'Z') || (expr[i] >= 'a' && expr[i] <= 'z') {
		temp1, temp2 := getvar(expr, i)
		cur.Tag = VAR
		cur.Image = temp1
		return cur, temp2

	}
	cur.Tag = ERROR
	return cur, len(expr)
}
func lexer(expr string) []Lexem {
	i := 0
	var lexems []Lexem
	if i == 0 {
		var cur Lexem
		cur.Tag = BEGIN
		cur.Image = ""
		lexems = append(lexems, cur)
	}

	var temp Lexem
	for i <= len(expr) {
		temp, i = getexpression(expr, i)
		lexems = append(lexems, temp)

	}

	return lexems
}

func cond(i, j Tag, equal bool) bool {

	return i == ERROR ||
		((i&(BEGIN|COMMA|VAR|EQUAL) == 0) && !equal) ||
		(i == j && i != MINUS && i != LPAREN && i != RPAREN) ||
		((i&(LPAREN|MUL|DIV|PLUS|MINUS) != 0) && (j&(MUL|DIV|END) != 0)) ||
		((i&(MUL|DIV) != 0) && (j&(RPAREN|MUL|DIV|PLUS|MINUS) != 0)) ||
		((i&(EQUAL) != 0) && (j&(END|RPAREN|DIV|MUL) != 0)) ||
		((i&(BEGIN|LPAREN) != 0) && (j&(EQUAL) != 0)) ||
		((i&(LPAREN) != 0) && (j&(RPAREN) != 0)) ||
		((i&(RPAREN) != 0) && (j&(LPAREN|VAR|NUMBER) != 0))
}
func balansed(lexems []Lexem) bool {
	equal := 0
	comma := 0
	paren := 0
	for i := 0; i < len(lexems); i++ {

		if lexems[i].Tag == LPAREN {
			paren++
		}
		if lexems[i].Tag == RPAREN {
			paren--
		}
		if lexems[i].Tag == EQUAL {
			equal++
		}
		if lexems[i].Tag == COMMA {
			if equal == 0 {
				comma++
			} else {
				comma--
			}
		}
	}
	if equal != 1 {
		return false
	}
	if comma != 0 {
		return false
	}
	if paren != 0 {
		return false
	}
	return true
}
func correctly(exection string) (bool, []Lexem) {
	lexems := lexer(exection)
	if !balansed(lexems) {
		return false, lexems
	}
	//fmt.Println(lexems)
	equal := false
	for i := 0; i < len(lexems)-1; i++ {
		if lexems[i].Tag == EQUAL {
			equal = true
		}
		if cond(lexems[i].Tag, lexems[i+1].Tag, equal) {
			return false, lexems
		}
	}
	return true, lexems
}

func parse(exection string) (Vertex, bool) {
	var v Vertex
	n := 0
	cond, lexems := correctly(exection)
	if !cond {
		return v, false
	}
	v.exection = exection
	v.rights = make(map[string]bool)
	v.lefts = make(map[string]bool)
	equal := 0
	for lexems[equal].Tag != EQUAL {
		equal++
	}
	if equal == len(exection) {
		return v, false
	}
	for n < equal {
		cur := lexems[n]
		if cur.Tag == VAR {
			if v.lefts[cur.Image] {
				return v, false
			}
			v.lefts[cur.Image] = true

		}
		n++
	}
	for n < len(lexems) {
		cur := lexems[n]
		if cur.Tag == VAR {
			v.rights[cur.Image] = true

		}
		n++

	}
	return v, true
}
func dfs(graph []Vertex, v int, order []string, cycles bool) ([]string, bool) {
	graph[v].used = true
	temp := graph[v].connections
	for temp != nil {

		if !graph[temp.val].visited && graph[temp.val].used {
			cycles = true
		}
		if !graph[temp.val].used {
			order, cycles = dfs(graph, temp.val, order, cycles)
		}
		temp = temp.next
	}
	graph[v].visited = true
	order = append(order, graph[v].exection)
	return order, cycles
}
func toStrings(input []byte) []string {
	var tempstring []byte
	var strings []string
	for i := 0; i < len(input); i++ {
		if input[i] == 10 {
			strings = append(strings, string(tempstring))
			tempstring = tempstring[:0]
		} else {
			tempstring = append(tempstring, input[i])
		}
	}
	return strings
}
func main() {
	var graph []Vertex
	var order []string
	cycles := false
	bs, _ := ioutil.ReadAll(os.Stdin)
	strings := toStrings(bs)
	for i := 0; i < len(strings); i++ {
		v, correct := parse(strings[i])
		if !correct {
			fmt.Println("syntax error")
			return
		}
		graph = append(graph, v)
	}
	varaibles_mask, val := make(map[string]bool), make(map[string]int)
	for i := 0; i < len(graph); i++ {
		for varaible := range graph[i].lefts {
			if varaibles_mask[varaible] {
				fmt.Println("syntax error")
				return
			} else {
				varaibles_mask[varaible] = true
				val[varaible] = i
			}
		}
	}

	for i := 0; i < len(graph); i++ {
		graph[i].help = 0
		for varaible := range graph[i].rights {
			if !varaibles_mask[varaible] {
				fmt.Println("syntax error")
				return
			} else {
				if graph[i].connections != nil {
					help := graph[i].connections
					for help.next != nil {
						help = help.next
					}
					var t helplify
					t.val = val[varaible]
					help.next = &t
					help = help.next
				} else {
					var t helplify
					t.val = val[varaible]
					graph[i].connections = &t
				}
				graph[val[varaible]].help++
			}
		}
	}
	for i := 0; i < len(graph); i++ {
		if !graph[i].used && graph[i].help == 0 {
			order, cycles = dfs(graph, i, order, cycles)
		}
	}
	if cycles || (len(order) == 0) {
		fmt.Println("cycle")
		return
	} else {
		for i := 0; i < len(order); i++ {
			fmt.Println(order[i])
		}
	}

}

