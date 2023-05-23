package parser

import (
	"2.3/lexer"
	"fmt"
	"strconv"
)

type Leaf struct {
	Token lexer.Token
}
type Node interface {
	Print(indent string)
}

func (leaf Leaf) Print(indent string) {
	fmt.Printf("%sЛист: %s\n", indent, leaf.Token.Value)
}

type Inner struct {
	NTerm    lexer.TokenType
	RuleID   int
	Children []Node
}

func (inner Inner) Print(indent string) {
	fmt.Printf("%sВнутренний узел: %v, правило: %v:\n", indent, inner.NTerm, inner.RuleID)
	for _, child := range inner.Children {
		child.Print(indent + "\t")
	}
}

type Table map[lexer.TokenType]map[lexer.TokenType]Rule
type Rule struct {
	ID         int
	TokenTypes []lexer.TokenType
}

func error(pos string, message string) {
	panic("Ошибка: " + pos + " " + message)
}

type Parser struct {
	Tokens   []lexer.Token
	CurIndex int
}

func NewLexer(input []lexer.Token) *Parser {
	return &Parser{
		Tokens:   input,
		CurIndex: 0,
	}
}
func (p *Parser) TopDownParse(delta Table) Node {
	sparent := Inner{} // Фиктивный родитель для аксиомы
	stack := make([]struct {
		Parent Inner
		X      lexer.TokenType
	}, 0)
	stack = append(stack, struct {
		Parent Inner
		X      lexer.TokenType
	}{sparent, lexer.Dollar})
	stack = append(stack, struct {
		Parent Inner
		X      lexer.TokenType
	}{sparent, SymbolS})
	a := lexer.NextToken()
	for {
		top := stack[len(stack)-1]
		parent := top.Parent
		X := top.X
		if X.isTerminal() {
			if X != a.Type {
				stack = stack[:len(stack)-1]
				parent.Children = append(parent.Children, Leaf{a})
				a = lexer.NextToken()
			} else {
				error(a.Pos, "Ожидался "+strconv.Itoa(int(X))+", получен "+a.Lexeme)
			}
		} else if rule, ok := delta[X][a.Type]; ok {
			stack = stack[:len(stack)-1]
			inner := Inner{X, rule.ID, []Node{}}
			parent.Children = append(parent.Children, inner)
			for i := len(rule.Symbols) - 1; i >= 0; i-- {
				stack = append(stack, struct {
					Parent Inner
					X      lexer.TokenType
				}{inner, rule.Symbols[i]})
			}
		} else {
			error(a.Pos, "Ожидался "+strconv.Itoa(int(X))+", получен "+a.Lexeme)
		}
		if X == SymbolDollar {
			break
		}
	}
	return sparent.Children[0]
}
