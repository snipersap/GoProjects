package main

import "testing"

var tests = []struct {
	testName                  string
	dividend, divisor, expect float32
	isError                   bool
}{
	{"correct-divide", 10, 5, 2.0, false},
	{"incorrect-divide", 10, 0, 0.0, true},
	{"fraction-result", 10, 2.5, 4.0, false},
	{"negative-divide", -10, 5, -2.0, false},
}

func TestDivide(t *testing.T) {

	for _, tt := range tests {

		//Check for errors
		res, err := divide(tt.dividend, tt.divisor)
		if err != nil && !tt.isError {
			t.Errorf("Did not expect an error but got one for test %s ", tt.testName)
		} else if err == nil && tt.isError {
			t.Errorf("Expected an error but didn't get one for test %s ", tt.testName)
		}

		//Check for expected values
		if res != tt.expect {
			t.Errorf("Expected %f but got %f in %s", tt.expect, res, tt.testName)
		}

	}

}
