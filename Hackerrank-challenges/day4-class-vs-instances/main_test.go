package main

import "testing"

var testsNewPerson = []struct {
	testName            string
	input, expNp, expYP int
}{
	{"Test0", -1, 0, 1},
	{"Test1", 10, 10, 11},
	{"Test2", 16, 16, 17},
	{"Test3", 18, 18, 19},
}

func TestPerson(t *testing.T) {

	for _, tt := range testsNewPerson {
		p := &person{}
		p = p.NewPerson(tt.input)
		if p.age != tt.expNp {
			t.Errorf("%s:Expected age to be %d but got %d", tt.testName, tt.expNp, p.age)
		}

		p = p.yearPasses()
		if p.age != tt.expYP {
			t.Errorf("%s:Expected age to be %d but got %d",tt.testName, tt.expYP, p.age)
		}
	}

}
