# ReadMe
This repo is the project *http-interface* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

The program calls google.com and prints the following: 
1. The response from google.com without the body
2. The body of the response using `r.Body.Read` function
3. The body of the response using `io.Copy` function
4. The body of the response using `ìo.Copy` function and passing a custom `Write` implementation using the `logWriter` struct type.

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 14.09.2023  
>*prints response from google.com*  
 Google.com >> *prints body from google.com*  
 Print Body using io.Copy() >>  
 *prints body from google.com*  
 , Written: 18453  
 Print Body using log writer() >>  
From custom Write implementation with 18379 bytes >>  
*prints body of google.com*  
, Wrote: 18379 bytes  


In case we use ftp instead of http in get: 
>2023/09/23 10:22:48 http.Get Error:Get "ftp://google.com": unsupported protocol scheme "ftp"
>exit status 1

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
