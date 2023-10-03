package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input string
var sizeOf int64
var intArr []int64

func main() {

	getUserInput()
	printReverseArr()

}

// printReverseArr loops over the integer array in reverse and prints it
func printReverseArr() {
	fmt.Println("The integer array printed in reverse (upto max size given) is:")
	for i := sizeOf - 1; i >= 0; i-- {
		fmt.Printf("%d ", intArr[i])
	}
}

// getUserInput gets the size of the array and the integer array as space separated values
func getUserInput() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	fmt.Println("Enter the size of the integer array:")
	ntemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	sizeOf = ntemp

	fmt.Print("Enter the space separated list of integers:")
	splitArr := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	for i, s := range splitArr {
		if i >= int(sizeOf) {
			break
		}
		itemp, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			intArr = append(intArr, itemp)
		}
	}
}

// checkError checks for the given error and exits with code 1 if there is an error
func checkError(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// readLine reads the user input and trims new lines from it
func readLine(reader *bufio.Reader) string {
	by, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return strings.TrimRight(string(by), "\r\n")
}
