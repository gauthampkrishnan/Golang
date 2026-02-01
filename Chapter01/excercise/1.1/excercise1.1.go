package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	sep := " "
	for i := 0; i < len(os.Args); i++ {
		s = s + os.Args[i] + sep

	}
	fmt.Println(s)
}

// go run excercise1.1.go 1 2 3
