package main

import "fmt"

var FUNCNAME string = "main"

func testfunc() {
	fmt.Println(FUNCNAME)
}
func main() {
	testfunc()
	fmt.Println(FUNCNAME)
}
