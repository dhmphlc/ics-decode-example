package main

import (
	"fmt"
	"time"

	ics "github.com/arran4/golang-ical"
)

func main() {
	// your ics calendar url
	url := "https://www.example.com/ics/export"

	GetCalendarEventsFromURL(url)
}

type Event struct {
	UID         string
	Timestamp   time.Time
	Start       time.Time
	End         time.Time
	Summary     string
	Description string
}

func GetCalendarEventsFromURL(url string) {
	cal, err := ics.ParseCalendarFromUrl(url)
	if err != nil {
		fmt.Printf("failed to parse calendar: %v", err)
	}

	events := []Event{}
	for _, event := range cal.Events() {
		resEvent := Event{}
		for _, prop := range event.Properties {
			switch prop.IANAToken {
			case "DTSTART":
				resEvent.Start, err = event.GetStartAt()
				if err != nil {
					fmt.Println("Error getting start time:", err)
				}
			case "DTEND":
				resEvent.End, err = event.GetEndAt()
				if err != nil {
					fmt.Println("Error getting end time:", err)
				}
			case "DTSTAMP":
				resEvent.Timestamp, err = event.GetDtStampTime()
				if err != nil {
					fmt.Println("Error getting timestamp:", err)
				}
			case "UID":
				resEvent.UID = prop.Value
			case "SUMMARY":
				resEvent.Summary = prop.Value
			case "DESCRIPTION":
				resEvent.Description = prop.Value
			}
		}
		events = append(events, resEvent)
	}

	for _, event := range events {
		fmt.Printf("UID: %s\n", event.UID)
		fmt.Printf("Summary: %s\n", event.Summary)
		fmt.Printf("Description: %s\n", event.Description)
		fmt.Printf("Start At: %s\n", event.Start)
		fmt.Printf("End At: %s\n", event.End)
		fmt.Printf("Timestamp: %s\n", event.Timestamp)
		fmt.Println()
	}
}
