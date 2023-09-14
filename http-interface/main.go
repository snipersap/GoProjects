package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	// Call google.com and print the response
	resp := callGoogle()
	fmt.Println(resp)
	printBodyLn(resp)
	resp = callGoogle()
	printBodyStd(resp)
	resp = callGoogle()
	printLog(resp)

}

// Call google.com and return response
func callGoogle() *http.Response {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("http.Get Error:", err)
		os.Exit(1)
	}
	return resp
}

// Print the body using println statements
func printBodyLn(r *http.Response) {
	document := make([]byte, 32*1024)
	readCount, err := r.Body.Read(document)
	if readCount == 0 && err != nil {
		fmt.Println("Bytes read:", readCount, "Body.Read Error:", err)
	}
	fmt.Println("Google.com >> ", string(document))
}

// Print response body using standard go functions
func printBodyStd(r *http.Response) {
	fmt.Println("\nPrint Body using io.Copy() >>")
	written, copyErr := io.Copy(os.Stdout, r.Body)
	if written == 0 || copyErr != nil {
		fmt.Println("\nCopy error: ", copyErr)
	} else {
		fmt.Println(", Written:", written)
	}
}

// Print response body using custom Log writer
func printLog(r *http.Response) {
	fmt.Println("\nPrint Body using log writer() >>")
	lw := logWriter{}
	written, copyErr := io.Copy(lw, r.Body)
	if written == 0 || copyErr != nil {
		fmt.Println("\nCustom Write error: ", copyErr)
	} else {
		fmt.Println(", Wrote:", written, "bytes")
	}

}

// Custom write implementation
func (lw logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println("\nFrom custom Write implementation with", len(bs), "bytes >>")
	fmt.Println(string(bs))
	return len(bs), nil //Follows spec of Writer interface
}
