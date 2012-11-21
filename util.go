package main

import (
	"encoding/json"
	"errors"
	"os"
)

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
