package util

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInputLineByLine(filename string) ([]string, error) {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, err
		}

		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ReadInput(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

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

func MustToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		log.Panicf("Unable to parse to int: %v", str)
	}

	return value
}
