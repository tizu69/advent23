package main

import (
	"fmt"
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
)

var CLI struct {
	Day int `arg:""`
}

type days struct{}

func main() {
	_ = kong.Parse(&CLI)

	log.Info("Looking for challenge day", "day", CLI.Day)

	function := reflect.ValueOf(&days{}).MethodByName(fmt.Sprintf("D%dMain", CLI.Day))
	if !function.IsValid() {
		log.Fatal("Couldn't find challenge for that day")
	}

	function.Call(nil)
}
