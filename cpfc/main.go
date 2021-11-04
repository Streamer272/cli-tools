package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 3 {
        fmt.Println("Usage: cpfc <source> <target>")
        os.Exit(1)
    }

	var sourcePath, targetPath string
	sourcePath = os.Args[1]
	targetPath = os.Args[2]

	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		fmt.Printf("Error: %v doesn't exist\n", targetPath)
		os.Exit(1)
	}

	err := copyFolder(sourcePath, targetPath)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}

func copyFolder(sourcePath string, destinationPath string) (err error) {
	sourceDirectory, _ := os.Open(sourcePath)

	sourceDirectoryFiles, err := sourceDirectory.Readdir(-1)

	for _, sourceDirectoryFile := range sourceDirectoryFiles {
		sourceDirectoryFilePath := sourcePath + "/" + sourceDirectoryFile.Name()

		destinationDirectoryFilePath := destinationPath + "/" + sourceDirectoryFile.Name()

		if sourceDirectoryFile.IsDir() {
			err = copyFolder(sourceDirectoryFilePath, destinationDirectoryFilePath)
			if err != nil {
				fmt.Printf("Error while copying directory: %v\n", err)
			}
		} else {
			err = copyFile(sourceDirectoryFilePath, destinationDirectoryFilePath)
			if err != nil {
				fmt.Printf("Error while copying file: %v\n", err)
			}
		}
	}

	return
}

func copyFile(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceInfo.Mode())
		}
	}

	return
}
