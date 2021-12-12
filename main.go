package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//	Get command-line arguments
	if len(os.Args) < 2 {
		log.Fatal("Missing argument.")
	}
	cmd := os.Args[1]

	fmt.Println(cmd)
}
