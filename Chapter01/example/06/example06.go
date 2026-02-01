package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		for line, n := range counts {
			if n > 1 {
				println(line, n)
			}
		}

	}
}
