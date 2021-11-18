package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: measure <command> <args>")
		os.Exit(1)
	}

	var args []string
	if len(os.Args) >= 3 {
        args = os.Args[2:]
    } else {
		args = []string{}
	}

	start := time.Now()

	commandOutput, err := exec.Command(os.Args[1], args...).Output()
	if err != nil {
	   fmt.Printf("Error occurred while executing command %v\n%v\n", os.Args[1], err)
	   os.Exit(1)
	}

	output := string(commandOutput)
	fmt.Printf("%v was executed in time %v with output\n%v\n", os.Args[1], time.Since(start), output)
}
