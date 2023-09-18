# ReadMe
This repo is the project *routines-and-channels* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 



## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 18.09.2023 
### Before using the go routine
>Yay!! http://google.com is up.  
Yay!! http://amazon.com is up.  
Yay!! http://honeywell.com is up.  
Maybe http://arunisgood.com is down!  

### After implementing the go routine, but without channels
> No output.

### After implementing wait for only 1 channel
> Maybe http://arunisgood.com is down!   
vvv DOWN vvv   

### After implementing wait for only 2 channels
>Maybe http://arunisgood.com is down!   
vvv DOWN vvv   
Yay!! http://google.com is up.   
^^^ UP ^^^   

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
