# ReadMe
This repo is the project *routines-and-channels* as described in the course  [Go: The Complete Developer's Guide (Golang)](https://udemy.com/course/go-the-complete-developers-guide/). 

The programs implements variations of the go routines, uses endless loops and function literal and sleep, and thus implements an endless web link status checker.

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

### After implementing wait for only 5 channels, although you have 4 links to check
>Yay!! http://google.com is up.   
^^^ UP ^^^   
Maybe http://arunisgood.com is down!   
vvv DOWN vvv   
Yay!! http://amazon.com is up.  
^^^ UP ^^^  
Yay!! http://honeywell.com is up.  
^^^ UP ^^^  
*Here the program waits indefinitely!*

### After implementing wait for as many links we are checking
>Maybe http://arunisgood.com is down!  
vvv DOWN vvv  
Yay!! http://google.com is up.  
^^^ UP ^^^  
Yay!! http://amazon.com is up.  
^^^ UP ^^^  
Yay!! http://honeywell.com is up.  
^^^ UP ^^^  

## Implement link check endlessly
>Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Yay!! http://google.com is up.  
Maybe http://arunisgood.com is down!  
Yay!! http://amazon.com is up.  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  

The above outcome is because the memory of the channel is being referenced by both main and child routines, where the child routine changes the value of the channel c, and the same link gets passed again to the checklink function within the endless loop. In rare cases, the checklink function receives other links to process too. 

## Alternative form of endless loop
>Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Yay!! http://google.com is up.  
Yay!! http://amazon.com is up.  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  
Maybe http://arunisgood.com is down!  

## USe function literal and remove sharing of variables between main and child routines
>Maybe http://arunisgood.com is down!  
Yay!! http://google.com is up.  
Yay!! http://amazon.com is up.  
Yay!! http://honeywell.com is up.  
Maybe http://arunisgood.com is down!  
Yay!! http://google.com is up.  
Yay!! http://amazon.com is up.  
Maybe http://arunisgood.com is down!  
Yay!! http://honeywell.com is up.  
Yay!! http://google.com is up.  
Maybe http://arunisgood.com is down!  
Yay!! http://amazon.com is up.  
Yay!! http://honeywell.com is up.  
Yay!! http://google.com is up.  
Maybe http://arunisgood.com is down!  

The above lines are printed with around 5 seconds of delay.

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
