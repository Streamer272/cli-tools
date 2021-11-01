package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error while reading STDIN\n")
		os.Exit(1)
	}

	message = strings.TrimSuffix(message, "\n")

	err = clipboard.WriteAll(message)
	if err != nil {
		fmt.Printf("Error while writing to clipboard\n")
		os.Exit(1)
	}
}
