package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	
	// when you don't care about the format but just want to see output, perhaps in debugging
	// any slice may be printed this way
	// fmt.Println(os.Args[0:])
}
