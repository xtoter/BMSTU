package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

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

type token struct {
	data   string
	coordx int
	coordy int
	tp     string
}

func isTerminal(s string) bool {
	if !(s == "PROG" || s == "RA" || s == "PROG_end" || s == "EXPR") {
		return true
	}
	return false
}

var delta map[string]map[string][]string = map[string]map[string][]string{
	"PROG": {
		"COMMENT":      {"COMMENT", "PROG"},
		"LEFT_bracket": {"LEFT_bracket", "RA"},
		"ENDOF":        {"ENDOF"},
	},
	"RA": {
		"axiom": {"axiom", "NONT", "RIGHT_bracket", "equal", "EXPR", "POINT", "PROG_end"},
		"NONT":  {"NONT", "RIGHT_bracket", "equal", "EXPR", "POINT", "PROG"},
		"ENDOF": {"ENDOF"},
	},
	"PROG_end": {
		"LEFT_bracket": {"LEFT_bracket", "NONT", "RIGHT_bracket", "equal", "EXPR", "POINT", "PROG_end"},
		"COMMENT":      {"COMMENT", "PROG_end"},
		"ENDOF":        {"ENDOF"},
	},
	"EXPR": {
		"LEFT_bracket": {"LEFT_bracket", "NONT", "RIGHT_bracket", "EXPR"},
		"TERM":         {"TERM", "EXPR"},
		"ALT":          {"ALT", "EXPR"},
		"ENDOF":        {"ENDOF"},
		"POINT":        {""},
	},
}

var tokens []token
var comment, LEFT_bracket, RIGHT_bracket, equal, alt, term, nont, axiom, point *regexp.Regexp

func nextToken() token {
	if len(tokens) != 0 {
		t := tokens[0]
		loc_s := comment.FindIndex([]byte(t.data))
		loc_lb := LEFT_bracket.FindIndex([]byte(t.data))
		loc_rb := RIGHT_bracket.FindIndex([]byte(t.data))
		loc_eq := equal.FindIndex([]byte(t.data))
		loc_alt := alt.FindIndex([]byte(t.data))
		loc_t := term.FindIndex([]byte(t.data))
		loc_nt := nont.FindIndex([]byte(t.data))
		loc_ax := axiom.FindIndex([]byte(t.data))
		loc_p := point.FindIndex([]byte(t.data))
		// loc_i := ident.FindIndex([]byte(t.data))
		out := token{}
		//fmt.Println(loc_s, t.data)
		switch {
		case len(loc_s) != 0 && loc_s[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_s[1]:]
			out.data = t.data[loc_s[0]:loc_s[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_s[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "COMMENT"

			return out
		case len(loc_rb) != 0 && loc_rb[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_rb[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_rb[0]:loc_rb[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_rb[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "RIGHT_bracket"
			return out
		case len(loc_lb) != 0 && loc_lb[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_lb[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_lb[0]:loc_lb[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_lb[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "LEFT_bracket"
			return out
		case len(loc_eq) != 0 && loc_eq[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_eq[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_eq[0]:loc_eq[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_eq[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "equal"
			return out

		case len(loc_alt) != 0 && loc_alt[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_alt[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_alt[0]:loc_alt[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_alt[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "ALT"
			return out
		case len(loc_p) != 0 && loc_p[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_p[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_p[0]:loc_p[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_p[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "POINT"
			return out
		case len(loc_ax) != 0 && loc_ax[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_ax[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_ax[0]:loc_ax[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_ax[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "axiom"
			return out
		case len(loc_nt) != 0 && loc_nt[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_nt[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_nt[0]:loc_nt[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_nt[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "NONT"
			return out
		case len(loc_t) != 0 && loc_t[0] == 0:
			//	fmt.Println("string (", t.coordx, t.coordy, ")", t.data[loc_s[0]:loc_s[1]])
			s := t.data[loc_t[1]:]
			//fmt.Println("rule (", t.coordx, t.coordy, ")", t.data, loc_r)
			out.data = t.data[loc_t[0]:loc_t[1]]
			out.coordx = t.coordx
			out.coordy = t.coordy
			if len(s) != 0 {
				tokens[0].data = s
				tokens[0].coordx += loc_t[1]
			} else {
				tokens = tokens[1:]
			}
			out.tp = "TERM"
			return out

		default:
			//fmt.Println("error", t.coordx, t.coordy)
			tok_err := token{}
			if len(t.data[1:]) == 0 {
				tokens = tokens[1:]
				break
			}
			tokens[0].data = tokens[0].data[1:]
			tokens[0].coordx += 1
			nextToken()
			tok_err.tp = "error"
			tok_err.coordx = t.coordx
			tok_err.coordy = t.coordy
			tok_err.data = ""
			return tok_err
		}
		//fmt.Println(loc)
	}
	return token{tp: "ENDOF"}
}

type Node struct {
}

type Inner struct {
	tp       string
	nterm    token
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

func (in Inner) print(ident string, rules *Inner) {
	if in.tp == "LEAF" {
		fmt.Println(ident + "Лист: " + in.nterm.data)
		// child := &Inner{in.tp, in.nterm, 0, nil}
		// rules.children = append(rules.children, child)
		if in.nterm.tp == "POINT" {
			rut.children = append(rut.children, start_rule)
		}
		return
	}

	if in.nterm.tp == "RA" || in.nterm.tp == "PROG_end" {
		start_rule = &Inner{in.tp, in.nterm, 0, nil}
		rules = start_rule
	}

	fmt.Println(ident + "Внутренний узел: " + in.nterm.data + ", Правило: " + in.tp)
	//fmt.Println(in.children)
	for i := 0; i < len(in.children); i++ {
		child := &Inner{in.children[i].tp, in.children[i].nterm, 0, nil}
		//rules.children = append(rules.children, child)

		in.children[i].print(ident+"	", child)
	}
}

func (in Inner) init(ident string, rules *Inner) {
	if in.tp == "LEAF" {
		fmt.Println(ident + "Лист: " + in.nterm.data)
		// child := &Inner{in.tp, in.nterm, 0, nil}
		// rules.children = append(rules.children, child)
		if in.nterm.tp == "POINT" {
			rut.children = append(rut.children, start_rule)
		}
		return
	}

	if in.nterm.tp == "RA" || in.nterm.tp == "PROG_end" {
		start_rule = &Inner{in.tp, in.nterm, 0, nil}
		rules = start_rule
	}

	fmt.Println(ident + "Внутренний узел: " + in.nterm.data + ", Правило: " + in.tp)
	//fmt.Println(in.children)
	for i := 0; i < len(in.children); i++ {
		child := &Inner{in.children[i].tp, in.children[i].nterm, 0, nil}
		rules.children = append(rules.children, child)

		in.children[i].print(ident+"	", child)
	}
}

func topDownParse() Inner {
	s := make(stack, 0)
	sparent := &Inner{}
	s = s.Push(Pair{sparent, "$"})
	s = s.Push(Pair{sparent, "PROG"})
	s_Prog := &Inner{} //sparent //.children
	a := nextToken()
	var X string
	var parent *Inner
	parent, X = s.Top()
	flag_prog_start := true
	var prog_start_pointer *Inner
	for X != "$" {
		//fmt.Println(s, a.data, a.tp, X)
		if isTerminal(X) {
			if X == a.tp {
				s, _ = s.Pop()

				parent.children = append(parent.children, &Inner{"LEAF", a, 0, nil})
				println(X)
				if X == "POINT" {
					flag_prog_start = true
					s_Prog.children = append(s_Prog.children, prog_start_pointer)
				}
				a = nextToken()

			} else if X == "" {
				s, _ = s.Pop()
			}
		} else if delta[X][a.tp] != nil {
			s, _ = s.Pop()
			inner := Inner{X, a, 0, []*Inner{}}
			if (X == "PROG_end" || X == "RA") && flag_prog_start {
				flag_prog_start = false
				prog_start_pointer = &inner
			}
			parent.children = append(parent.children, &inner)
			for i := len(delta[X][a.tp]) - 1; i >= 0; i-- {
				s = s.Push(Pair{&inner, delta[X][a.tp][i]})
			}
		} else {
			fmt.Println("error", a.coordx, a.coordy)
		}
		parent, X = s.Top()
	}
	return *sparent.children[0]
}

func main2() {

	comment = regexp.MustCompile(`;([a-z]| )*`)
	LEFT_bracket = regexp.MustCompile(`( )*\(`)
	RIGHT_bracket = regexp.MustCompile(`( )*\)`)
	equal = regexp.MustCompile(`( )*\=`)
	alt = regexp.MustCompile(`( )*\|`)
	term = regexp.MustCompile(`( )*([^\(\)\|\\])|\\\(|\\\)|\\\|| |\+|\*`)
	nont = regexp.MustCompile(`( )*[A-Z][0-9]*`)
	axiom = regexp.MustCompile(`( )*axiom`)
	point = regexp.MustCompile(`( )*\.`)
	//loc_r := str.FindIndex([]byte("; akbdsba ashdsaksdh"))
	//fmt.Println(loc_r, str)
	//return
	f, err := os.Open("test")
	if err != nil {
		panic(err)
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	str_data := string(c)
	f.Close()
	text := strings.Split(str_data, "\n")
	//fmt.Println(text[0])
	for i := 0; i < len(text); i++ {

		a := token{}
		a.data = text[i]
		a.coordx = 0
		a.coordy = i
		tokens = append(tokens, a)
	}

	sp := topDownParse()
	rut = Inner{}
	sp.print("", nil)
}
