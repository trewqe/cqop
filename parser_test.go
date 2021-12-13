package main

import (
	"testing"
)

var parser Parser

func init() {
	parser = Parser{}
}

func TestParse(t *testing.T) {
	// Arrange
	tests := []struct {
		input  string
		expect string
	}{
		{input: "ggl/test", expect: "https://google.com/search?q=test"},
		{input: "ggl/test space ing", expect: "https://google.com/search?q=test+space+ing"},
		{input: "ggl/test   space     mult", expect: "https://google.com/search?q=test+++space+++++mult"},
		{input: "ytb/test", expect: "https://youtube.com/results?search_query=test"},
		{input: "ytb/test space ing", expect: "https://youtube.com/results?search_query=test+space+ing"},
		{input: "ytb/test   space     mult", expect: "https://youtube.com/results?search_query=test+++space+++++mult"},
	}

	for _, test := range tests {
		// Act
		got, _ := parser.Parse(test.input)

		// Assert
		if got != test.expect {
			t.Errorf("Parse(%s): got: %s, expected: %s", test.input, got, test.expect)
		}
	}
}

func TestParseErrors(t *testing.T) {
	// Arrange
	tests := []string{
		"invalid/site",
		"invalid format",
	}

	for _, test := range tests {
		// Act
		_, err := parser.Parse(test)

		//Assert
		if err == nil {
			t.Errorf("Parse(%s) error expected.", test)
		}
	}
}
