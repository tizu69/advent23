package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

// what even is this.. forgive me (again) 🙏

func (d *days) D4Main() (int, string) {
	if CLI.PartTwo {
		log.Fatal("Part 2 is not (yet) implemented!")
	}
	returnVal, returnSum := 0, ""

	// Fetch data
	lineInput := strings.Split(string(GetInput()), "\n")
	input := [][][]int{}
	for i := 0; i < len(lineInput); i++ {
		v := lineInput[i]
		log.Debug("Processing line", "i", i, "v", v)

		thatLine, state := strings.Split(v, " "), 0
		inputResult := [][]int{{}, {}}

		for j := 0; j < len(thatLine); j += 1 {
			val := thatLine[j]

			switch true {
			case strings.HasSuffix(val, ":"):
				inputResult[0] = append(inputResult[0], i)
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

	for i := 0; i < len(input); i++ { // for each line
		v := input[i]

		res /*, ind*/ := 0 /*, 0*/
		log.Debug("Processing", "i", i, "v", v)
		for j, w := range v[1] { // for each chosen
			for k, x := range v[0][1:] { // for each winning
				log.Debug("Checking winnings", "i", i, "j", j, "k", k, "w", w, "x", x, "v[0]", v[0], "v[1]", v[1])
				if x == w {
					log.Debug("... won!", "res", res)
					returnSum += fmt.Sprintf("Winner was found      -> {L %d} %d\n", i+1, x)
					res = int(math.Max(float64(res*2), 1))
					/* if CLI.PartTwo {
						ind++
						input = append(input, input[v[0][0]+ind])
						returnSum += fmt.Sprintf("P2 card was added     -> {L %d} %d\n", i+1, ind)
					} */
					break
				}
			}
		}
		returnVal += res
	}

	if CLI.PartTwo {
		returnVal = len(input)
	}
	return returnVal, returnSum
}
