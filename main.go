package main

import (
	"fmt"
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

var CLI struct {
	Day          int    `arg:""`
	NoVerbose    bool   `help:"DISABLE verbose logging"`
	PartTwo      bool   `short:"2" help:"Use P2 logic"`
	FileName     string `help:"Name of input file" short:"f" default:"input"`
	ExpectedData string `help:"Expected output" short:"e"`
}
var invalid = false

type days struct{}

func main() {
	_ = kong.Parse(&CLI)

	if !CLI.NoVerbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("Looking for challenge day", "day", CLI.Day)

	function := reflect.ValueOf(&days{}).MethodByName(fmt.Sprintf("D%dMain", CLI.Day))
	if !function.IsValid() {
		log.Fatal("Couldn't find challenge for that day")
	}

	function.Call(nil)

	if invalid {
		log.Fatal("Error was encountered, output may be invalid!")
	}

	if CLI.ExpectedData != "" {
		log.Info("Expected output", "expected", CLI.ExpectedData)
	}
	if CLI.PartTwo {
		log.Info("Note: Part 2 results may differ from part 1!")
	}
}
