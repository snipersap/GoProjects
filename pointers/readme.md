# ReadMe
This repo is the project *pointers* as described in the course  [Building Modern Web Applications with Go (Golang)](https://udemy.com/course/building-modern-web-applications-with-go/). 

This repo defines a color of type string and updates it using the function `myFavColor` where value at operator (`*`) is used to update the color. It also defines a struct, and updates the color in the struct by passing a pointer to th function `myfavDeepColor` but without using the value at operator.

Some experiments were added, like slices and their sorting.

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 13.09.2023  
>Pointers with standard type string >>  
2023/09/24 12:39:07 Current color: Red     
2023/09/24 12:39:07 Fav color: Green       
Pointers with struct>>  
2023/09/24 12:39:07 Current color: Orange  
2023/09/24 12:39:07 Fav color: Blue     

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
