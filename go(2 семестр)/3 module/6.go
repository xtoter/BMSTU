package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

var alphabetglobal []string

type perehod struct {
	next  int
	label string
}
type perehod2 struct {
	old     []bool
	new     []bool
	Markers map[string]bool
}
type perehod3 struct {
	old     []int
	new     []int
	Markers []string
}

func Dfs(fortunes [][]perehod, q perehod, C map[int]bool, markers map[string]bool) map[string]bool {
	//if C[q.next] == false {
	C[q.next] = true
	markers[q.label] = true
	for i := 0; i < len(fortunes[q.next]); i++ {
		if fortunes[q.next][i].label == "lambda" {
			Dfs(fortunes, fortunes[q.next][i], C, markers)
		}
	}
	//}
	return markers
}

func Closurenew(fortunes [][]perehod, z int) (map[int]bool, map[string]bool) {
	C := make(map[int]bool)
	markers := make(map[string]bool)
	for i := 0; i < len(fortunes[z]); i++ {
		if fortunes[z][i].label != "lambda" {
			markers = Dfs(fortunes, fortunes[z][i], C, markers)
		}
	}
	return C, markers
}
func ClosureArray(fortunes [][]perehod, z []bool) (map[int]bool, map[string]bool) {
	C := make(map[int]bool)
	markers := make(map[string]bool)
	for j := 0; j < len(z); j++ {
		if z[j] == true {
			for i := 0; i < len(fortunes[j]); i++ {
				if fortunes[j][i].label != "lambda" {
					markers = Dfs(fortunes, fortunes[j][i], C, markers)
				}
			}
		}
	}
	return C, markers
}

/*func Closure(fortunes [][]perehod, z int) map[int]bool {
	C := make(map[int]bool)
	for i := 0; i < len(fortunes[z]); i++ {
		Dfs(fortunes, fortunes[z][i], C)
	}
	return C
}*/
func MaptoArray(a map[int]bool, N int) []bool {
	var b []bool
	for i := 0; i < N; i++ {
		b = append(b, a[i])
	}
	return b
}
func ArraytoString(a []bool) string {
	var string string
	if len(a) > 0 {
		if a[0] == true {
			string = fmt.Sprint(string, 1)
		} else {
			string = fmt.Sprint(string, 0)
		}
		if len(a) > 1 {
			for i := 1; i < len(a); i++ {
				if a[i] == true {
					string = fmt.Sprint(string, 1)
				} else {
					string = fmt.Sprint(string, 0)
				}
			}
		}
	}
	return string
}
func containsMarkers(a, b map[string]bool) bool {
	for key, _ := range a {
		if b[key] == true {
			return true
		}
	}
	return false
}
func UnionMarkers(a, b perehod2) perehod2 {
	var out perehod2
	out.old = a.old
	for i := 0; i < len(a.new); i++ {
		if a.new[i] == true {
			out.new = append(out.new, true)
		} else {
			if b.new[i] == true {
				out.new = append(out.new, true)
			} else {
				out.new = append(out.new, false)
			}
		}
	}
	for key, _ := range a.Markers {
		b.Markers[key] = true
	}
	out.Markers = b.Markers
	return out
}
func Det(fortunes [][]perehod, Final []bool, q int, N int) ([][]bool, []perehod2) {
	temp4 := MaptoArray(Closurelambda(fortunes, q), N)
	q0 := temp4
	Q := make(map[string]bool)
	Q[ArraytoString(q0)] = true
	var delta []perehod2
	var F [][]bool
	var stack [][]bool

	//fmt.Println(temp4)
	//fmt.Println(q0)

	stack = append(stack, q0)
	for len(stack) > 0 {
		z := stack[len(stack)-1]
		//fmt.Println("z", z)
		stack = stack[:len(stack)-1]
		for i := 0; i < len(z); i++ {
			u := z[i]

			if u && Final[i] {
				//fmt.Println(z)
				F = append(F, z)
				break
			}
		}
		//var stacktemp [][]bool
		var deltatemp []perehod2
		for i := 0; i < N; i++ {
			if z[i] {
				first, second := Closurenew(fortunes, i)
				zz := MaptoArray(first, N)
				/*if Q[ArraytoString(zz)] == false {
					Q[ArraytoString(zz)] = true
					//fmt.Println(zz)
					stacktemp = append(stacktemp, zz)
					//fmt.Println(zz, second)
				}*/
				var temp perehod2
				temp.old = z
				temp.new = zz
				temp.Markers = second
				if len(temp.Markers) > 0 {
					deltatemp = append(deltatemp, temp)
				}
			}

		}
		for i := 0; i < len(deltatemp); i++ {
			for j := 0; j < len(deltatemp); j++ {
				if i != j {
					if containsMarkers(deltatemp[i].Markers, deltatemp[j].Markers) {
						deltatemp[i] = UnionMarkers(deltatemp[i], deltatemp[j])
						deltatemp[j] = deltatemp[len(deltatemp)-1]
						deltatemp = deltatemp[:len(deltatemp)-1]

					}
				}
			}
		}
		//fmt.Println(deltatemp)

		for i := 0; i < len(deltatemp); i++ {
			delta = append(delta, deltatemp[i])
		}
		for i := 0; i < len(deltatemp); i++ {
			zz := deltatemp[i].new
			if Q[ArraytoString(zz)] == false {
				Q[ArraytoString(zz)] = true
				stack = append(stack, zz)
			}

		}
	}
	return F, delta
}

func FindNumber(a []bool) []int {
	var out []int
	if len(a) > 0 {
		for i := 0; i < len(a); i++ {
			if a[i] == true {
				out = append(out, i)
			}
		}

	}
	return out
}
func MaptoArrayString(input map[string]bool) []string {
	var out []string
	for k, _ := range input {
		if k != "lambda" {
			out = append(out, k)
		}
	}
	return out
}
func MaptoArrayStringtrue(input map[string]bool) []string {
	var out []string
	for k, cond := range input {
		if cond {
			out = append(out, k)
		}
	}
	return out
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true

}
func Dfslambda(fortunes [][]perehod, q int, C map[int]bool, markers map[string]bool) map[string]bool {
	//if C[q.next] == false {
	C[q] = true
	for i := 0; i < len(fortunes[q]); i++ {
		if fortunes[q][i].label == "lambda" {
			Dfslambda(fortunes, fortunes[q][i].next, C, markers)
		}
	}
	//}
	return markers
}
func Closurelambda(fortunes [][]perehod, z int) map[int]bool {
	C := make(map[int]bool)
	markers := make(map[string]bool)
	Dfslambda(fortunes, z, C, markers)

	return C
}
func Unionstrings(b, a []string) []string {
	var result []string
	for i := 0; i < len(a); i++ {
		result = append(result, a[i])
	}
	for i := 0; i < len(b); i++ {
		d := true
		for j := 0; j < len(a); j++ {
			if b[i] == a[j] {
				d = false
			}

		}
		if d {
			result = append(result, b[i])
		}
	}
	return result
}
func sortString(strings []string) []string {
	var temp []string
	for i := 0; i < len(alphabetglobal); i++ {
		for j := 0; j < len(strings); j++ {
			if strings[j] == alphabetglobal[i] {
				temp = append(temp, strings[j])
			}
		}
	}
	return temp
}
func Canonize(accept [][]bool, perehod []perehod2) ([]perehod3, [][]int) {
	var newperehod []perehod3
	var newaccept [][]int
	for i := 0; i < len(accept); i++ {
		newaccept = append(newaccept, FindNumber(accept[i]))
	}
	for i := 0; i < len(perehod); i++ {
		var temp perehod3
		temp.old = FindNumber(perehod[i].old)
		temp.new = FindNumber(perehod[i].new)
		temp.Markers = MaptoArrayString(perehod[i].Markers)
		temp.Markers = sortString(temp.Markers)
		newperehod = append(newperehod, temp)
	}
	n := len(newperehod)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {

			if equal(newperehod[i].old, newperehod[j].old) {
				if equal(newperehod[i].new, newperehod[j].new) {
					newperehod[i].Markers = Unionstrings(newperehod[i].Markers, newperehod[j].Markers)
					newperehod[i].Markers = sortString(newperehod[i].Markers)
					newperehod[j] = newperehod[len(newperehod)-1]
					newperehod = newperehod[:len(newperehod)-1]
				}
			}
		}
	}
	return newperehod, newaccept
}

type perehod4 struct {
	old     int
	new     int
	Markers []string
}

func ArraytoInt(a []int) int {
	result := 0
	desat := 1
	for i := 0; i < len(a); i++ {
		result = result + a[i]*desat
		desat = desat * 10
	}
	return result
}
func Canonizenumber(perehod []perehod3, accept [][]int) (map[int]int, int, [][]int) {
	C := make(map[int]int)
	var temp [][]int
	count := 1
	for i := 0; i < len(perehod); i++ {
		if C[ArraytoInt(perehod[i].old)] == 0 {
			C[ArraytoInt(perehod[i].old)] = count
			temp = append(temp, perehod[i].old)
			count++
		}
		if C[ArraytoInt(perehod[i].new)] == 0 {
			C[ArraytoInt(perehod[i].new)] = count
			temp = append(temp, perehod[i].new)
			count++
		}
	}
	return C, count, temp
}

func output(perehod []perehod3, accept [][]int, alphabet1 []string) {
	fmt.Printf("digraph {\nrankdir = LR\ndummy [label = \"\", shape = none]\n")
	out1, out2, out3 := Canonizenumber(perehod, accept)
	out2--
	for i := 0; i < len(out3); i++ {
		z := out1[ArraytoInt(out3[i])]
		cond := false
		for j := 0; j < len(accept); j++ {
			if z == out1[ArraytoInt(accept[j])] {
				cond = true
			}
		}
		z--
		if cond {
			fmt.Printf("%d  [label = \"", z)
			fmt.Print(out3[i])
			fmt.Printf("\", shape = doublecircle]\n")
		} else {
			fmt.Printf("%d  [label = \"", z)
			fmt.Print(out3[i])
			fmt.Printf("\", shape = circle]\n")
		}

	}

	checkMarker := make(map[string]bool)
	for i := 0; i < len(out3); i++ {
		for j := 0; j < len(alphabet1); j++ {
			checkMarker[fmt.Sprint(alphabet1[j], i)] = true
		}

	}
	fmt.Println("dummy -> 0")
	for i := 0; i < len(perehod); i++ {
		if len(perehod[i].Markers) > 0 {
			fmt.Printf("%d -> %d [label = \"", out1[ArraytoInt(perehod[i].old)]-1, out1[ArraytoInt(perehod[i].new)]-1)

			fmt.Print(perehod[i].Markers[0])
			checkMarker[fmt.Sprint(perehod[i].Markers[0], out1[ArraytoInt(perehod[i].old)]-1)] = false
			for j := 1; j < len(perehod[i].Markers); j++ {
				fmt.Print(", ", perehod[i].Markers[j])
				checkMarker[fmt.Sprint(perehod[i].Markers[j], out1[ArraytoInt(perehod[i].old)]-1)] = false
			}

			fmt.Print("\"]\n")
		}
	}
	test := 0
	for i := 0; i < len(out3); i++ {

		z := out1[ArraytoInt(out3[i])]
		var outresult []string
		for j := 0; j < len(alphabet1); j++ {
			if checkMarker[fmt.Sprint(alphabet1[j], z-1)] {
				outresult = append(outresult, alphabet1[j])
			}
		}
		if len(outresult) > 0 {
			fmt.Printf("%d -> %d [label = \"", z-1, out2)
			test++
			fmt.Print(outresult[0])
			for j := 1; j < len(outresult); j++ {
				fmt.Print(", ", outresult[j])
			}
			fmt.Print("\"]\n")
		}

	}
	if test > 0 {
		fmt.Printf("%d -> %d [label = \"", out2, out2)
		fmt.Print(alphabetglobal[0])
		for j := 1; j < len(alphabetglobal); j++ {
			fmt.Print(", ", alphabetglobal[j])
		}
		fmt.Print("\"]\n")
		fmt.Printf("%d  [label = \"[]\", shape = circle]\n", out2)
	}

	//fmt.Println(checkMarker)
	fmt.Print("}\n")
}
func main() {
	var N, M int
	input.Scanf("%d%d", &N, &M)
	var fortunes [][]perehod
	for i := 0; i < N; i++ {
		var temp []perehod
		fortunes = append(fortunes, temp)
	}
	alphabet := make(map[string]bool)
	for i := 0; i < M; i++ {
		var newfortunes int
		var temp1 int
		var temp2 string
		input.Scanf("%d%d%s", &newfortunes, &temp1, &temp2)
		if temp2 != "lambda" {
			if !alphabet[temp2] {
				alphabetglobal = append(alphabetglobal, temp2)
			}
			alphabet[temp2] = true

		}
		var newperehod perehod
		newperehod.label = temp2
		newperehod.next = temp1
		fortunes[newfortunes] = append(fortunes[newfortunes], newperehod)
	}
	alphabet1 := MaptoArrayString(alphabet)
	alphabet1 = sortString(alphabet1)
	var Final []bool
	for i := 0; i < N; i++ {
		var temp int
		input.Scanf("%d", &temp)
		if temp == 1 {
			Final = append(Final, true)
		} else {
			Final = append(Final, false)
		}
	}
	var q0 int
	input.Scanf("%d", &q0)
	//fmt.Println(N, " ", M, " ", fortunes, " ", Final, " ", q0)
	//fmt.Println(Closurenew(fortunes, q0))
	//fmt.Println(MaptoArray(Closure(fortunes, q0), N))
	//fmt.Println(ArraytoString(MaptoArray(Closure(fortunes, q0), N)))
	//fmt.Println("newdata")
	out1, out2 := Det(fortunes, Final, q0, N)
	//fmt.Println(out1, out2)
	/*for i := 0; i < len(out1); i++ {
		fmt.Println(FindNumber(out1[i]))
	}
	//fmt.Println("new")
	for i := 0; i < len(out2); i++ {
		fmt.Println(FindNumber(out2[i].old), "->", FindNumber(out2[i].new), "marker", out2[i].Markers)
	}*/
	newout1, newout2 := Canonize(out1, out2)
	output(newout1, newout2, alphabet1)
}

