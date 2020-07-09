package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// for loop with a "blank identifier" ( _ ), since go doesn't permit unused local variables
	for _, arg := range os.Args[1:] {
		// fmt.Println("index:", i, "value:", arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
