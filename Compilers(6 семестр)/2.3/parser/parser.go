package parser

import (
	"2.3/lexer"
	"fmt"
)

var delta map[string]map[string][]string = map[string]map[string][]string{
	"Grammar": {
		"COMMENT": {"COMMENT", "Grammar"},
		"NONT":    {"Rule", "Grammar"},
		"axiom":   {"axiom", "Rule", "Grammar"},
		"ENDOF":   {"$"},
	},
	"Rule": {
		"NONT":  {"NONT", "Exprbracket"},
		"axiom": {"axiom", "NONT", "Exprbracket"},
		"ENDOF": {"$"},
	},
	"Exprbracket": {
		"LEFT_bracket": {"LEFT_bracket", "Expr", "RIGHT_bracket", "Exprbracket"},
		"ENDOF":        {"$"},
		"NONT":         {""},
		"axiom":        {""},
	},
	"Expr": {
		"TERM":          {"TERM", "Expr"},
		"NONT":          {"NONT", "Expr"},
		"ENDOF":         {"ENDOF"},
		"RIGHT_bracket": {""},
	},
}

type Node struct {
}

type stack []Pair

func (s stack) Top() (*Inner, string) {
	l := len(s)
	if l == 0 {
		return nil, ""
		//panic("empty stack")
	}
	return s[l-1].node, s[l-1].symb
}
func (s stack) Push(v Pair) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, Pair) {

	l := len(s)
	return s[:l-1], s[l-1]
}

type Inner struct {
	tp       string
	nterm    lexer.Token
	reluId   int
	children []*Inner
}
type Pair struct {
	node *Inner
	symb string
}

var rut Inner

// var rules *Inner
var start_rule *Inner
var allToken []lexer.Token

func nextToken() lexer.Token {
	if len(allToken) > 0 {
		out := allToken[0]
		allToken = allToken[1:]
		return out
	}
	return lexer.Token{}
}
func isTerminal(s string) bool {
	if !(s == "Grammar" || s == "Rule" || s == "Exprbracket" || s == "Expr") {
		return true
	}
	return false
}

var tokens []string = []string{"TERM", "NONT", "LEFT_bracket", "RIGHT_bracket", "COMMENT", "axiom", "Dollar", "ENDOF", "TokenError"}

func (in Inner) Print(ident string, rules *Inner) {
	if in.tp == "LEAF" {
		fmt.Println(ident + "Лист: " + in.nterm.Value)
		// child := &Inner{in.tp, in.nterm, 0, nil}
		// rules.children = append(rules.children, child)
		if tokens[in.nterm.Type] == "POINT" {
			rut.children = append(rut.children, start_rule)
		}
		return
	}

	if tokens[in.nterm.Type] == "RA" || tokens[in.nterm.Type] == "PROG_end" {
		start_rule = &Inner{in.tp, in.nterm, 0, nil}
		rules = start_rule
	}

	fmt.Println(ident + "Внутренний узел: " + in.nterm.Value + ", Правило: " + in.tp)
	//fmt.Println(in.children)
	for i := 0; i < len(in.children); i++ {
		child := &Inner{in.children[i].tp, in.children[i].nterm, 0, nil}
		//rules.children = append(rules.children, child)

		in.children[i].Print(ident+"	", child)
	}
}
func TopDownParse(in []lexer.Token) Inner {
	allToken = in
	s := make(stack, 0)
	sparent := &Inner{}
	s = s.Push(Pair{sparent, "$"})
	s = s.Push(Pair{sparent, "Grammar"})
	a := nextToken()
	var X string
	var parent *Inner
	parent, X = s.Top()
	for X != "$" {
		//fmt.Println("test", "x", X, "token", tokens[a.Type], delta[X][tokens[a.Type]], "a", a)
		if isTerminal(X) {
			if X == tokens[a.Type] {
				s, _ = s.Pop()

				parent.children = append(parent.children, &Inner{"LEAF", a, 0, nil})
				println(X)
				a = nextToken()

			} else if X == "" {
				s, _ = s.Pop()
			}
		} else if delta[X][tokens[a.Type]] != nil {
			s, _ = s.Pop()
			inner := Inner{X, a, 0, []*Inner{}}

			parent.children = append(parent.children, &inner)
			for i := len(delta[X][tokens[a.Type]]) - 1; i >= 0; i-- {
				s = s.Push(Pair{&inner, delta[X][tokens[a.Type]][i]})
			}
		} else {
			fmt.Println("error", a.Row, a.Column, X)
		}
		parent, X = s.Top()
	}
	return *sparent.children[0]
}
