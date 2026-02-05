package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {

		fmt.Print(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print("Error", err)
			os.Exit(1)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Print("Error", err)
			os.Exit(1)
		}
		fmt.Print(string(body))
		fmt.Print(resp.StatusCode)
	}
}
