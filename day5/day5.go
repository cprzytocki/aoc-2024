package aoc

import (
	lib "aoc/2024/lib"
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Order struct {
	before int
	after  int
}

func scanManual(filepath string) []Order {
	orders := []Order{}
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Declare a slice to hold the parsed integers
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Get the line from the scanner
		line := scanner.Text()
		parts := strings.Split(line, "|")

		before, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Println("Error converting string to int:", err)
			continue
		}

		after, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Println("Error converting string to int:", err)
			continue
		}

		orders = append(orders, Order{before, after})
	}

	return orders
}

func printRows(rows interface{}) {
	switch rows := rows.(type) {
	case [][]int:
		for _, row := range rows {
			fmt.Println(row)
		}
	case map[int][]int:
		for key, value := range rows {
			fmt.Println(key, value)
		}
	case []Order:
		for _, order := range rows {
			fmt.Println(order)
		}
	}
}

func buildAdjacencyList(orders []Order) map[int][]int {
	adjacencyList := map[int][]int{}
	for _, order := range orders {
		adjacencyList[order.before] = append(adjacencyList[order.before], order.after)
	}

	return adjacencyList
}

func rowPasses(row []int, adjacencyList map[int][]int) bool {
	for i := len(row) - 1; i > 0; i-- {
		if lib.ArraysHaveSameElement(row[:i], adjacencyList[row[i]]) {
			return false
		}
	}
	return true
}

func getSums(rows [][]int) int {
	sum := 0
	for _, row := range rows {
		middleIndex := (len(row) - 1) / 2
		sum += row[middleIndex]
	}
	return sum
}

func Day5() int {
	rows := lib.ScanFileToRows("day5/input", ",")
	orders := scanManual("day5/input_manual")
	adjacencyList := buildAdjacencyList(orders)

	// printRows(rows)
	// printRows(orders)
	// printRows(adjacencyList)

	passedRows := [][]int{}
	for _, row := range rows {
		if rowPasses(row, adjacencyList) {
			passedRows = append(passedRows, row)
		}
	}

	fmt.Println("Passed rows:")
	printRows(passedRows)

	return getSums(passedRows)
}

//
//  PART 2
//

func trimDependencies(dependencies []int, arrayToSort []int) []int {
	newDependencies := []int{}
	for _, dependency := range dependencies {
		if slices.Contains(arrayToSort, dependency) {
			newDependencies = append(newDependencies, dependency)
		}
	}
	return newDependencies
}

func trimAdjacencyList(adjacencyList map[int][]int, arrayToSort []int) map[int][]int {
	newAdjacencyList := map[int][]int{}
	for key, value := range adjacencyList {
		if slices.Contains(arrayToSort, key) {
			newAdjacencyList[key] = trimDependencies(value, arrayToSort)
		}
	}
	return newAdjacencyList
}

func sortRow(arrayToSort []int, adjacencyList map[int][]int) []int {
	newArray := []int{}
	// Trim the adjacency list to only include the items in the array to sort
	trimmedAdjacencyList := trimAdjacencyList(adjacencyList, arrayToSort)

	for len(newArray) != len(arrayToSort) {
		for _, item := range arrayToSort {
			dependencies := trimmedAdjacencyList[item]
			if !slices.Contains(newArray, item) {
				if len(dependencies) == 0 {
					newArray = slices.Insert(newArray, 0, item)
				} else if lib.ArrayContainsAllElements(newArray, dependencies) {
					newArray = slices.Insert(newArray, 0, item)
				}
			}
		}
	}
	return newArray
}

func Day5Part2() int {
	rows := lib.ScanFileToRows("day5/input", ",")
	orders := scanManual("day5/input_manual")
	adjacencyList := buildAdjacencyList(orders)

	// printRows(rows)
	// printRows(orders)
	// printRows(adjacencyList)

	failedRows := [][]int{}
	for _, row := range rows {
		if !rowPasses(row, adjacencyList) {
			failedRows = append(failedRows, row)
		}
	}
	sortedFailedRows := [][]int{}
	for _, row := range failedRows {
		sortedFailedRows = append(sortedFailedRows, sortRow(row, adjacencyList))
	}

	return getSums(sortedFailedRows)
}
