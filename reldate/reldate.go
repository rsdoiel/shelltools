//
// Package reldate generates a date in YYYY-MM-DD format based on a
// relative time description (e.g. -1 week, +3 years)
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
// copyright (c) 2014 all rights reserved.
// Released under the Simplified BSD License
// See: http://opensource.org/licenses/bsd-license.php
//
package reldate

import (
	"errors"
	"strings"
	"time"
)

const (
	// The standard format for dates I find convient
	YYYYMMDD = "2006-01-02"
	// Version of this package
	Version = "v0.0.2"
)

// finds the end of the month value (e.g. 28, 29, 30, 31)
func EndOfMonth(t1 time.Time) string {
	location := t1.Location()
	year := t1.Year()
	month := t1.Month()
	if month == 12 {
		year++
	}
	month++
	t2 := time.Date(year, month, 1, 0, 0, 0, 0, location)
	return t2.Add(-time.Hour).Format(YYYYMMDD)
}

// computes the offset of a weekday time for a given weekday
func weekdayOffset(weekday time.Weekday) int {
	switch {
	case weekday == time.Sunday:
		return 0
	case weekday == time.Monday:
		return 1
	case weekday == time.Tuesday:
		return 2
	case weekday == time.Wednesday:
		return 3
	case weekday == time.Thursday:
		return 4
	case weekday == time.Friday:
		return 5
	case weekday == time.Saturday:
		return 6
	}
	return 0
}

// relativeWeekday converts the weekday name into an offset time and error
func relativeWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	// Normalize to Sunday then add weekday constant
	switch {
	case t.Weekday() == time.Sunday:
		return t.AddDate(0, 0, weekdayOffset(weekday)), nil
	case t.Weekday() == time.Monday:
		return t.AddDate(0, 0, (-1 + weekdayOffset(weekday))), nil
	case t.Weekday() == time.Tuesday:
		return t.AddDate(0, 0, (-2 + weekdayOffset(weekday))), nil
	case t.Weekday() == time.Wednesday:
		return t.AddDate(0, 0, (-3 + weekdayOffset(weekday))), nil
	case t.Weekday() == time.Thursday:
		return t.AddDate(0, 0, (-4 + weekdayOffset(weekday))), nil
	case t.Weekday() == time.Friday:
		return t.AddDate(0, 0, (-5 + weekdayOffset(weekday))), nil
	case t.Weekday() == time.Saturday:
		return t.AddDate(0, 0, (-6 + weekdayOffset(weekday))), nil
	}
	return t, errors.New("Expecting Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, or Saturday.")
}

// RelativeTime takes a time, an integer ammount (positive or negative)
// and a unit value (day of week, days, weeks, month, years) and
// computes the relative time in days from time returning a new
// time and error.
func RelativeTime(t time.Time, i int, u string) (time.Time, error) {
	switch {
	case strings.HasPrefix(u, "sun"):
		return relativeWeekday(t, time.Sunday)
	case strings.HasPrefix(u, "mon"):
		return relativeWeekday(t, time.Monday)
	case strings.HasPrefix(u, "tue"):
		return relativeWeekday(t, time.Tuesday)
	case strings.HasPrefix(u, "wed"):
		return relativeWeekday(t, time.Wednesday)
	case strings.HasPrefix(u, "thu"):
		return relativeWeekday(t, time.Thursday)
	case strings.HasPrefix(u, "fri"):
		return relativeWeekday(t, time.Friday)
	case strings.HasPrefix(u, "sat"):
		return relativeWeekday(t, time.Saturday)
	case strings.HasPrefix(u, "day"):
		return t.AddDate(0, 0, i), nil
	case strings.HasPrefix(u, "week"):
		return t.AddDate(0, 0, 7*i), nil
	case strings.HasPrefix(u, "month"):
		return t.AddDate(0, i, 0), nil
	case strings.HasPrefix(u, "year"):
		return t.AddDate(i, 0, 0), nil
	}
	return t, errors.New("Time unit must be day(s), week(s), month(s) or year(s) or weekday name.")
}
