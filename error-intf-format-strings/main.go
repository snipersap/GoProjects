package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	// return formatted error string
	eS := fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent", cost, recipient)
	return eS
}
