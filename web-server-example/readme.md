# ReadMe
This repo is the project *web-server-example* as described in the course  [Design Patterns in Go](https://udemy.com/course/design-patterns-go/). 

This repo demonstrates the builder pattern in go, with its various types. 

## How to test the repo?
1. Run the files `main.go` using the command `go run main.go`.
2. Go to your web browser and put in the URL `http:\\localhost:8080`.

Run the tests with 
`go test`

### Test Result:
 

In case of confusion, check the commits. 

## Expected output
### Web browser
#### Home page
About >> is a link to the about page.
```
Hello World!, You are looking at the home page. About >>
```
#### About page
Home >> is a link to the Home page
```
Now we are in the About page.Home >> Result of addition of 2 numbers is:6
```

### Console
Two responses for / and one for the About page.
```
2023/09/25 18:05:17 Starting web server on port: :8080
2023/09/25 18:05:33 Bytes written to response: 90
2023/09/25 18:05:34 Bytes written to response: 90
2023/09/25 18:05:34 Bytes written to response: 90
```
## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
