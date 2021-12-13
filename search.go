package main

import (
	_ "errors"
	"log"
	"os/exec"
	"runtime"
)

var osCommand string

func init() {
	// Get the command for opening browser
	switch runtime.GOOS {
	case "linux":
		osCommand = "xdg-open"
	case "windows":
		osCommand = "rundll32"
	default:
		log.Fatal("Unsupported platform")
	}
}

type Parserer interface {
	Parse(input string) (string, error)
}

type Search struct {
	parser Parserer
}

// Search opens web browser to an url based on given query string
func (s Search) Search(query string) {
	url, err := s.parser.Parse(query)

	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(osCommand, url)

	cmd.Start()
}
