package events

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"potential-doodle/utilities"

	"github.com/jedib0t/go-pretty/table"
	"google.golang.org/api/calendar/v3"
)

// DisplayEvents display events displays upcoming 10 events
func DisplayEvents() {
	t := time.Now().Format(time.RFC3339)
	srv, err := utilities.GetCalendarService()
	if err != nil {
		log.Fatal("error while getting calendar service")
		return
	}
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

// WriteToTable writes output to console ina pretty table format
func WriteToTable(events *calendar.Events) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle("UPCOMING EVENTS")
	t.AppendHeader(table.Row{"#", "EVENT NAME", "DATE", "TIME"})
	for i, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		dateTime := strings.Split(date, "T")
		t.AppendRow(table.Row{i + 1, item.Summary, dateTime[0], dateTime[1]})
	}
	t.SetStyle(table.StyleColoredBright)
	t.Render()
}
