package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
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

// Deserialize a JSON array from a file
func loadArrayFromFile(f string) ([]string, error) {
	fi, fierr := os.Stat(f)
	if fierr != nil {
		return nil, fierr
	}

	file, err := os.Open(f)
	if err != nil {
		return nil, fierr
	}

	data := make([]byte, fi.Size())
	count, err := file.Read(data)
	if count == 0 {
		return nil, errors.New("Unable to read from file")
	}
	if err != nil {
		file.Close()
		return nil, err
	}
	file.Close()

	var d []string
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
