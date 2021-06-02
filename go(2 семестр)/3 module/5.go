package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func getString(a int, b string) string {
	return fmt.Sprint(a, ",", b)
}
func getMealyview(automat Myre) Myre {
	var newi, newj []int
	for i := 0; i < automat.N; i++ {
		for j := 0; j < automat.K1; j++ {
			cond := false
			for z := 0; z < len(newi); z++ {
				if (automat.perehod[newi[z]*automat.K1+newj[z]] == automat.perehod[i*automat.K1+j]) && (automat.exits[newi[z]*automat.K1+newj[z]] == automat.exits[i*automat.K1+j]) {
					cond = true
					break
				}
			}
			if !cond {
				newi = append(newi, i)
				newj = append(newj, j)
			}
		}
	}
	for i := 0; i < len(newi); i++ {
		for j := 0; j < automat.K1; j++ {
			automat.newperehod = append(automat.newperehod, 0)
			automat.newexits = append(automat.newexits, "")
		}
	}
	for i := 0; i < len(newi); i++ {
		for j := 0; j < automat.K1; j++ {
			automat.newperehod[i*automat.K1+j] = automat.perehod[newi[i]*automat.K1+newj[i]]
			automat.newexits[i*automat.K1+j] = automat.exits[newi[i]*automat.K1+newj[i]]
		}
	}
	automat.n = len(newi)
	return automat
}

type Myre struct {
	K1, N, n   int
	perehod    []int
	exits      []string
	newperehod []int
	newexits   []string
	values     []string
	alphabet   map[string]int
	K1elem     []string
}

func Print(automat Myre) {
	fmt.Printf("digraph {\nrankdir = LR\n")
	for i := 0; i < automat.K1*automat.N; i++ {
		c := automat.alphabet[getString(automat.perehod[i], automat.exits[i])]
		fmt.Printf("%d [label = \"(%s)\"]\n", c, automat.values[c])

	}
	for i := 0; i < automat.n; i++ {
		for j := 0; j < automat.K1; j++ {
			fmt.Printf("%d -> %d [label = \"%s\"]\n", automat.alphabet[getString(automat.newperehod[i*automat.K1+j], automat.newexits[i*automat.K1+j])], automat.alphabet[getString(automat.perehod[automat.newperehod[i*automat.K1+j]*automat.K1+j], automat.exits[automat.newperehod[i*automat.K1+j]*automat.K1+j])], automat.K1elem[j])
		}
	}

	fmt.Printf("}\n")
}
func main() {
	var automat Myre
	var K1, K2, N, temp int
	var temp1 string
	input.Scanf("%d", &K1)
	var K2elem []string
	for i := 0; i < K1; i++ {
		input.Scanf("%s", &temp1)
		automat.K1elem = append(automat.K1elem, temp1)
	}
	input.Scanf("%d", &K2)
	for i := 0; i < K2; i++ {
		input.Scanf("%s", &temp1)
		K2elem = append(K2elem, temp1)
	}

	input.Scanf("%d", &N)
	for i := 0; i < K1*N; i++ {
		input.Scanf("%d", &temp)
		automat.perehod = append(automat.perehod, temp)
	}
	for i := 0; i < K1*N; i++ {
		input.Scanf("%s", &temp1)
		automat.exits = append(automat.exits, temp1)
	}
	automat.K1 = K1
	automat.N = N
	automat = getMealyview(automat)
	automat.alphabet = make(map[string]int)

	for i := 0; i < K1*N; i++ {
		automat.alphabet[getString(automat.perehod[i], automat.exits[i])] = i
		automat.values = append(automat.values, getString(automat.perehod[i], automat.exits[i]))

	}
	//fmt.Println(alphabet)
	Print(automat)
	//fmt.Print(K1elem, K2elem, N, perehod, exits)
}

