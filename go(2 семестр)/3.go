package main

import (
	"fmt"

	"github.com/skorobogatov/input"
)

func main() {
	s := input.Gets()
	curstring := []rune(s)
	input.Scanf("\n")
	atemp := input.Gets()
	a := []rune(atemp)
	var a1, b1, len1, len2 int
	a1, b1, len1 = -1, -1, 1000001
	//fmt.Print(curstring[4], a[0], a[2])
	for i, val := range curstring {
		if val == a[0] {
			a1 = i
			if b1 != -1 {
				len2 = a1 - b1 - 1
				if len2 < len1 {
					len1 = len2
				}
			}
		}
		if val == a[2] {
			b1 = i
			if a1 != -1 {
				len2 = b1 - a1 - 1
				if len2 < len1 {
					len1 = len2
				}
			}
		}
	}
	fmt.Print(len1)
}
