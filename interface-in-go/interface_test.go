package main

import "testing"

func TestGreeting(t *testing.T) {
	// Setup
	eb := englishBot{}
	db := deutscheBot{}
	expectedEbGreeting := "Hello there!"
	expectedDbGreeting := "Halli Hallo!"

	// Execute
	ebGreeting := eb.greeting()
	dbGreeting := db.greeting()

	if ebGreeting != expectedEbGreeting {
		t.Errorf("Expected englishBot greeting to be %s, but got %s", expectedEbGreeting, ebGreeting)
	}

	if dbGreeting != expectedDbGreeting {
		t.Errorf("Expected deutsdcheBot greeting to be %s, but got %s", expectedDbGreeting, dbGreeting)
	}

	// Teardown
	ebGreeting = ""
	dbGreeting = ""
	expectedEbGreeting = ""
	expectedDbGreeting = ""

}
