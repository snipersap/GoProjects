package main

import (
	"fmt"
	"time"
)

func main() {
	//send a birthday message
	birthdayMsg := birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "Donald Duck",
	}
	sendMessage(birthdayMsg)

	//send a report
	reportMsg := sendingReport{
		reportName:    "Communication insights",
		numberOfSends: 400,
	}
	sendMessage(reportMsg)
}

func sendMessage(msg message) (string, int) {
	contentOfMessage := msg.getMessage()
	costOfMessage := len(contentOfMessage) * 3
	fmt.Println("Message:", contentOfMessage, ", and it cost you (in cents):", costOfMessage) //output to console.
	return contentOfMessage, costOfMessage
}

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}
