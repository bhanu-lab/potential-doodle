package main

import (
	"fmt"
	"log"
	"potential-doodle/utilities"
	"time"

	"google.golang.org/api/calendar/v3"
)

func main() {
	cfg, err := utilities.GetConfig()

	if err != nil {
		log.Fatal(err)
		return
	}
	client := utilities.GetClient(cfg)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		WriteToTable(events)
	}
}
