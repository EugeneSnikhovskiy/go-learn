package main

import (
	"fmt"
	"errors"
)

func main() {
	fmt.Println(calendarPrediction("mon"))
	fmt.Println(calendarPrediction("tue"))
	fmt.Println(calendarPrediction("wrong"))
}

func calendarPrediction(day string) (string, error) {
	switch day {
	case "mon":
		return "Good day", nil
	case "tue":
		return "Bad day", nil
	default:
		return "other", errors.New("invalid day of week")
	}
}
