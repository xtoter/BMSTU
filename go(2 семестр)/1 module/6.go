package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/skorobogatov/input"
)

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
func lexer(expr string, lexems chan Lexem) {
	i := 0
	var temp Lexem
	for i < len(expr) {
		temp, i = getexpression(expr, i)
		lexems <- temp
		//fmt.Println(temp, i)
	}
}
func E(lindex int, alllexems []Lexem, tokens []string) (int, []Lexem, []string) {
	lindex, alllexems, tokens = T(lindex, alllexems, tokens)
	lindex, alllexems, tokens = EE(lindex, alllexems, tokens)
	return lindex, alllexems, tokens
}
func T(lindex int, alllexems []Lexem, tokens []string) (int, []Lexem, []string) {
	lindex, alllexems, tokens = F(lindex, alllexems, tokens)
	lindex, alllexems, tokens = TT(lindex, alllexems, tokens)
	return lindex, alllexems, tokens
}

func F(lindex int, alllexems []Lexem, tokens []string) (int, []Lexem, []string) {
	if len(alllexems) > lindex {
		cur := alllexems[lindex]
		if cur.Tag&(NUMBER|VAR) != 0 {
			lindex++
			tokens = append(tokens, cur.Image)
		} else if cur.Tag&MINUS != 0 {
			lindex++
			tokens = append(tokens, "-1")
			lindex, alllexems, tokens = F(lindex, alllexems, tokens)
			tokens = append(tokens, "*")
		} else if cur.Tag&LPAREN != 0 {
			lindex++
			lindex, alllexems, tokens = E(lindex, alllexems, tokens)
			if len(alllexems) > lindex {
				cur := alllexems[lindex]
				lindex++
				if cur.Tag&RPAREN == 0 {
					fmt.Println("error")
					os.Exit(0)
				}
			} else {
				fmt.Println("error")
				os.Exit(0)
			}
		} else {
			fmt.Println("error")
			os.Exit(0)
		}
	} else {
		fmt.Println("error")
		os.Exit(0)
	}
	return lindex, alllexems, tokens
}
func TT(lindex int, alllexems []Lexem, tokens []string) (int, []Lexem, []string) {
	if len(alllexems) > lindex {
		cur := alllexems[lindex]
		if cur.Tag&(DIV|MUL) != 0 {
			lindex++
			lindex, alllexems, tokens = F(lindex, alllexems, tokens)
			tokens = append(tokens, cur.Image)
			lindex, alllexems, tokens = TT(lindex, alllexems, tokens)
		}
	}
	return lindex, alllexems, tokens
}
func EE(lindex int, alllexems []Lexem, tokens []string) (int, []Lexem, []string) {
	if len(alllexems) > lindex {
		cur := alllexems[lindex]
		if cur.Tag&(PLUS|MINUS) != 0 {
			lindex++
			lindex, alllexems, tokens = T(lindex, alllexems, tokens)
			tokens = append(tokens, cur.Image)
			lindex, alllexems, tokens = EE(lindex, alllexems, tokens)
		} else if (cur.Tag & (VAR | NUMBER)) != 0 {
			fmt.Println("error")
			os.Exit(0)
		}
	}
	return lindex, alllexems, tokens
}
func startparse(lindex int, alllexems []Lexem, tokens []string) (int, []string) {
	lindex, _, tokens = E(lindex, alllexems, tokens)
	return lindex, tokens
}

func getanswer(tokens []string, datanum map[string]int) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		if tokens[i][0] >= '0' && tokens[i][0] <= '9' || (tokens[i][0] == '-' && len(tokens[i]) > 1) {

			value, _ := strconv.ParseInt(tokens[i], 10, 0)
			stack = append(stack, int(value))
		} else if (tokens[i][0] >= 'A' && tokens[i][0] <= 'Z') || (tokens[i][0] >= 'a' && tokens[i][0] <= 'z') {
			value := datanum[tokens[i]]
			stack = append(stack, value)
		} else {
			if tokens[i] == "+" {
				stack = append(stack[:len(stack)-2], stack[len(stack)-1]+stack[len(stack)-2])

			} else if tokens[i] == "-" {
				stack = append(stack[:len(stack)-2], stack[len(stack)-2]-stack[len(stack)-1])

			} else if tokens[i] == "*" {
				stack = append(stack[:len(stack)-2], stack[len(stack)-1]*stack[len(stack)-2])

			} else if tokens[i] == "/" {
				if stack[len(stack)-2] == 0 {
					stack = append(stack[:len(stack)-2], 0)
				} else {
					stack = append(stack[:len(stack)-2], stack[len(stack)-2]/stack[len(stack)-1])
				}
			}
		}
	}
	return stack[0]
}

func main() {
	var lexems chan Lexem

	a := input.Gets()
	input.Scanf("\n")
	lexems = make(chan Lexem, len(a))
	lexer(a, lexems)
	var alllexems []Lexem
	for 0 < len(lexems) {
		alllexems = append(alllexems, <-lexems)
	}
	//fmt.Println(alllexems)
	lindex := 0
	var tokens []string
	lindex, tokens = startparse(lindex, alllexems, tokens)
	var x int
	n := 0
	datanum := make(map[string]int)
	for i := 0; i < len(alllexems); i++ {
		if (alllexems[i].Tag&VAR != 0) && datanum[alllexems[i].Image] == 0 {
			datanum[alllexems[i].Image] = n + 1
			n++
		}
	}
	//fmt.Println(datanum)
	var newdatanumber []int
	for i := 0; i < n; i++ {
		input.Scanf("%d", &x)
		newdatanumber = append(newdatanumber, x)

	}
	//fmt.Println(newdatanumber)
	for k, v := range datanum {
		datanum[k] = newdatanumber[v-1]
	}
	//fmt.Println(datanum)
	result := getanswer(tokens, datanum)

	fmt.Println(result)

}
