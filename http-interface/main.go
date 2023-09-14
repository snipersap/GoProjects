package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("http.Get Error:", err)
	}
	fmt.Println(resp)

	document := make([]byte, 99999)
	readCount, err := resp.Body.Read(document)
	if readCount == 0 && err != nil {
		fmt.Println("Bytes read:", readCount, "Body.Read Error:", err)
	}
	fmt.Println("Google.com >> ", string(document))
}
