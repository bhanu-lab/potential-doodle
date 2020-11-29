package main

import (
	"fmt"
	"os"
	"potential-doodle/events"

	"google.golang.org/api/calendar/v3"
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
	} else if argsWithoutProg[0] == "-c" {
		duration := 0
		recurrence := ""
		startTime := &calendar.EventDateTime{
			DateTime: "",
			TimeZone: "",
		}

		endTime := &calendar.EventDateTime{
			DateTime: "",
			TimeZone: "",
		}

		events.CreateEvent(argsWithoutProg[1], argsWithoutProg[2], startTime, endTime, duration, recurrence)
	}
}
