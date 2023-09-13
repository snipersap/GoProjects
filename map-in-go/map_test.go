package main

import "testing"

func TestUpdateColors(t *testing.T) {
	//Setup
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
	}
	expectedColorKey := "turquiose"
	expectedHexCode := "#00ffef"
	// Execute
	updateColors(colors)
	if colors[expectedColorKey] != expectedHexCode {
		t.Errorf("Expected color code of %s as %s, but found %s", expectedColorKey, expectedHexCode, colors[expectedColorKey])
	}
}
