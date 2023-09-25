# ReadMe
This repo is the project *automate-tests* as described in the course [Building Modern Web Applications with Go (Golang)](https://udemy.com/course/building-modern-web-applications-with-go/). 

This repo creates a simple `divide` function that returns an error and the result. Using the `divide` function, we create a few tests that we automate. 

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:
Last updated 13.09.2023  
```
=== RUN   TestDivide
--- PASS: TestDivide (0.00s)
PASS
ok      automate-tests  0.076s
```  
## Coverage
```
PASS
coverage: 33.3% of statements
ok      automate-tests  1.142s
```
### Report
Get coverage report by running the command: 
`go test -coverprofile=coverage.out && go tool cover -html=coverage.out`
[Coverage report from test](automate-tests_%20Go%20Coverage%20Report.html)
In case of confusion, check the commits. 

## Expected output
When y = 0.0
```
2023/09/25 14:41:48 Divide error: divide by zero error
```
OR when y = 20.0
```
Result of dividing 10.000000 by 20.000000 is: 0.500000
```



## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
