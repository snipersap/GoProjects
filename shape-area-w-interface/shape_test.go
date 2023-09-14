package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	// Setup
	s := square{181}
	tr := triangle{191, 200}
	expArea := float64(181 * 181)
	expTrArea := float64(191 * 200 * 0.5)

	// Execute
	a := s.area()
	if a != expArea {
		t.Errorf("Expected area of square as %v, got %v", expArea, a)
	}

	at := tr.area()
	if at != expTrArea {
		t.Errorf("Expected area of triangle as %v, got %v", expTrArea, at)
	}

}
