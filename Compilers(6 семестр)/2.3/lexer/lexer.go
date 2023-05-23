package lexer

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type TokenType int

const (
	Star TokenType = iota
	Plus
	Leftbracket
	Rightbracket
	Terminal
	NonTerminal
	Left
	Right
	Aksiom
	Dollar
	TokenEOF
	TokenError
)

func (t TokenType) String() string {
	a := strconv.Itoa(int(t))
	return a
}

type Token struct {
	Type        TokenType
	Value       string
	row, column int
}

var ptr int = 0
var x, y int

type lexem struct {
	data        string
	row, column int
}

func GetData() []lexem {
	var out []lexem
	row, column := 0, 0
	data, err := ioutil.ReadFile("/home/xtoter/github/BMSTU/Compilers(6 семестр)/2.3/data/task.txt")

	if err != nil {
		fmt.Println(err)
	}
	var sb strings.Builder
	for _, curRune := range string(data) {
		switch curRune {
		case '\n':

			if size := sb.Len(); size > 0 {
				lexem := lexem{string(sb.String()), row, column - size}
				out = append(out, lexem)
				sb.Reset()
			}
			row++
			column = -1
		case '\t':
			if size := sb.Len(); size > 0 {
				lexem := lexem{string(sb.String()), row, column - size}
				out = append(out, lexem)
			}
			sb.Reset()
		case ' ':
			if size := sb.Len(); size > 0 {
				lexem := lexem{string(sb.String()), row, column - size}
				out = append(out, lexem)
			}
			sb.Reset()
		default:
			sb.WriteRune(curRune)
		}

		column++
	}
	x, y = row, column
	if size := sb.Len(); size > 0 {
		lexem := lexem{string(sb.String()), row, column - size}
		out = append(out, lexem)
	}
	return out
}
func out(typee string, x, y int, value string) {
	fmt.Printf("%s (%d, %d): %s\n", typee, x, y, value)
}
func nextToken(in []lexem) (lexem, bool) {
	if len(in) == ptr {
		return lexem{}, false
	}
	res := in[ptr]
	ptr++
	return res, true
}

func lexer(in []lexem) []Token {
	var result []Token
	pattern := []*regexp.Regexp{
		regexp.MustCompile(`\"\*\"`),
		regexp.MustCompile(`\"\+\"`),
		regexp.MustCompile(`\"\(\"`),
		regexp.MustCompile(`\"\)\"`),
		regexp.MustCompile(`\"[a-z0-9]\"`),
		regexp.MustCompile(`[A-Z]\'?`),
		regexp.MustCompile(`\(`),
		regexp.MustCompile(`\)`),
		regexp.MustCompile(`\*`)}
	lexem, cond := lexem{}, true
	for cond {
		lexem, cond = nextToken(in)
		//fmt.Println(lexem)
		var loc [][]int
		for len(lexem.data) > 0 {
			loc = [][]int{}
			for i := 0; i < len(pattern); i++ {
				loc = append(loc, pattern[i].FindIndex([]byte(lexem.data)))
			}
			cond := true
			for i := 0; i < len(loc) && cond; i++ {
				if len(loc[i]) > 0 && loc[i][0] == 0 {
					//	fmt.Println("test", loc)
					cur := Token{TokenType(i), lexem.data[:loc[i][1]], lexem.row, lexem.column}
					result = append(result, cur)
					//out(i, lexem.row, lexem.column, lexem.data[:loc[i][1]])
					lexem.data = lexem.data[loc[i][1]:]
					lexem.column += loc[i][1]
					cond = false
				}
			}
			if cond {
				fmt.Println("parse error")
				os.Exit(0)

			}

		}
	}
	result = append(result, Token{TokenEOF, "", x, y})
	return result
}
func GetTokens() []Token {
	return lexer(GetData())
}
