package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var CP, SP, BP, P int
var stack []int
var buffer []string

func push_asm_prog(x int) {
	stack[P] = x
	P++

}
func call_instruction(x int) {
	switch x {
	case -1:
		y := pop()
		x := pop()
		push(x + y)

	case -2:
		y := pop()
		x := pop()
		push(x - y)

	case -3:
		y := pop()
		x := pop()
		push(x / y)

	case -4:
		y := pop()
		x := pop()
		push(x % y)

	case -5:
		y := pop()
		x := pop()
		push(x * y)

	case -6:
		x := pop()
		push(0 - x)

	case -7:
		y := pop()
		x := pop()
		push(x & y)

	case -8:
		y := pop()
		x := pop()
		push(x | y)

	case -9:
		x := pop()
		push(^x)

	case -10:
		x := pop()
		push(x)
		push(x)

	case -11:
		pop()

	case -12:
		y := pop()
		x := pop()
		push(y)
		push(x)

	case -13:
		z := pop()
		y := pop()
		x := pop()
		push(y)
		push(z)
		push(x)

	case -14:
		y := pop()
		x := pop()
		push(x)
		push(y)
		push(x)

	case -15:
		x := pop()
		push(stack[x])

	case -16:
		a := pop()
		v := pop()
		stack[v] = a

	case -17:
		y := pop()
		x := pop()
		if x < y {
			push(-1)
		} else if x == y {
			push(0)
		} else {
			push(1)
		}

	case -18:
		CP = pop() - 1
	case -19:
		a := pop()
		x := pop()
		if x < 0 {
			CP = a - 1
		} else {
			//CP = CP + 1
		}
	case -20:
		a := pop()
		x := pop()
		if x > 0 {
			CP = a - 1
		} else {
			//CP = CP + 1
		}
	case -21:
		a := pop()
		x := pop()
		if x == 0 {
			CP = a - 1
		} else {
			//CP = CP + 1
		}
	case -22:
		a := pop()
		x := pop()
		if x <= 0 {
			CP = a - 1
		} else {
			//	CP = CP + 1
		}
	case -23:
		a := pop()
		x := pop()
		if x >= 0 {
			CP = a - 1
		} else {
			//CP = CP + 1
		}
	case -24:
		a := pop()
		x := pop()
		if x != 0 {
			CP = a - 1
		} else {
			//CP = CP + 1
		}
	case -25:
		a := pop()
		push(CP + 1)
		CP = a - 1
	case -26:
		n := pop()
		a := pop()
		for i := 0; i < n; i++ {
			pop()
		}
		CP = a - 1
	case -27:
		push(SP)

	case -28:
		a := pop()
		SP = a

	case -29:
		//fmt.Println("BP", BP)
		push(BP)

	case -30:
		a := pop()
		BP = a

	case -31:
		push(CP)

	case -32:
		os.Exit(0)
		pop()
	case -33:
		var a rune
		fmt.Scanf("%c", &a)
		push(int(a))

	case -34:
		fmt.Print(string(pop()))

	case -35:
		n := pop()
		for i := 0; i < n; i++ {
			pop()
		}
		//SP += n + 1

	case -36:
		n := pop()
		for i := 0; i < n; i++ {
			push(-55)
		}
		//SP -= n - 1
	case -37:
		a := pop()
		CP = a - 1

	default:
		push(x)

	}
}

func get_command_code(x string) int {
	switch x {
	case "ADD":
		return -1
	case "SUB":
		return -2
	case "DIV":
		return -3
	case "MOD":
		return -4
	case "MUL":
		return -5
	case "NEG":
		return -6
	case "BITAND":
		return -7
	case "BITOR":
		return -8
	case "BITNOT":
		return -9
	case "DUP":
		return -10
	case "DROP":
		return -11
	case "SWAP":
		return -12
	case "ROT":
		return -13
	case "OVER":
		return -14
	case "READ":
		return -15
	case "WRITE":
		return -16
	case "CMP":
		return -17
	case "JMP":
		return -18
	case "JLT":
		return -19
	case "JGT":
		return -20
	case "JEQ":
		return -21
	case "JLE":
		return -22
	case "JGE":
		return -23
	case "JNE":
		return -24
	case "CALL":
		return -25
	case "RETN":
		return -26
	case "GOTO": //хз
		return -37
	case "GETSP":
		return -27
	case "SETSP":
		return -28
	case "GETBP":
		return -29
	case "SETBP":
		return -30
	case "GETCP":
		return -31
	case "HALT":
		return -32
	case "IN":
		return -33
	case "OUT":
		return -34
	case "DROPN":
		return -35
	case "PUSHN":
		return -36
	default:
		intVar, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println(err)
		}
		return intVar
	}
}
func push_str_prog(x string) {
	buffer = append(buffer, x)
}
func push(x int) {
	SP--
	stack[SP] = x
}
func pop() int {
	res := stack[SP]
	stack[SP] = -47
	SP++
	return res
}
func InitProgram() []string {
	CP, BP, P = 0, 0, 0
	size := len(os.Args) - 1
	if size == 0 {
		fmt.Println("not file in arguments")
		return []string{}
	}
	if len(os.Args[1]) > 12 && os.Args[1][:12] == "-stack_size=" {
		intVar, err := strconv.Atoi(os.Args[1][12:])
		if err != nil {
			fmt.Println(err)
			return []string{}
		}
		stack = make([]int, intVar)
		SP = intVar
		return os.Args[2:]
	}
	stack = make([]int, 1000)
	SP = 1000
	return os.Args[1:]

}
func GetCode(input_files []string) {
	for i := 0; i < len(input_files); i++ {
		f, err := os.Open(input_files[i])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			cur_string := sc.Text()
			if len(cur_string) > 0 {

				res := regexp.MustCompile(`;`).Split(cur_string, -1)
				if len(res[0]) > 0 {
					s := strings.Fields(res[0])
					for i := 0; i < len(s); i++ {
						push_str_prog(s[i])
					}
				}
			}
		}
	}
}
func calculateConstant() {
	constants := make(map[string]string)
	length := len(buffer)
	for i := 0; i < length; i++ { //собираем
		if buffer[i][0] == ':' {
			res := regexp.MustCompile(`;`).Split(buffer[i][1:], -1)
			copy(buffer[i:], buffer[i+1:])
			buffer[len(buffer)-1] = ""
			buffer = buffer[:len(buffer)-1]
			if len(res) > 1 {
				constants[res[0]] = res[1]
				//replace(res[0], res[1])
			} else {
				constants[res[0]] = strconv.Itoa(i)
				//replace(res[0], strconv.Itoa(i))
			}
			//calculateConstant()
			i--
			length--
		}
	}
	for i := 0; i < len(buffer); i++ { //заменяем
		if val, ok := constants[buffer[i]]; ok {
			buffer[i] = val
		}
	}
}

func generateCode() {
	for i := 0; i < len(buffer); i++ {
		push_asm_prog(get_command_code(buffer[i]))
	}
}
func seestack() {
	fmt.Println("->", CP, "CP", SP, "SP", BP, "BP", P, "P")
	for i := len(stack) - 1; i >= SP; i-- {
		fmt.Print(stack[i], " ")
	}
	fmt.Println()

}
func run_interpreter() {
	i := 0
	for {

		fmt.Println("command ", stack[CP], "len", 1000-SP)
		fmt.Println("all", stack, "until")
		//seestack()
		call_instruction(stack[CP])
		CP++
		fmt.Println("index", i)
		i++
		//fmt.Println("after")
		//seestack()
	}
}
func main() {
	input_files := InitProgram()
	if len(input_files) == 0 {
		fmt.Println("not file in arguments")
		return
	}
	GetCode(input_files)

	//println("\ntest\n")
	calculateConstant()

	generateCode()
	//fmt.Println("run")
	run_interpreter()
	//for i := 0; i < len(buffer); i++ {
	//	fmt.Println(buffer[i])
	//}

}
