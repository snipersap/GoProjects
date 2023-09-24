# ReadMe
This repo is the project *read-write-json* as described in the course [Building Modern Web Applications with Go (Golang)](https://udemy.com/course/building-modern-web-applications-with-go/). 

This repo reads the json returned from `initJson` function and decodes it into a `Person` struct. It then marshals the `Person` struct into a json. Then finally comapares the initial json with the encoded json and panics if they are not the same. 

## How to test the repo?
Run the files `main.go` using the command 
`go run main.go`.   

Run the tests with 
`go test`

### Test Result:


In case of confusion, check the commits. 

## Expected output
Last updated 13.09.2023  
```
Persons recovered are: [{Raju Gentleman India 10} {Miroslav Guzmann Ukraine 25}]  
Person Marshaled into:  
[  
    {  
                "first_name": "Raju",  
                "last_name": "Gentleman",  
                "country": "India",  
                "age": 10  
        },  
        {  
                "first_name": "Miroslav",  
                "last_name": "Guzmann",  
                "country": "Ukraine",  
                "age": 25  
        }  
]
2023/09/24 20:37:14 Encoded json does not match initial json!  
exit status 1     
```  
  
## Report bugs or suggestions
Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome. 

## Disclaimer
Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result. 
