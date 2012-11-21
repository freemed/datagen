package main

import (
	"math/rand"
)

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
