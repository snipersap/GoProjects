package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file := openFile(os.Args[1])
	readFile(file)

}

func readFile(file *os.File) {
	bytes, err := io.Copy(os.Stdout, file)
	if bytes != 0 && err != nil {
		fmt.Println("File", file.Name(), "could not be read. Error:", err)
		os.Exit(1)
	} else {
		fmt.Println("\n<< Read", bytes, "bytes from", file.Name())
	}
}

func openFile(fName string) *os.File {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Println("File open error:", err)
		os.Exit(1)
	} else {
		fmt.Println("File", file.Name(), "opened for reading >>")
	}
	return file

}
