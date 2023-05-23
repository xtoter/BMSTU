package lexer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TokenType int

const (
	TokenNonterminal TokenType = iota
	TokenTerminal
	TokenLeftParenthesis
	TokenRightParenthesis
	TokenComma
	TokenCurlyBracketOpen
	TokenCurlyBracketClose
	TokenEOF
	TokenError
)

func (t TokenType) String() string {
	a := strconv.Itoa(int(t))
	return a
}

type Token struct {
	Type  TokenType
	Value string
}

func (t *TokenType) ToString() string {
	return strconv.Itoa(int(*t))
}
func replaceQuotes(str string) string {
	return strings.ReplaceAll(str, "\"", "'")
}
func (t *Token) ToString() string {
	return "val : " + replaceQuotes(t.Value) + " Type : " + t.Type.ToString()

	return ""
}

type Lexer struct {
	input    string
	position int
}

var (
	nonterminalPattern     = regexp.MustCompile(`^[A-Z]+\b`)
	terminalPattern        = regexp.MustCompile(`^"[^"]*"`)
	leftParenthesisRegex   = regexp.MustCompile(`^\(`)
	rightParenthesisRegex  = regexp.MustCompile(`^\)`)
	commaRegex             = regexp.MustCompile(`^,`)
	curlyBracketOpenRegex  = regexp.MustCompile(`^{`)
	curlyBracketCloseRegex = regexp.MustCompile(`^}`)
	whitespaceRegex        = regexp.MustCompile(`^\s+`)
)

func NewLexer(input string) *Lexer {
	return &Lexer{
		input:    input,
		position: 0,
	}
}

func (l *Lexer) GetNextToken() Token {
	if l.position >= len(l.input) {
		return Token{Type: TokenEOF}
	}

	remainingInput := l.input[l.position:]

	if matches := nonterminalPattern.FindString(remainingInput); matches != "" {
		l.position += len(matches)
		return Token{Type: TokenNonterminal, Value: matches}
	}

	if matches := terminalPattern.FindString(remainingInput); matches != "" {
		l.position += len(matches)
		return Token{Type: TokenTerminal, Value: matches}
	}

	if leftParenthesisRegex.MatchString(remainingInput) {
		l.position++
		return Token{Type: TokenLeftParenthesis, Value: "("}
	}

	if rightParenthesisRegex.MatchString(remainingInput) {
		l.position++
		return Token{Type: TokenRightParenthesis, Value: ")"}
	}

	if commaRegex.MatchString(remainingInput) {
		l.position++
		return Token{Type: TokenComma, Value: ","}
	}

	if curlyBracketOpenRegex.MatchString(remainingInput) {
		l.position++
		return Token{Type: TokenCurlyBracketOpen, Value: "{"}
	}

	if curlyBracketCloseRegex.MatchString(remainingInput) {
		l.position++
		return Token{Type: TokenCurlyBracketClose, Value: "}"}
	}

	if whitespaceRegex.MatchString(remainingInput) {
		// Пропускаем пробелы
		l.position += len(whitespaceRegex.FindString(remainingInput))
		return l.GetNextToken()
	}
	l.position++
	// Invalid character
	return Token{Type: TokenError, Value: string(remainingInput[0])}
}

func Lex(input string) []Token {

	lexer := NewLexer(input)
	var tokens []Token
	for {
		token := lexer.GetNextToken()

		if token.Type == TokenEOF {
			break
		} else if token.Type == TokenError {
			fmt.Printf("Lexical error: Invalid character '%s'\n", token.Value)

		}
		tokens = append(tokens, token)
		token.Type.String()
		fmt.Printf("Token: %s, Value: %s\n", token.Type, token.Value)
	}
	return tokens
}
