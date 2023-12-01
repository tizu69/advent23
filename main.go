package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

var CLI struct {
	Day       int  `arg:""`
	NoVerbose bool `help:"DISABLE verbose logging"`
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
}

func GetInput() []byte {
	log.Info("Create a file (if neccessary) called 'input.txt', edit it to your liking, then press enter to continue")
	fmt.Scanln()

	// Read data from file
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("An error occurred while reading ", "err", err)
	}

	return input
}
