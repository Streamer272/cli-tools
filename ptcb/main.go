package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Printf("Error (%v) occurred while reading from clipboard\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v", text)
}
