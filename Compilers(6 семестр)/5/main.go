package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	text  string
	line  int
	pos   int
	index int
}

func NewPosition(text string) *Position {
	return &Position{text: text, line: 1, pos: 1, index: 0}
}
func (p *Position) GetText() string {
	return p.text
}
func (p *Position) GetLine() int {
	return p.line
}

func (p *Position) SetLine(line int) {
	p.line = line
}

func (p *Position) GetPos() int {
	return p.pos
}

func (p *Position) SetPos(pos int) {
	p.pos = pos
}

func (p *Position) GetIndex() int {
	return p.index
}

func (p *Position) SetIndex(index int) {
	p.index = index
}

func (p *Position) GetCp() int {
	if p.index == len(p.text) {
		return -1
	}
	return int(p.text[p.index])
}

func (p *Position) IsWhiteSpace() bool {
	return p.index != len(p.text) &&
		(int(p.GetCp()) == ' ')
}
func (p *Position) IsEOF() bool {
	return p.index == len(p.text)
}
func (p *Position) IsNumber() bool {
	temp := p.GetCp()

	return p.index != len(p.text) && (temp >= '0' && temp <= '9')
}
func (p *Position) IsLetter() bool {
	temp := p.GetCp()

	return p.index != len(p.text) && ((temp >= 'a' && temp <= 'z') || (temp >= 'A' && temp <= 'Z'))
}

func (p *Position) IsLetterOrDigit() bool {
	return p.IsNumber() || p.IsLetter()
}

func (p *Position) IsNewLine() bool {
	if p.index == len(p.text) {
		return true
	}
	if '\r' == p.text[p.index] && p.index+1 < len(p.text) {
		return '\n' == p.text[p.index+1]
	}
	return '\n' == p.text[p.index]
}

func (p *Position) Next() *Position {
	if p.index < len(p.text) {
		if p.IsNewLine() {
			if '\r' == p.text[p.index] {
				p.index++
			}
			p.line++
			p.pos = 1
		} else {
			p.pos++
		}
		p.index++
	}
	return p
}

func (p *Position) Copy() *Position {
	return &Position{text: p.text, line: p.line, pos: p.pos, index: p.index}
}

func (p *Position) CompareTo(other *Position) int {
	return p.index - other.index
}

func (p *Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.line, p.pos)
}

type Fragment struct {
	starting, following Position
}

func NewFragment(starting, following Position) *Fragment {
	return &Fragment{starting: starting, following: following}
}

func (f *Fragment) String() string {
	return fmt.Sprintf("%s-%s", f.starting.String(), f.following.String())
}

type Automata struct {
	messages map[*Position]string
	program  string
	pos      *Position
	state    int
}

func NewAutomata(program string) *Automata {
	return &Automata{
		messages: make(map[*Position]string),
		program:  program,
		pos:      NewPosition(program),
		state:    0,
	}
}

func (a *Automata) get_code(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return 0 //DIGIT
	case c == '(':
		return 1 // (
	case c == ')':
		return 2
	case c == ' ' || c == '\n' || c == '\r' || c == '\t':
		return 3 //newline
	}

	switch c {
	case 'u':
		return 4 //u
	case 'n':
		return 5 //n
	case 's':
		return 6 //s
	case 'e':
		return 7 //e
	case 't':
		return 8 //t

	}

	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return 9
	}

	return -1
}

func (a *Automata) get_state_name(state int) string {
	switch state {
	case 1:
		return "Number"
	case 2:
		return "WhiteSpace"
	case 3:
		return "Ident"
	case 4:
		return "Keyword"
	case 6:
		return "STR_LITERAL"
	default:
		return "ERROR"
	}
}

var err bool = false

func (a *Automata) run() {
	fmt.Println("\nTokens:")
	for a.pos.GetCp() != -1 {
		word := ""
		state := 0
		final_state := false
		start := a.pos.Copy()
		cond := true
		for cond {
			if a.pos.GetCp() == -1 {
				cond = false
			}
			//fmt.Println("cur char jump code ", curr_char, jump_code)
			var curr_char byte
			if a.pos.GetIndex() == len(a.program) {
				final_state = true
				//fmt.Print("(-1)\n")
				cond = false
				break
			}
			curr_char = a.program[a.pos.GetIndex()]

			jump_code := a.get_code(curr_char)
			if jump_code == -1 {
				cond = false
				if !err {
					a.messages[a.pos.Copy()] = "Unexpected characters"
					err = true
				}

				break
			}
			err = false

			//fmt.Printf("(%d)->", state)
			//fmt.Printf("[%c]->", curr_char)

			next_state := table[state][jump_code]
			//fmt.Println("nextstate,jumpcode", next_state, jump_code)
			if next_state == -1 {
				final_state = true
				//fmt.Print("(-1)\n")
				cond = false
				break
			}

			word += string(curr_char)
			state = next_state
			a.pos.Next()
		}
		//fmt.Println("exit")

		if final_state {
			//	fmt.Println("final")
			frag := Fragment{*start, *a.pos}
			fmt.Printf("%s %s: %s\n", a.get_state_name(state), frag.String(), strings.ReplaceAll(word, "\n", " "))
			continue
		}

		a.pos.Next()
	}
}

func (a *Automata) output_messages() {
	fmt.Println("\nMessages:")
	for key, value := range a.messages {
		fmt.Print("ERROR ")
		fmt.Printf("(%d, %d): %s\n", key.GetLine(), key.GetPos(), value)
	}
}

var table = [][]int{
	/*  digit   (   )   space   u   n   s  e  t  .  oth*/
	/*  START   */ {1, 5, -1, 2, 7, -1, 9, -1, -1, 3, -1},
	/*  Number1 */ {1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	/*  Space2  */ {-1, -1, -1, 2, -1, -1, -1, -1, -1, -1, -1},
	/*  Ident3  */ {3, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1},
	/*  keyword4*/ {-1, -1, -1, -1, -1, -1, -1, -1, -1, 3, -1},
	/*  word5   */ {-1, -1, 6, -1, -1, -1, -1, -1, -1, 5, -1},
	/*  strlit6 */ {-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
	/*  u_7     */ {-1, -1, -1, -1, -1, 8, -1, -1, -1, 3, -1},
	/*  n_8     */ {-1, -1, -1, -1, -1, -1, 9, -1, -1, 3, -1},
	/*  s_9     */ {-1, -1, -1, -1, -1, -1, -1, 10, -1, 3, -1},
	/*  e_10    */ {-1, -1, -1, -1, -1, -1, -1, -1, 4, -1, -1},
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	text := ""
	fmt.Println("   1234567890123456789")
	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Printf("%d [%s]\n", i, line)
		text += line + "\n"
		i++
	}
	auto := NewAutomata(text)
	auto.run()
	//auto.output_messages()
}
