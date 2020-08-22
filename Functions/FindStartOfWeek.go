package Functions

import "time"

func FindStartOfWeek(eventTime time.Time) string {
	var dayOfWeek = int(eventTime.Weekday())

	var startOfWeekDate time.Time

	if dayOfWeek != 0 {
		startOfWeekDate = eventTime.AddDate(0, 0, -(dayOfWeek - 1))
	} else {
		startOfWeekDate = eventTime.AddDate(0, 0, -6)
	}

	var extractedStartOfWeek = startOfWeekDate.Format("2006-01-02")

	return extractedStartOfWeek

}
