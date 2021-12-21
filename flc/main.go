package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Please specify at least one file...\n")
		os.Exit(1)
	}

	sum := 0

	for i, file := range os.Args {
		if i == 0 {
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Unexpected error (%v) occurred while reading %v...\n", err, file)
		}

		lineCount := len(strings.Split(string(content), "\n"))
		fmt.Printf("Line count of %v is %v\n", file, lineCount)
		sum += lineCount
	}

	if len(os.Args) > 2 {
		fmt.Printf("Sum of line counts in %v is %v\n", strings.Join(os.Args[1:], ", "), sum)
	}
}
