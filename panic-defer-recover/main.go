package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	//handle any panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)

			//now exit using log.Fatal
			log.Fatal("Exiting program with a fatal error log.")

			//the below print statement is not executed
			fmt.Println("Statement after Fatal error logged.")
		}
	}()

	//trigger the panic
	triggerPanic()
}

func triggerPanic() {
	fmt.Println("I am here to trigger a panic. Triggering it ...")
	err := errors.New("Panic was triggerred.")
	panic(err)
}
