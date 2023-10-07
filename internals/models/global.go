package models

import "time"

type CalScform []CALSCFORMElement

type CALSCFORMElement struct {
	Uid         string   `json:"uid"`
	Dtstamp     string   `json:"dtstamp"`
	Dtstart     string   `json:"dtstart"`
	Dtend       string   `json:"dtend"`
	Summary     string   `json:"summary"`
	Location    string   `json:"location"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Priority    int64    `json:"priority"`
	Class       string   `json:"class"`
	Sequence    int64    `json:"sequence"`
}

// Create a month struct to store the month data

type Month struct {
	MonthNumber int
	MonthName   string
	MonthYear   int
	Weeks       []Week
}

// Create a week struct to store the week data
type Week struct {
	WeekNumber int
	Days       []Day
}

// Create a day struct to store the day data
type Day struct {
	DayNumber int
	DayName   string
	DayDate   time.Time
	DayEvents []Event
}

// Create an event struct to store the event data
type Event struct {
	EventName  string
	EventStart time.Time
	EventEnd   time.Time
	EventLink  string
}
