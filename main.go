package main

import (
	"log"
	"os"
)

func main() {
	//	Get command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("Missing argument.")
	}
	cmd := os.Args[1]

	parser := Parser{}
	search := Search{
		parser: parser,
	}

	search.Search(cmd)
}
