package main

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func (d *days) D1Main() {
	// Fetch data
	input := GetInput()
	// input = append(input, '\n')

	//#region Task 2
	if CLI.PartTwo {
		wordedNumbers := map[string]byte{
			"zero":  '0',
			"one":   '1',
			"two":   '2',
			"three": '3',
			"four":  '4',
			"five":  '5',
			"six":   '6',
			"seven": '7',
			"eight": '8',
			"nine":  '9',
		}
		for i := 0; i < len(input); i++ {
			for k, v := range wordedNumbers {
				if strings.HasPrefix(string(input[i:]), k) {
					input[i] = v
					break
				}
			}
		}
		log.Debug("Converted for p2", "input", string(input))
	}
	//#endregion

	results, thisLine := []int{}, []byte{}
	for i, v := range input {
		log.Debug("New run", "i", i, "v", string(v))

		if v == '\n' {
			toAdd, err := strconv.Atoi(string(thisLine[0]) + string(thisLine[len(thisLine)-1]))
			if err != nil {
				invalid = true
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
