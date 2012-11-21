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

type Gender int

const (
	_ = 0
	Male
	Female
)

type NameGenerator struct {
	LastNames   []string
	FFirstNames []string
	MFirstNames []string
}

func NewNameGenerator() *NameGenerator {
	n := &NameGenerator{}
	n.Init()
	return n
}

// Load data for name generator from files
func (self *NameGenerator) Init() {
	var e error
	self.LastNames, e = loadArrayFromFile("data/names/dist.all.last.txt.json")
	if e != nil {
		panic(e.Error())
	}
	self.FFirstNames, e = loadArrayFromFile("data/names/dist.female.first.txt.json")
	if e != nil {
		panic(e.Error())
	}
	self.MFirstNames, e = loadArrayFromFile("data/names/dist.male.first.txt.json")
	if e != nil {
		panic(e.Error())
	}
}

// Generate a randomized name.
func (self *NameGenerator) Generate(s Gender) (first string, last string) {
	last = self.LastNames[rand.Intn(len(self.LastNames))]
	if s == Male {
		first = self.MFirstNames[rand.Intn(len(self.MFirstNames))]
	} else {
		first = self.FFirstNames[rand.Intn(len(self.FFirstNames))]
	}
	return
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
