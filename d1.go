package main

import (
	"strconv"

	"github.com/charmbracelet/log"
)

func (d *days) D1Main() {
	// Fetch data
	input := GetInput()
	input = append(input, '\n')

	results, thisLine := []int{}, []byte{}
	for i, v := range input {
		log.Debug("New run", "i", i, "v", string(v))

		if v == '\n' {
			toAdd, err := strconv.Atoi(string(thisLine[0]) + string(thisLine[len(thisLine)-1]))
			if err != nil {
				log.Error("Couldn't convert to int", "err", err)
			}

			results = append(results, toAdd)
			thisLine = []byte{}
			log.Debug("Appended to results", "now", results, "added", toAdd)
			continue
		}

		// If not a number, skip
		if !('0' <= v && v <= '9') {
			continue
		}

		thisLine = append(thisLine, v)
		log.Debug("Appended to the current line", "now", thisLine, "added", string(v))
	}

	log.Info("Got intermediate results", "results", results)

	returnVal := 0
	for _, v := range results {
		returnVal += v
	}

	log.Info("Found an answer!", "answer", returnVal)
}
