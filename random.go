package main

import (
	"math/rand"
	"time"
)

// Generate a random date, usable for date of birth generation.
func GenerateRandomDate() time.Time {
	day := int(rand.Int31n(27)) + 1
	month := (time.Month)(rand.Int31n(11) + 1)
	if month == time.February && day > 27 {
		day = 27
	}
	year := time.Now().Year() - int(rand.Int31n(90))
	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}
