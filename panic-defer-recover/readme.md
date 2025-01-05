# Panic

As we've seen, the proper way to handle errors in Go is to make use of the error interface. Pass errors up the call stack, treating them as normal values:
```go
func enrichUser(userID string) (User, error) {
    user, err := getUser(userID)
    if err != nil {
        // fmt.Errorf is GOATed: it wraps an error with additional context
        return User{}, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}
```
However, there is another way to deal with errors in Go: the [`panic`](https://golang.org/ref/spec#Handling_panics) function. When a function calls panic, the program crashes and prints a stack trace.

As a general rule, do not use panic!

The `panic` function yeets control out of the current function and up the call stack until it reaches a function that [defers a `recover`](https://go.dev/blog/defer-panic-and-recover). If no function calls `recover`, the goroutine (often the entire program) crashes.
```go
func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
```
Sometimes new Go developers look at panic/recover, and think, "This is like try/catch! I like this!" Don't be like them.

I use error values for all "normal" error handling, and if I have a truly unrecoverable error, I use [log.Fatal](https://golang.org/pkg/log/#Fatal) to print a message and exit the program.

## Assignment
Write a function in go to trigger a panic, and handle it gracefully by calling the `recover` function within a deferred function. After recovering, log a fatal error using `log.Fatal` function to exit the program. Provide one print statement after logging the fatal error, to show that the program exits after logging the fatal error. 

### Test results
```bash
$> go run .\main.go
I am here to trigger a panic. Triggering it ...
Recovered from error: Panic was triggerred.
2025/01/05 18:45:01 Exiting program with a fatal error log.
exit status 1
```