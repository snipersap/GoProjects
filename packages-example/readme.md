# ReadMe
This repo is the project *pacakges-example* as described in the course  [Building Modern Web Applications with Go (Golang)](https://udemy.com/course/building-modern-web-applications-with-go/). 

This repo demonstrates the use of custom pacakges in go. The `helpers` package defines a structure `Animal` which contains multiple public fields and one private field. It also defines a public function `PrintInfo`, which then is called from `main` and adds a scientific name based on the breed of the animal. 

Important was to define a folder and another go file in the folder, that contained the package and related structure and function. We import this package into our `main.go` file and can then consume it.

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 13.09.2023  
>Name: Shizu  
Breed: Dog  
Number of Logs: 4          
Science calls it a: Canis    

## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
