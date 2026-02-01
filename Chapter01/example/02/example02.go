package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, val := range os.Args[1:] {
		s = s + sep + val
		sep = " "
	}
	fmt.Println(s)
}

// go run example02.go 1 2 3
