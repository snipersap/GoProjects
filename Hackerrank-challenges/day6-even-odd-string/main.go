package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	tests := int(scanTestCases())
	texts := scanTestTexts(tests)
	printEvenOddIndexedTexts(texts)

}
func scanTestCases() int64 {
	var tests int64
	var input string

	_, err := fmt.Scan(&input)
	checkErr(err)
	tests, err = strconv.ParseInt(strings.TrimRight(input, "\r\n"), 10, 64)

	return tests
}
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func scanTestTexts(tests int) []string {
	var texts []string
	var input string

	for i := 0; i < tests; i++ {
		_, err := fmt.Scan(&input)
		checkErr(err)
		texts = append(texts, strings.TrimRight(input, "\r\n"))
	}

	return texts
}

func printEvenOddIndexedTexts(texts []string) {

	for _, t := range texts {
		var even, odd []byte
		for j := 0; j < len(t); j++ {
			if j%2 == 0 {
				even = append(even, t[j])
			} else {
				odd = append(odd, t[j])
			}
		}
		fmt.Printf("%s %s\n", string(even), string(odd))
	}

}
