package handler

import (
	"calendar/internals/icals"
	"calendar/internals/models"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	_ "time/tzdata"

	"github.com/gofiber/fiber/v2"
)

// invertFormation map
var invertFormation = map[string]string{
	"dev":   "cyber",
	"cyber": "dev",
}

func GetHandleMonth(c *fiber.Ctx) (err error) {

	ParisLocation, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		return c.SendString(err.Error())
	}

	formation := c.Query("formation", "dev")

	currentDay := time.Now().In(ParisLocation)

	currentYear := time.Now().In(ParisLocation).Year()
	currentMonth := time.Now().In(ParisLocation).Month()

	log.Println(currentMonth)

	year := c.Query("year")

	month := c.Query("month")

	monthInt := int(currentMonth)
	var formattedMonth string
	if math.Abs(float64(monthInt)) < 10 {
		// append a 0 before the int
		formattedMonth = fmt.Sprintf("0%d", monthInt)
	} else {
		formattedMonth = fmt.Sprintf("%d", monthInt)
	}
	log.Println(formattedMonth)
	// if the query is empty, redirect to the current month and year
	if year == "" || month == "" {
		return c.Redirect(fmt.Sprintf("/month?formation=%s&year=%d&month=%s", formation, currentYear, formattedMonth))
	}

	// log.Println(formation, year, month)

	// Create a month time.Time from year and month number
	MonthTime, err := time.Parse("2006-01", fmt.Sprintf("%s-%s", year, month))
	if err != nil {
		return c.SendString(err.Error())
	}

	log.Println(MonthTime, int(MonthTime.Month()))
	Weeks := GenerateWeek(MonthTime.Year(), int(MonthTime.Month()), formation)

	// return c.SendString("ok")
	MonthM := models.Month{
		MonthNumber: int(MonthTime.Month()),
		MonthName:   models.FrenchMonthMap[MonthTime.Month().String()],
		MonthYear:   MonthTime.Year(),
		Weeks:       Weeks,
		// generate a week struct for each week in the month
	}

	// Convert a letter from lower case to upper case
	formationFirstLetter := strings.ToUpper(string(formation[0]))
	// get the rest of the string
	formationRest := formation[1:]
	// add the first letter to the rest of the string
	formationFirstLetterUpper := formationFirstLetter + formationRest

	RenderMap := fiber.Map{
		"InvertFormation": invertFormation[formation],
		"Month":           MonthM,
		"CurrentDay":      currentDay,
		"Year":            year,
		"Formation":       formationFirstLetterUpper,
	}

	//

	return c.Render("month", RenderMap, "layouts/main")

}

// Function to generate a []models.Week struct for each week in the month
func GenerateWeek(year int, month int, formation string) (Weeks []models.Week) {

	// get the first day of the month
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	// get the last day of the month
	lastDay := firstDay.AddDate(0, 1, -1)
	// get the first day of the first week of the month
	firstWeekDay := firstDay.AddDate(0, 0, -int(firstDay.Weekday()))
	// get the last day of the last week of the month
	lastWeekDay := lastDay.AddDate(0, 0, 6-int(lastDay.Weekday()))

	// get the number of weeks in the month
	weeks := int(lastWeekDay.Sub(firstWeekDay).Hours()/24/7) + 1

	calendarEvents, err := icals.GetIcal(formation)
	if err != nil {
		log.Println(err)
	}

	// log.Printf("number of weeks: %d", weeks)

	// generate a week struct for each week in the month
	for i := 0; i < weeks; i++ {
		// Generate all the days in the week
		// get the first day of the week
		firstDayOfWeek := firstWeekDay.AddDate(0, 0, i*7)
		// get the last day of the week
		lastDayOfWeek := firstDayOfWeek.AddDate(0, 0, 6)
		// get the number of days in the week for a (monday to friday)
		days := int(lastDayOfWeek.Sub(firstDayOfWeek).Hours()/24) + 1

		// log.Printf("first day of the week: %s\n last day of the week %s\n days: %d", firstDayOfWeek, lastDayOfWeek, days)

		// days := 5
		// var Days []models.Day
		Days := make([]models.Day, 0, days)
		// generate a day struct for each day in the week
		for j := 0; j < days; j++ {
			var EmptyDay = false
			// get the day number
			currentDay := firstDayOfWeek.AddDate(0, 0, j)
			if currentDay.Month() != firstDay.Month() {
				// If the current day is not in the same month as the first day, set to an empty time.Time
				// log.Println("not in the same month")
				// EmptyDay = true

			}
			dayNumber := currentDay.Day()
			// get the day name
			dayName := currentDay.Weekday().String()

			// If the day is a Saturday or Sunday or an empty time, skip
			if currentDay.IsZero() || currentDay.Weekday() == time.Saturday || currentDay.Weekday() == time.Sunday {
				continue
			}

			var dayEvents []models.Event

			ParisLocation, err := time.LoadLocation("Europe/Paris")
			if err != nil {
				log.Println(err)
			}

			// loop over all the events in the calendar test if it is in the day
			for _, event := range calendarEvents {

				DtStart, err := time.Parse("20060102T150405Z", event.Dtstart)
				if err != nil {
					log.Println(err)
				}

				DtEnd, err := time.Parse("20060102T150405Z", event.Dtend)
				if err != nil {
					log.Println(err)
				}

				// print dtStart with format day month year
				// fmt.Println(DtStart.Format("02/01/2006 15:04:05"))

				// log.Println(DtStart)
				// check if the event date is the same as the day date
				// test if the event day is the same as the day day and if the event month is the same as the day month and if the event year is the same as the day year
				if DtStart.Day() == currentDay.Day() && DtStart.Month() == currentDay.Month() && DtStart.Year() == currentDay.Year() {
					// log.Println(DtStart.Format("15:04"), DtEnd.Format("15:04"))

					// convert DtStart to a time in paris timezone

					DtStartParis := DtStart.In(ParisLocation)

					DtEndParis := DtEnd.In(ParisLocation)

					// log.Println(event.Summary, DtStart)
					// if the event date is the same as the day date
					// add the event to the day events
					dayEvents = append(dayEvents, models.Event{
						EventName:  event.Summary,
						EventStart: DtStartParis,
						EventEnd:   DtEndParis,
						EventLink:  event.Description,
					})
				}
			}

			fullDayName := fmt.Sprintf("%s %d %s %d", models.FrenchDayMap[dayName], dayNumber, models.FrenchMonthMap[time.Month(currentDay.Month()).String()], year)
			// log.Println(fullDayName)

			// generate a day struct for each day in the week
			Days = append(Days, models.Day{
				DayNumber: dayNumber,
				DayName:   fullDayName,
				DayDate:   currentDay,
				DayEvents: dayEvents,
				Empty:     EmptyDay,
			})
		}

		Weeks = append(Weeks, models.Week{
			WeekNumber: i + 1,
			Days:       Days,
		})
	}

	// print the last week day with each day

	return Weeks
}
