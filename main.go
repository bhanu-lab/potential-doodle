package main

import (
	"fmt"
	"os"
	"potential-doodle/events"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Printf("no options are specified to process. check with -h option")
	}

	if argsWithoutProg[0] == "-d" {
		events.DisplayEvents()
	} else if argsWithoutProg[0] == "-h" {
		fmt.Printf("\t-d\tDisplay next 10 upcoming events \n")
	}
}
