package main

import (
	"fmt"
	"log"
	"os"
	"potential-doodle/events"
	"strconv"
	"strings"
	"time"

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
		fmt.Printf("\t-c\tCreates Event mention date and time in below format \n")
		fmt.Printf("\t\t23/09/2020-5:30 \n")
	} else if argsWithoutProg[0] == "-c" {
		var summary, description string
		dateTime := argsWithoutProg[1]
		if len(argsWithoutProg) == 3 || len(argsWithoutProg) > 3 {
			summary = argsWithoutProg[2]
		}
		if len(argsWithoutProg) == 4 || len(argsWithoutProg) > 4 {
			description = argsWithoutProg[3]
		}
		eventDateTime := strings.Split(dateTime, "-")
		eventDate, eventTime := eventDateTime[0], eventDateTime[1]
		date := strings.Split(eventDate, "/")
		tym := strings.Split(eventTime, ":")
		y, err := strconv.Atoi(date[0])
		m, err := strconv.Atoi(date[1])
		d, err := strconv.Atoi(date[2])
		h, err := strconv.Atoi(tym[0])
		min, err := strconv.Atoi(tym[1])
		//s, err := strconv.Atoi(tym[2])
		if err != nil {
			log.Fatal("error occured while trying to convert to year string to integer")
			return
		}
		loc, err := time.LoadLocation("GMT")
		if err != nil {
			log.Fatal("error while loading location")
			return
		}
		t1 := time.Date(y, time.Month(m), d, h, min, 0, 0, loc)
		t2 := time.Date(y, time.Month(m), d, h, min+30, 0, 0, loc)

		duration := 30
		recurrence := ""
		startTime := &calendar.EventDateTime{
			DateTime: t1.String(),
			TimeZone: loc.String(),
		}

		endTime := &calendar.EventDateTime{
			DateTime: t2.String(),
			TimeZone: loc.String(),
		}

		events.CreateEvent(summary, description, startTime, endTime, duration, recurrence)
	}
}
