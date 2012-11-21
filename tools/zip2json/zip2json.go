package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Zip struct {
	ZipCode     string  `json:"zip"`
	StateAbbrev string  `json:"state"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lon"`
	City        string  `json:"city"`
	State       string  `json:"state_long"`
}

func main() {
	flag.Parse()
	args := flag.Args()
	for _, el := range args {
		fmt.Println("Processing " + el)
		file, err := os.Open(el)
		if err != nil {
			fmt.Println("Unable to open " + el)
			continue
		}
		items := []Zip{}
		reader := bufio.NewReader(file)

		// Read and ignore header line
		reader.ReadString('\n')

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				//fmt.Println("EOF")
				break
			}
			parts := strings.Split(strings.TrimSpace(line), "\"")
			lat, err := strconv.ParseFloat(strings.TrimSpace(parts[5]), 64)
			if err != nil {
				panic(err.Error())
			}
			lon, err := strconv.ParseFloat(strings.TrimSpace(parts[7]), 64)
			if err != nil {
				panic(err.Error())
			}
			z := Zip{
				ZipCode:     parts[1],
				StateAbbrev: parts[3],
				Latitude:    lat,
				Longitude:   lon,
				City:        parts[9],
				State:       parts[11],
			}
			items = append(items, z)
		}
		fmt.Printf("%s: Found %d items\n", el, len(items))

		j, err := json.Marshal(items)

		out, err := os.Create(el + ".json")
		if err != nil {
			fmt.Println("Unable to open " + el + ".json")
			file.Close()
			continue
		}
		out.Write(j)
		out.Close()

		file.Close()
	}
}
