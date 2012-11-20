package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

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
		items := []string{}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				//fmt.Println("EOF")
				break
			}
			items = append(items, strings.TrimSpace(line))
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
