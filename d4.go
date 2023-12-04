package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func (d *days) D4Main() {
	if CLI.PartTwo {
		log.Fatal("Part 2 is not (yet) implemented!")
	}

	// Fetch data
	lineInput := strings.Split(string(GetInput()), "\n")
	input := [][][]int{}
	for i, v := range lineInput {
		log.Debug("Processing line", "i", i, "v", v)

		thatLine, state := strings.Split(v, " "), 0
		inputResult := [][]int{{}, {}}

		for i := 0; i < len(thatLine); i += 1 {
			val := thatLine[i]

			switch true {
			case strings.HasSuffix(val, ":"):
				state = 1
				continue
			case val == "|":
				state = 2
				continue
			case val == "":
				continue
			case val == "Card":
				continue
			}

			intVal, err := strconv.Atoi(val)
			if err != nil {
				invalid = true
				log.Error("Failed to convert to int, skipped", "err", err, "val", val)
				continue
			}

			switch state {
			case 1:
				inputResult[0] = append(inputResult[0], intVal)
			case 2:
				inputResult[1] = append(inputResult[1], intVal)
			}

			log.Debug("Appending to input", "inputResult", inputResult)
		}

		input = append(input, inputResult)
	}

	log.Debug("Input processed", "input", input)

	returnVal := 0
	for i, v := range input { // for each line
		res := 0
		log.Debug("Processing", "i", i, "v", v)
		for j, w := range v[1] { // for each chosen
			for k, x := range v[0] { // for each winning
				log.Debug("Checking winnings", "i", i, "j", j, "k", k, "w", w, "x", x, "v[0]", v[0], "v[1]", v[1])
				if x == w {
					log.Debug("... won!", "res", res)
					res = int(math.Max(float64(res*2), 1))
					break
				}
			}
		}
		returnVal += res
	}

	log.Info("Found an answer!", "answer", returnVal)
}
