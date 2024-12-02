package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ScanFileToNumbers(filepath string) []int {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Declare a slice to hold the parsed integers
	numbers := []int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get the line from the scanner
		line := scanner.Text()

		// Split the line by spaces (you can modify this depending on your file format)
		parts := strings.Fields(line)

		// Convert each part to an integer and append it to the slice
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Println("Error converting string to int:", err)
				continue
			}
			numbers = append(numbers, num)
		}
	}
	return numbers

}

func SumInts(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
