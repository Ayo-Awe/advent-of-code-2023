package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Algorithm:
	// read input from file line by line
	// identify the numbers on each line
	// retrieve first and last number on the line
	// derive calibration value
	// sum calibration values for all lines

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go input_file.txt")
	}

	filePath := os.Args[1]

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fileContent := string(bytes)
	lines := strings.Split(fileContent, "\n")
	sum := 0

	for _, line := range lines {
		digits := getLineDigits(line)
		calibrationValue := calcCalibrationValue(digits)
		sum += calibrationValue
	}

	fmt.Printf("Sum of calibration values = %v\n", sum)

}

func getLineDigits(line string) []int {
	var digits []int

	for idx := range strings.Split(line, "") {

		// check if numeric digit exists at index
		digit, found := checkNumericDigit(line, idx)
		if found {
			digits = append(digits, digit)
			continue
		}

		// else check if word digit exists at index
		digit, found = checkWordDigit(line, idx)
		if found {
			digits = append(digits, digit)
			continue
		}

	}

	return digits
}

func calcCalibrationValue(lineDigits []int) int {

	if len(lineDigits) == 0 {
		return 0
	}

	firstDigit := lineDigits[0]
	lastDigit := lineDigits[len(lineDigits)-1]

	// Calibration value is the combination of first and last digits on a line e.g 1, 2 = 12
	calibrationValue := firstDigit*10 + lastDigit

	return calibrationValue
}

func checkNumericDigit(str string, idx int) (digit int, found bool) {
	indexOutOfRange := int(idx) > (len(str) - 1)
	if indexOutOfRange {
		return 0, false
	}

	digit, err := strconv.Atoi(string(str[idx]))
	if err != nil {
		return 0, false
	}

	return digit, true
}

func checkWordDigit(str string, idx int) (digit int, found bool) {
	wordDigits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	indexOutOfRange := idx > (len(str) - 1)
	if indexOutOfRange {
		return 0, false
	}

	stringOffset := str[idx:]

	for key, value := range wordDigits {
		// Check if word begins with a word digit
		if strings.HasPrefix(stringOffset, key) {
			return value, true
		}
	}

	return 0, false
}
