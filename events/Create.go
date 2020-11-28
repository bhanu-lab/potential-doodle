package events

import (
	"fmt"
	"log"
	"potential-doodle/utilities"

	"google.golang.org/api/calendar/v3"
)

/*CreateEvent creates event using summary description start time and duration of event
example struct

event := &calendar.Event{
	Summary:     "Google I/O 2015",
	Location:    "800 Howard St., San Francisco, CA 94103",
	Description: "A chance to hear more about Google's developer products.",
	Start: &calendar.EventDateTime{
		DateTime: "2015-05-28T09:00:00-07:00",
		TimeZone: "America/Los_Angeles",
	},
	End: &calendar.EventDateTime{
		DateTime: "2015-05-28T17:00:00-07:00",
		TimeZone: "America/Los_Angeles",
	},
	Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
	Attendees: []*calendar.EventAttendee{
		&calendar.EventAttendee{Email: "lpage@example.com"},
		&calendar.EventAttendee{Email: "sbrin@example.com"},
	},
}
*/
func CreateEvent(summary string, description string, start *calendar.EventDateTime, end *calendar.EventDateTime, duration int, recurrence string) error {
	event := &calendar.Event{
		Summary:     summary,
		Description: description,
		Start:       start,
		End:         end,
		Recurrence:  []string{recurrence},
	}

	calendarID := "primary"
	srv, err := utilities.GetCalendarService()
	if err != nil {
		log.Fatal("error whille getting calendar service")
		return err
	}
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %+v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	return nil
}
