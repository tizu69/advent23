package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/charmbracelet/log"
)

// what even is this.. forgive me 🙏

func (d *days) D3Main() {
	if CLI.PartTwo {
		log.Fatal("Part 2 is not (yet) implemented!")
	}

	// Fetch data
	lineInput := append(GetInput(), '\n')
	line := 0
	input := [][]byte{}
	for i, v := range lineInput {
		if v == '\n' {
			// HACK: This appends a . to the end of the line as my code didn't consider the last column and I'm lazy.
			input = append(input, append(lineInput[line:i], '.'))
			line = i + 1
		}
	}

	returnVal, gearRatio, results := 0, 0, []string{}
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			num, valid, gear := d3getNum(x, y, input, false)
			if valid {
				returnVal += num
				gearRatio += int(math.Max(0, float64(gear)))
				results = append(results, fmt.Sprintf("x=%d, y=%d, num=%d, valid=%t gearRatio=%d ||", x, y, num, valid, gear))
			}
		}
	}
	log.Info("Got intermediate result", "results", results)
	log.Info("Found an answer!", "answer", returnVal, "gearRatio", gearRatio)
}

var d3directions = map[string][]int{
	"top":          {0, -1},
	"bottom":       {0, 1},
	"left":         {-1, 0},
	"right":        {1, 0},
	"top-right":    {1, -1},
	"bottom-right": {1, 1},
	"bottom-left":  {-1, 1},
	"top-left":     {-1, -1},
}

// Returns the complete number at the given coordinates (-1) if not a number, and wether or not it is a valid part
// Return values: number, valid, gearRatio
func d3getNum(x int, y int, input [][]byte, nested bool) (int, bool, int) {
	v := input[y][x]
	log.Debug("Checking", "x", x, "y", y, "v", v, "input[y][x]", input[y][x], "input[y]", input[y], "d3directions[\"right\"]", d3directions["right"],
		"sting(input[y])", string(input[y]))

	if len(input) <= y || len(input[y]) <= x+d3directions["right"][0] {
		log.Debug("Not even considering, out of bounds", "x", x, "y", y, "string(v)", string(v))
		return -1, false, -1
	}
	if !('0' <= v && v <= '9') {
		log.Debug("Not a number", "x", x, "y", y, "string(v)", string(v))

		//#region Task 2
		/* if CLI.PartTwo && !nested && v == '*' {
			gearRatio, gears := 1, 0
			for key, direction := range d3directions {
				tY, tX := y+direction[1], y+direction[0]

				if tY < 0 || tX < 0 || tY >= len(input) || tX >= len(input[0]) {
					log.Debug("Out of bounds, not considering", "x", x, "y", y,
						"tY", tY, "tX", tX, "key", key, "len(input)", len(input), "len(input[0])", len(input[0]))
					continue
				}

				number, _, _ := d3getNum(tX, tY, input, true)
				if number == -1 {
					continue
				}

				gearRatio *= number
				gears++
				log.Debug("New gear data", "gears", gears, "gearRatio", gearRatio, "number", number, "key", key)
			}

			if gears == 2 {
				return -1, true, gearRatio
			} else if gears == 0 {
				return -1, false, -1

			}
		} */
		//#endregion
		return -1, false, -1
	}
	if val := input[y][x+d3directions["right"][0]]; '0' <= val && val <= '9' {
		log.Debug("Number, but handled elsewhere", "x", x, "y", y, "v", v)
		return -1, false, -1 // Return -1 so that the last digit actually counts as the number
	}

	num := ""
	validPart := false
	for i := x; true; i-- {
		if i < 0 || i >= len(input[0]) {
			log.Debug("Out of bounds", "x", x, "y", y, "i", i)
			break
		}

		val := input[y][i]
		if !('0' <= val && val <= '9') {
			break
		}
		num = string(val) + num

		for key, direction := range d3directions {
			tY, tX := y+direction[1], i+direction[0]
			if tY < 0 || tX < 0 || tY >= len(input) || tX >= len(input[0]) {
				log.Debug("Out of bounds, not considering", "x", x, "y", y, "i", i,
					"tY", tY, "tX", tX, "key", key, "len(input)", len(input), "len(input[0])", len(input[0]))
				continue
			}
			thatOne := input[tY][tX]
			log.Debug("Considering special char", "string(thatOne)", string(thatOne), "key", key, "tY", tY, "tX", tX)
			if !InSlice[byte](thatOne, []byte{'.', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}) {
				validPart = true
				log.Debug("... found!")

			}
		}
	}

	convertedNum, err := strconv.Atoi(num)
	if err != nil {
		invalid = true
		log.Error("Couldn't convert to int", "err", err, "num", num)
		return -1, false, -1
	}

	log.Debug("Returning num", "num", convertedNum, "validPart", validPart)
	return convertedNum, validPart, -1
}
