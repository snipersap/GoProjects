# ReadMe
This repo is the project *interface-in-go* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

The program implements the *bot* interface with the function *greeting( )* that returns a value of type *string*. Such that the *struct* types *englishBot* and *deutscheBot* both get access to the interface function, and therefore can be passed to the *print( )* function, which can now accept both *englishBot* and *deutschBot* types. 

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:
PASS  
ok      interface-in-go 1.271s  

In case of confusion, check the commits. 

## Expected output
Last updated 13.09.2023  
>Hello there!  
Halli Hallo!  

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
