package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
)

func InSlice[T comparable](a T, list []T) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func GetInput() []byte {
	log.Info("Create a file (if neccessary) called '" + CLI.FileName + ".txt', edit it to your liking, then press enter to continue")
	fmt.Scanln()

	// Read data from file
	input, err := os.ReadFile(CLI.FileName + ".txt")
	if err != nil {
		log.Fatal("An error occurred while reading ", "err", err)
	}

	return input
}

func MultiplyDigits(input string) int {
	result := 1
	for _, char := range input {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("Error converting character to integer:", err)
			return 0
		}
		result *= digit
	}
	return result
}
