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

	for _, char := range strings.Split(line, "") {
		digit, err := strconv.Atoi(char)

		// Skip if character is not a digit
		if err != nil {
			continue
		}

		digits = append(digits, digit)
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
