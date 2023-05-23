package first

import (
	"2.4/syntaxAnalizer"
	"fmt"
)

var first map[string][]string

type Set map[string]bool

func GetFirst(tree syntaxAnalizer.Tree) {
	first = make(map[string][]string)

	for i := 0; i < len(tree.Leaves); i++ {
		GetFirstRecRule(tree.Leaves[i])
	}

	firstSets := make(map[string]Set)
	for symbol := range first {
		firstSets[symbol] = make(Set)
	}

	for symbol := range first {
		expandFirstSet(symbol, firstSets)
	}

	for symbol, firstSet := range firstSets {
		fmt.Printf("FIRST(%s): ", symbol)
		for terminal := range firstSet {
			fmt.Printf("%s ", terminal)
		}
		fmt.Println()
	}
}

func expandFirstSet(symbol string, firstSets map[string]Set) {
	if len(firstSets[symbol]) > 0 {
		return
	}
	productions := first[symbol]
	for _, production := range productions {
		if production >= "A" && production <= "Z" {
			expandFirstSet(production, firstSets)
			firstSets[symbol] = union(firstSets[symbol], firstSets[production])
		} else {
			firstSets[symbol][production] = true
		}
	}
}

func union(setA, setB Set) Set {
	result := make(Set)
	for key := range setA {
		result[key] = true
	}
	for key := range setB {
		result[key] = true
	}
	return result
}
func getRules(tree syntaxAnalizer.Tree) []string {
	var result []string
	for i := 0; i < len(tree.Leaves); i++ {
		if tree.Leaves[i].Str == "Term" {
			result = append(result, GetFirstRecTerm(tree.Leaves[i])...)
		}
	}
	return result
}
func GetFirstRecRule(tree syntaxAnalizer.Tree) {
	if tree.Str == "Rule" {
		first[tree.Leaves[0].Leaves[0].CurrentToken.Value] = getRules(tree.Leaves[2])

	}
}
func GetFirstRecTerm(tree syntaxAnalizer.Tree) []string {
	var result []string
	for i := 0; i < len(tree.Leaves); i++ {

		subtree := tree.Leaves[i]
		subsubtree := subtree.Leaves[0]
		if subsubtree.CurrentToken.Value != "" {
			result = append(result, subsubtree.CurrentToken.Value)
		} else if subsubtree.Str == "Terminal" {
			result = append(result, subsubtree.Leaves[0].CurrentToken.Value)
		} else if subsubtree.Str == "Option" {
			result = append(result, subsubtree.Leaves[0].Leaves[0].CurrentToken.Value)

		}
	}

	return result
}
