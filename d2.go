package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func (d *days) D2Main() {
	// Fetch data
	lineInput := strings.Split(string(GetInput()), "\n")
	input := [][]string{}
	for i, v := range lineInput {
		log.Debug("Processing line", "i", i, "v", v)
		re := regexp.MustCompile(`[\W]+`)
		input = append(input, strings.Split(re.ReplaceAllString(v, " "), " "))
	}

	p2r, p2g, p2b := make([]int, len(input)), make([]int, len(input)), make([]int, len(input))

	const (
		containsRed   = 12
		containsGreen = 13
		containsBlue  = 14
	)

	impossible := make([]bool, len(input))
	for i := 0; i < len(input); i += 1 {
		for j := 0; j < len(input[i]); j += 2 {
			if j+1 >= len(input[i]) {
				invalid = true
				log.Warn("Input does not match format", "i", i, "j", j, "input[i]", input[i])
				break
			}

			v := input[i][j]
			vA := input[i][j+1]
			log.Debug("New run", "i", i, "j", j, "v", v, "vA", vA, "input[i]", input[i])

			if v == "Game" {
				continue
			}

			count, err := strconv.Atoi(v)
			if err != nil {
				invalid = true
				log.Error("Failed to convert to int, skipped", "err", err, "v", v)
				continue
			}

			//#region Task 2
			if CLI.PartTwo {
				switch vA {
				case "red":
					p2r[i] = max(p2r[i], count)
				case "green":
					p2g[i] = max(p2g[i], count)
				case "blue":
					p2b[i] = max(p2b[i], count)
				default:
					invalid = true
					log.Error("Invalid color for part 2", "color", vA)
				}
			}
			//#endregion

			switch vA {
			case "red":
				log.Debug("Appended to results", "before", impossible, "added", count > containsRed, "color", vA)
				impossible[i] = impossible[i] || count > containsRed
			case "green":
				log.Debug("Appended to results", "before", impossible, "added", count > containsGreen, "color", vA)
				impossible[i] = impossible[i] || count > containsGreen
			case "blue":
				log.Debug("Appended to results", "before", impossible, "added", count > containsBlue, "color", vA)
				impossible[i] = impossible[i] || count > containsBlue
			default:
				invalid = true
				log.Error("Invalid color", "color", vA)
			}
		}
	}

	returnVal := 0
	for i, v := range impossible {
		if !v {
			log.Debug("Processing results", "i", i, "v", v, "returnVal", returnVal)
			returnVal += i + 1
		}
	}
	log.Info("Found an answer!", "answer", returnVal)

	//#region Task 2
	if CLI.PartTwo {
		log.Debug("Processing p2 results", "p2r", p2r, "p2g", p2g, "p2b", p2b)
		p2result := 0

		for i := range p2r {
			p2result += p2r[i] * p2g[i] * p2b[i]
		}

		log.Info("Found an answer for p2!", "answer", p2result)
	}
	//#endregion
}
