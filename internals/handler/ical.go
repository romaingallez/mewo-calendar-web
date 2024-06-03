package handler

import (
	"calendar/internals/icals"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"

	ical "github.com/arran4/golang-ical"
)

func GetHandleIcal(c *fiber.Ctx) error {

	formation := c.Query("formation", "dev")

	// ParisLocation, err := time.LoadLocation("Europe/Paris")
	// if err != nil {
	// 	return c.SendString(err.Error())
	// }

	cal_raw, err := icals.GetIcalRaw(formation)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	cal := ical.NewCalendar()

	for _, element := range cal_raw {
		event := cal.AddEvent(element.Uid)
		event.SetSummary(element.Summary)
		startEvent, err := time.Parse("20060102T150405Z", element.Dtstart)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		endEvent, err := time.Parse("20060102T150405Z", element.Dtend)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		dtTimeStamp, err := time.Parse("20060102T150405Z", element.Dtstamp)
		if err != nil {
			log.Println(err)
			dtTimeStamp = startEvent
		}
		event.SetDtStampTime(dtTimeStamp)
		event.SetStartAt(startEvent)
		event.SetEndAt(endEvent)
		event.SetLocation(element.Location)
		event.SetDescription(element.Description)
	}

	// Convert calendar to string
	icalString := cal.Serialize()

	// log.Println(icalString)

	// Set the Content-Type to text/calendar
	c.Set(fiber.HeaderContentType, "text/calendar")
	return c.SendString(icalString)

}
