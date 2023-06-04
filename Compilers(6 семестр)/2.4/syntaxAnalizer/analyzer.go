package syntaxAnalizer

import (
	"fmt"

	"2.4/lexer"
)

type Parser struct {
	tokens         []lexer.Token
	currentToken   lexer.Token
	nextTokenIndex int
}

var counter int

type Tree struct {
	Str          string
	Num          int
	Leaves       []Tree
	CurrentToken lexer.Token
}

func (tree Tree) ToString() string {

	if tree.Str == "" {
		return tree.CurrentToken.ToString()
	}
	return tree.Str
}
func generateDOT(tree Tree) string {

	dotCode := "digraph Tree {\n"
	dotCode += "  node [shape=box];\n"
	dotCode += generateNodeDOT(tree)
	dotCode += "}\n"
	return dotCode
}
func generateNodeDOT(tree Tree) string {

	nodeID := fmt.Sprintf("%d", tree.Num)
	dotCode := fmt.Sprintf("  %s [label=\"%s\"];\n", nodeID, tree.ToString())
	for i := 0; i < len(tree.Leaves); i++ {

		leafID := fmt.Sprintf("%d", tree.Leaves[i].Num)
		dotCode += fmt.Sprintf("  %s -> %s;\n", nodeID, leafID)
		dotCode += generateNodeDOT(tree.Leaves[i])
	}
	return dotCode
}

func NewParser(input string) *Parser {
	tokens := lexer.Lex(input)
	return &Parser{
		tokens:         tokens,
		currentToken:   lexer.Token{},
		nextTokenIndex: 0,
	}
}

func (p *Parser) Parse() Tree {
	p.nextToken()
	counter = 0

	cur := p.parseGrammar()
	cur.Num = counter
	counter++
	fmt.Println(generateDOT(cur))
	if p.currentToken.Type != lexer.TokenEOF {
		fmt.Println("Parsing error: unexpected token", p.currentToken)
	} else {
		fmt.Println("Parsing completed successfully.")
	}
	return cur
}

func (p *Parser) parseGrammar() Tree {
	var cur Tree
	cur.Str = "Grammar"
	cur.Num = counter
	counter++
	cur.Leaves = append(cur.Leaves, p.parseRule())
	for p.nextTokenIndex != len(p.tokens) {
		cur.Leaves = append(cur.Leaves, p.parseRule())
	}

	return cur
}

func (p *Parser) parseRule() Tree {
	var cur Tree
	cur.Str = "Rule"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenNonterminal {
		cur.Leaves = append(cur.Leaves, p.parseNonterminal())

		if p.currentToken.Type == lexer.TokenLeftParenthesis {
			cur.Leaves = append(cur.Leaves, p.nextToken())
			cur.Leaves = append(cur.Leaves, p.parseExpression())

			if p.currentToken.Type == lexer.TokenRightParenthesis {
				cur.Leaves = append(cur.Leaves, p.nextToken())
			} else {
				fmt.Println("Parsing error: expected ')', but got", p.currentToken)
			}
		} else {
			fmt.Println("Parsing error: expected '(', but got", p.currentToken)
		}
	} else {
		fmt.Println("Parsing error: expected '(', but got", p.currentToken)
	}

	return cur
}

func (p *Parser) parseExpression() Tree {
	var cur Tree
	cur.Str = "Expression"
	cur.Num = counter
	counter++
	cur.Leaves = append(cur.Leaves, p.parseTerm())
	for p.currentToken.Type == lexer.TokenComma {
		cur.Leaves = append(cur.Leaves, p.nextToken())
		cur.Leaves = append(cur.Leaves, p.parseTerm())
	}

	return cur

}

func (p *Parser) parseTerm() Tree {
	var cur Tree
	cur.Str = "Term"
	cur.Num = counter
	counter++
	cur.Leaves = append(cur.Leaves, p.parseFactor())
	for p.currentToken.Type != lexer.TokenEOF && (p.currentToken.Type == lexer.TokenTerminal || p.currentToken.Type == lexer.TokenNonterminal || p.currentToken.Type == lexer.TokenLeftParenthesis || p.currentToken.Type == lexer.TokenCurlyBracketOpen) {
		cur.Leaves = append(cur.Leaves, p.parseFactor())
	}

	return cur
}

func (p *Parser) parseFactor() Tree {
	var cur Tree
	cur.Str = "Factor"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenNonterminal {
		cur.Leaves = append(cur.Leaves, p.parseNonterminal())
	} else if p.currentToken.Type == lexer.TokenTerminal {
		cur.Leaves = append(cur.Leaves, p.parseTerminal())
	} else if p.currentToken.Type == lexer.TokenLeftParenthesis {
		cur.Leaves = append(cur.Leaves, p.parseGrouping())
	} else if p.currentToken.Type == lexer.TokenCurlyBracketOpen {
		cur.Leaves = append(cur.Leaves, p.parseOption())

	} else {
		fmt.Println("Parsing error: unexpected token", p.currentToken.Value)
	}

	return cur
}
func (p *Parser) parseGrouping() Tree {
	var cur Tree
	cur.Str = "Grouping"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenLeftParenthesis {
		cur.Leaves = append(cur.Leaves, p.nextToken())
		cur.Leaves = append(cur.Leaves, p.parseExpression())
		cur.Leaves = append(cur.Leaves, p.nextToken())
	} else {
		fmt.Println("Parsing error: expected '(', but got", p.currentToken.Value)
	}

	return cur
}
func (p *Parser) parseOption() Tree {
	var cur Tree
	cur.Str = "Option"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenCurlyBracketOpen {
		cur.Leaves = append(cur.Leaves, p.nextToken())
		cur.Leaves = append(cur.Leaves, p.parseExpression())
		cur.Leaves = append(cur.Leaves, p.nextToken())
	} else {
		fmt.Println("Parsing error: expected '{', but got", p.currentToken.Value)
	}

	return cur
}

func (p *Parser) parseNonterminal() Tree {
	var cur Tree
	cur.Str = "NonTerminal"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenNonterminal {
		cur.Leaves = append(cur.Leaves, p.nextToken())
	} else {
		fmt.Println("Parsing error: expected Nonterminal, but got", p.currentToken.Value)
	}

	return cur
}

func (p *Parser) parseTerminal() Tree {
	var cur Tree
	cur.Str = "Terminal"
	cur.Num = counter
	counter++
	if p.currentToken.Type == lexer.TokenTerminal {
		cur.Leaves = append(cur.Leaves, p.nextToken())
	} else {
		fmt.Println("Parsing error: expected Terminal, but got", p.currentToken.Value)
	}

	return cur
}

func (p *Parser) nextToken() Tree {
	var cur Tree
	cur.Num = counter
	counter++
	cur.CurrentToken = p.currentToken
	if p.nextTokenIndex < len(p.tokens) {
		p.currentToken = p.tokens[p.nextTokenIndex]
		p.nextTokenIndex++
	} else {
		p.currentToken = lexer.Token{
			Type:  lexer.TokenEOF,
			Value: "",
		}
	}

	return cur
}
