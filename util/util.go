package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() ([]string, error) {
	if len(os.Args) < 2 {
		return nil, errors.New("usage: go run main.go input_file.txt")
	}

	filePath := os.Args[1]

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	fileContent := string(bytes)
	fileContent = strings.TrimSuffix(fileContent, "\n") // remove trailing new line

	lines := strings.Split(fileContent, "\n")

	return lines, nil
}

func IsDigit(str string) bool {
	if len(str) > 1 {
		return false
	}

	_, err := strconv.Atoi(str)
	return err == nil
}

func ToIntArray(strArray []string) []int {
	var intArray []int

	for _, s := range strArray {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Unable to parse to int: %v", s)
		}

		intArray = append(intArray, int(value))
	}

	return intArray
}
