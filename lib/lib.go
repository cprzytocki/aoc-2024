package aoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ScanFileToNumbers(filepath string, separator ...string) []int {
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

		var parts []string
		if len(separator) > 0 {
			parts = strings.Split(line, separator[0])
		} else {
			parts = strings.Fields(line)
		}

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

func ScanFileToRows(filepath string, separator ...string) [][]int {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Declare a slice to hold the parsed integers
	numbersArray := [][]int{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get the line from the scanner
		line := scanner.Text()
		var parts []string
		if len(separator) > 0 {
			parts = strings.Split(line, separator[0])
		} else {
			parts = strings.Fields(line)
		}

		numbersRow := []int{}

		// Convert each part to an integer and append it to the slice
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Println("Error converting string to int:", err)
				continue
			}
			numbersRow = append(numbersRow, num)
		}
		numbersArray = append(numbersArray, numbersRow)
	}
	return numbersArray
}

func ScanFileToMatrix(filepath string) [][]string {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	matrix := [][]string{}
	// Declare a slice to hold the parsed strings
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get the line from the scanner
		line := scanner.Text()
		stringArray := []string{}

		for _, char := range line {
			stringArray = append(stringArray, string(char))
		}
		matrix = append(matrix, stringArray)
	}
	return matrix
}

func SumInts(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func Abs(num int) int {
	return max(num, -num)
}

func RemoveElementByIndex(slice []int, index int) ([]int, error) {
	// Check if the index is valid
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index out of range")
	}

	// Create a new slice and append elements excluding the one at the specified index
	newSlice := append([]int(nil), slice[:index]...) // Copy elements before index
	newSlice = append(newSlice, slice[index+1:]...)  // Append elements after index

	return newSlice, nil
}

type Pair struct {
	Num1 int
	Num2 int
}

func MultiplyPair(pair Pair) int {
	return pair.Num1 * pair.Num2
}

func ArraysHaveSameElement(arr1 []int, arr2 []int) bool {
	for _, item := range arr1 {
		if slices.Contains(arr2, item) {
			return true
		}
	}
	return false
}

func ArrayContainsAllElements(arr1 []int, arr2 []int) bool {
	for _, item := range arr2 {
		if !slices.Contains(arr1, item) {
			return false
		}
	}
	return true
}
