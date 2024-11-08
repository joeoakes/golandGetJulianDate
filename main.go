package main

import (
	"fmt"
	"time"
)

// calculates the Julian date for the current date and time
func getJulianDate() float64 {
	// Define the Julian date for the Unix epoch (1970-01-01 00:00:00 UTC)
	julianDateEpoch := 2440587.5

	// Get the current time and calculate the duration since the Unix epoch
	now := time.Now()
	durationSinceEpoch := now.Sub(time.Unix(0, 0))

	// Calculate the days since the Unix epoch
	daysSinceEpoch := durationSinceEpoch.Hours() / 24.0

	// Calculate and return the Julian date
	return julianDateEpoch + daysSinceEpoch
}

func main() {
	// Get and print the Julian date
	julianDate := getJulianDate()
	fmt.Printf("The current Julian date is: %f\n", julianDate)
}
