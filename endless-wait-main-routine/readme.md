# ReadMe
This repo is the project *endless-wait-main-routine* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

This program tests whether the main routine will wait for the returning value of the channel even if no wait to channel has been explicitly provided. This was aked in Q5 of Quiz 11 in the above course. 
> Another tough one!  Is there any issue with the following code?  
>
>package main  
>   
>func main() {  
>     c := make(chan string)  
>     c <- "Hi there!"  
>}

>Answer in Quiz: 
The syntax of this program is OK but the program will never exit because, it will wait for something to receive the value we passed into the channel c. 

## Conclusion
Based on the output in the section *Expected output*, we can conclude that when all go routines are sleeping, we cannot pass any value to the channel. 

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 18.09.2023 
>fatal error: all goroutines are asleep - deadlock!  
>
>goroutine 1 [chan send]:  
>main.main()  
>        .../GoProjects/GoProjects/endless-wait-main-routine/main.go:5 +0x28  
>exit status 2

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
