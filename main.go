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
	message := ""

	for {
		char, err := reader.ReadByte()
		if err != nil {
			break
		}

		message += string(char)
	}

	message = strings.TrimSuffix(message, "\n")

	err := clipboard.WriteAll(message)
	if err != nil {
		fmt.Printf("Error while writing to clipboard\n")
		os.Exit(1)
	}
}
