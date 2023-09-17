package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now().Format("01-02-2006")
	// now := time.Now().Format(time.RFC3339)
	// createdDate:= time.Date()
	// fmt.Println(now)

	// Your date string
	dateString := "2023-09-12T08:17:53+06:00"

	// Define the layout that matches your date format
	layout := "2006-01-02T15:04:05-07:00"

	// Parse the date string into a time.Time object
	t, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}

	// Print the parsed time
	fmt.Println("Parsed Time:", t.Year())
}
