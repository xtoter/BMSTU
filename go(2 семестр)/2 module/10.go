package main

import (
	"fmt"
	"strconv"
)

type Graph struct {
	a          []int
	parent     []int
	namestring string
	totaldist  int
	index      int
	dist       int
	help       int
	T1         int
	low        int
	comp       int
	color      int
}

var time, count = 1, 0

func findMax(graphs []Graph) int {
	out := graphs[0].dist
	for i := 0; i < len(graphs); i++ {
		if graphs[i].dist > out && graphs[i].color != 2 {
			out = graphs[i].dist
		}
	}
	return out
}
func Tarjan(graph *[]Graph) {
	var s []int
	c := *graph
	for i := 0; i < len(*graph); i++ {
		if c[i].T1 == 0 {
			s, _ = visitVertexTarjan(c, i, s)
		}
	}
}
func visitVertexTarjan(graphs []Graph, numV int, s []int) ([]int, []Graph) {
	graphs[numV].T1 = time
	graphs[numV].low = time
	time++
	s = append(s, numV)
	for _, e := range graphs[numV].a {
		if graphs[e].T1 == 0 {
			s, graphs = visitVertexTarjan(graphs, e, s)
		}
		if (graphs[e].comp == -1) && (graphs[numV].low > graphs[e].low) {
			graphs[numV].low = graphs[e].low
		}
	}
	if graphs[numV].low == graphs[numV].T1 {
		u := s[len(s)-1]
		s = s[:len(s)-1]
		if u == numV && getindex(graphs[u].a, u) != -1 {
			coloringblue(&graphs, u)
		}
		graphs[u].comp = count
		for u != numV {
			u = s[len(s)-1]
			s = s[:len(s)-1]
			graphs[u].comp = count
			coloringblue(&graphs, u)
		}
		count++
	}
	return s, graphs
}
func relax(u, v *Graph, help int) bool {
	changed := u.dist+help > v.dist
	if changed {
		v.dist = u.dist + help
		v.parent = make([]int, 0)
		v.parent = append(v.parent, u.totaldist)
	}
	return changed
}

func getindex(slice []int, element int) int {
	for i, x := range slice {
		if x == element {
			return i
		}
	}
	return -1
}
func coloringblue(graphs *[]Graph, i int) {
	(*graphs)[i].color = 2
	(*graphs)[i].dist = -1
	for j := 0; j < len((*graphs)[i].a); j++ {
		e := (*graphs)[i].a[j]
		if (*graphs)[e].color != 2 {
			coloringblue(graphs, e)
		}
	}
}
func coloringred(graphs *[]Graph, i int) {
	(*graphs)[i].color = 1
	for j := 0; j < len((*graphs)[i].parent); j++ {
		e := (*graphs)[i].parent[j]
		if (*graphs)[e].color != 1 {
			coloringred(graphs, e)
		}
	}
}
func equality(a, b string) bool {
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
func output(graphs []Graph) {
	fmt.Println("digraph {")
	color := [2]string{"red", "blue"}
	for i := 0; i < len(graphs); i++ {
		fmt.Printf("%s [label = \"%s(%d)\"", graphs[i].namestring, graphs[i].namestring, graphs[i].help)
		if graphs[i].color > 0 {
			fmt.Printf(", color = %s", color[graphs[i].color-1])
		}
		fmt.Printf("]\n")
	}
	for i := 0; i < len(graphs); i++ {
		for j := 0; j < len(graphs[i].a); j++ {
			fmt.Printf("%s -> %s ", graphs[i].namestring, graphs[graphs[i].a[j]].namestring)
			if graphs[i].color == 2 {
				fmt.Printf("[color = blue]")
			}
			if graphs[i].color == 1 && graphs[graphs[i].a[j]].color == 1 {
				if getindex(graphs[graphs[i].a[j]].parent, graphs[i].totaldist) != -1 {
					fmt.Printf("[color = red]")
				}
			}
			fmt.Printf("\n")
		}
	}
	fmt.Println("}")
}

func getentry(a, b string) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b[0] {
			if equality(a[i:i+len(b)], b) {
				return i
			}
		}
	}
	return -1
}

func getGraph(dist, help int, token string) Graph {
	var out Graph
	out.color = 0
	out.totaldist = dist
	out.namestring = token
	out.help = help
	out.comp = -1
	out.dist = help
	return out
}
func getstring(a string, num int) (string, string, int) {
	var out1, out2 []byte

	if num >= len(a) {
		return string(out1), string(out2), num
	}
	c := a

	i := num
	for i < len(c) && c[i] == ' ' {
		i++
	}
	for i < len(c) && c[i] != ' ' {
		out1 = append(out1, c[i])
		i++
	}
	for i < len(c) && c[i] == ' ' {
		i++
	}
	for i < len(c) && c[i] != ' ' {
		out2 = append(out2, c[i])
		i++
	}
	return string(out1), string(out2), i
}

func inputdata() []Graph {
	//num := 0
	/*scanner := bufio.NewScanner(os.Stdin)
	var source string
	for scanner.Scan() {
		source = fmt.Sprintf("%s %s", source, scanner.Text())
	}*/
	//fmt.Println(source)
	var token1, token2, tokenProm string
	var cond1, cond2 bool
	var graphs []Graph
	was := make(map[string]int)
	fmt.Scanf("%s %s", &token1, &tokenProm)
	//token1, tokenProm, num = getstring(source, num)
	//fmt.Println(token1, "- ", tokenProm)
	firstentry := getentry(token1, ";")
	if firstentry != -1 {
		cond1 = true
		token1 = stringchange(token1, token1[firstentry:], "", -1)
	}

	left := getentry(token1, "(")
	right := getentry(token1, ")")

	help, _ := strconv.ParseInt((token1[left+1 : right]), 10, 0)
	token1 = stringchange(token1, token1[left:], "", -1)
	graphs = append(graphs, getGraph(0, int(help), token1))
	i := 0
	if !(tokenProm == "" && !cond1) {
		for {
			if cond1 {
				cond2 = true
				cond1 = false
			}
			var tokenProm string
			fmt.Scanf("%s %s", &token2, &tokenProm)
			//token2, tokenProm, num = getstring(source, num)

			firstentry = getentry(token2, ";")
			if firstentry != -1 {
				cond1 = true
				token2 = stringchange(token2, token2[firstentry:], "", -1)
			}
			left = getentry(token2, "(")
			if left != -1 {
				right = getentry(token2, ")")

				help, _ = strconv.ParseInt((token2[left+1 : right]), 10, 0)
				token2 = stringchange(token2, token2[left:], "", -1)
				was[token2] = i + 1
				graphs = append(graphs, getGraph(i+1, int(help), token2))
				if !cond2 {
					t := was[token1]
					if getindex(graphs[t].a, i+1) == -1 {
						graphs[t].a = append(graphs[t].a, i+1)
					}
				}
				i++
			} else {
				if !cond2 {
					t := was[token1]
					if getindex(graphs[t].a, was[token2]) == -1 {
						graphs[t].a = append(graphs[t].a, was[token2])
					}

				}
			}
			token1 = token2
			if tokenProm == "" && !cond1 {
				return graphs
			}
			cond2 = false
		}
	}
	return graphs
}
func stringchange(a, b, c string, n int) string {
	var temp []byte
	if n == -1 {
		n = 99999
	}
	i := 0
	for i < len(a) {
		if n > 0 {
			if a[i] != b[0] {
				temp = append(temp, a[i])
				i++
			} else if equality(a[i:i+len(b)], b) {
				for j := 0; j < len(c); j++ {
					temp = append(temp, c[j])

				}

				i = i + len(b)
				n--
			} else {
				temp = append(temp, a[i])
				i++
			}
		} else {
			break
		}

	}
	for i < len(a) {
		temp = append(temp, a[i])
		i++
	}
	return string(temp)
}

func forred(graph *[]Graph) {
	cond := false
	for i, _ := range *graph {
		if (*graph)[i].color != 2 {
			for _, e := range (*graph)[i].a {
				relaxed := relax(&(*graph)[i], &(*graph)[e], (*graph)[e].help)
				if (*graph)[i].dist+(*graph)[e].help == (*graph)[e].dist {
					if getindex((*graph)[e].parent, (*graph)[i].totaldist) == -1 {
						(*graph)[e].parent = append((*graph)[e].parent, (*graph)[i].totaldist)
					}
				}
				if relaxed {
					cond = true
				}
			}
		}
	}
	if cond {
		forred(graph)
	}
}
func main() {
	graphs := inputdata()
	//fmt.Println(graphs)
	//fmt.Println("test")
	//fmt.Println(graphs)
	Tarjan(&graphs)
	//fmt.Println(graphs)
	forred(&graphs)
	//fmt.Println(graphs)
	maxdist := findMax(graphs)
	//fmt.Println(maxdist, graphs)
	//fmt.Println(maxdist)
	//fmt.Println(graphs)
	for i := 0; i < len(graphs); i++ {
		if graphs[i].dist == maxdist && graphs[i].color != 2 {
			coloringred(&graphs, i)
		}
	}
	//fmt.Println(graphs)
	output(graphs)

}

