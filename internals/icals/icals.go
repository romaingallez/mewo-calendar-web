package icals

import (
	"calendar/internals/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

func MergeEventsByDay(events []models.CALSCFORMElement) []models.CALSCFORMElement {
	// Group by day and summary
	groupedEvents := make(map[string]map[string][]models.CALSCFORMElement)

	for _, event := range events {
		date, err := time.Parse("20060102T150405Z", event.Dtstart)
		if err != nil {
			log.Println(err)
			continue
		}

		dayKey := date.Format("2006-01-02")
		summaryKey := event.Summary

		if _, ok := groupedEvents[dayKey]; !ok {
			groupedEvents[dayKey] = make(map[string][]models.CALSCFORMElement)
		}

		groupedEvents[dayKey][summaryKey] = append(groupedEvents[dayKey][summaryKey], event)
	}

	// Merge events for each day and summary
	var mergedEvents []models.CALSCFORMElement
	for _, dayEventsMap := range groupedEvents {
		for _, summaryEvents := range dayEventsMap {
			if len(summaryEvents) == 0 {
				continue
			}

			// Sort events by start time within the day
			sort.Slice(summaryEvents, func(i, j int) bool {
				return summaryEvents[i].Dtstart < summaryEvents[j].Dtstart
			})

			mergedEvent := summaryEvents[0]
			mergedEvent.Dtend = summaryEvents[len(summaryEvents)-1].Dtend
			mergedEvents = append(mergedEvents, mergedEvent)
		}
	}

	return mergedEvents
}

func ParseTime(timeString string) time.Time {
	layout := "20060102T150405Z"
	t, err := time.Parse(layout, timeString)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func GetIcal(formation string) ([]models.CALSCFORMElement, error) {

	var url string

	proxyURL := os.Getenv("PROXY_URL")
	// log.Println(proxyURL)
	if len(proxyURL) == 0 {
		return nil, errors.New("PROXY_URL not set")
	}

	url = proxyURL + "/ical?formation=" + formation

	// log.Panicln(url)
	log.Println(url)
	// fmt.Scanln()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("User-Agent", "romain-bot")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// log.Println(string(body))
	// marshal body to []CALSCFORMElement
	var CalendarEvents []models.CALSCFORMElement

	err = json.Unmarshal(body, &CalendarEvents)
	if err != nil {

		log.Println(err)
		return nil, err
	}

	log.Println(len(CalendarEvents))

	// loop over ical events if event is the same date as the previous one, add it to the same event
	// disabled for now as there is event on the same day but different on the morning and afternoon
	CalendarEvents = MergeEventsByDay(CalendarEvents)

	// log.Println(len(CalendarEvents))

	return CalendarEvents, nil

}
