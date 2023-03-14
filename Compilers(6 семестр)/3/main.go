package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var ptr int = 0

type lexem struct {
	data        string
	row, column int
}

func GetData() []lexem {
	var out []lexem
	for i := 1; i < len(os.Args); i++ {
		row, column := 0, 0
		data, err := ioutil.ReadFile(os.Args[i])
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
		if size := sb.Len(); size > 0 {
			lexem := lexem{string(sb.String()), row, column - size}
			out = append(out, lexem)
		}
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
func lexer(in []lexem) {
	pattern1 := regexp.MustCompile(`mov|eax`)
	pattern2 := regexp.MustCompile(`[0-9a-fA-F]+h`)
	pattern3 := regexp.MustCompile(`[0-9]+`)
	pattern4 := regexp.MustCompile(`[a-zA-Z\p{L}]+[a-zA-Z\p{L}0-9]*`)
	lexem, cond := lexem{}, true
	for cond {
		lexem, cond = nextToken(in)
		//fmt.Println(lexem)
		for len(lexem.data) > 0 {

			loc1 := pattern1.FindIndex([]byte(lexem.data))
			loc2 := pattern2.FindIndex([]byte(lexem.data))
			loc3 := pattern3.FindIndex([]byte(lexem.data))
			loc4 := pattern4.FindIndex([]byte(lexem.data))
			//	fmt.Println(lexem)
			//	fmt.Println(loc1, loc2, loc3, loc4)
			if len(loc1) > 0 && loc1[0] == 0 {
				if loc1[1] == len(lexem.data) {
					out("KeyWord", lexem.row, lexem.column, lexem.data)
					lexem.data = ""
					continue
				}
				out("KeyWord", lexem.row, lexem.column, lexem.data[:loc1[1]])
				lexem.data = lexem.data[loc1[1]:]
				lexem.column += loc1[1]
				continue
			}
			if len(loc2) > 0 && loc2[0] == 0 {
				if loc2[1] == len(lexem.data) {
					out("Hex", lexem.row, lexem.column, lexem.data)
					lexem.data = ""
					continue
				}
				out("Hex", lexem.row, lexem.column, lexem.data[:loc2[1]])
				lexem.data = lexem.data[loc2[1]:]
				lexem.column += loc2[1]
				continue
			}
			if len(loc3) > 0 && loc3[0] == 0 {
				if loc3[1] == len(lexem.data) {
					out("Dec", lexem.row, lexem.column, lexem.data)
					lexem.data = ""
					continue
				}
				out("Dec", lexem.row, lexem.column, lexem.data[:loc3[1]])
				lexem.data = lexem.data[loc3[1]:]
				lexem.column += loc3[1]
				continue
			}
			if len(loc4) > 0 && loc4[0] == 0 {
				if loc4[1] == len(lexem.data) {
					out("Word", lexem.row, lexem.column, lexem.data)
					lexem.data = ""
					continue
				}
				out("Word", lexem.row, lexem.column, lexem.data[:loc4[1]])
				lexem.data = lexem.data[loc4[1]:]
				lexem.column += loc4[1]
				continue
			}
			fmt.Println("parse error")
			os.Exit(0)
		}
	}
}
func main() {
	lexer(GetData())
}
