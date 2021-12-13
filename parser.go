package main

import (
	"errors"
	"strings"
)

type Parser struct{}

// Parse parses input string into a corresponding url
// input string should be in format "site/query"
func (p Parser) Parse(input string) (string, error) {
	split := strings.SplitN(input, "/", 2)
	var res string

	switch split[0] {
	case "ggl":
		res = "https://google.com/search?q=" + parseQuery(split[1])
	case "ytb":
		res = "https://youtube.com/results?search_query=" + parseQuery(split[1])
	default:
		return "", errors.New("Unsupported site.")
	}

	return res, nil
}

func parseQuery(query string) string {
	return strings.ReplaceAll(query, " ", "+")
}
