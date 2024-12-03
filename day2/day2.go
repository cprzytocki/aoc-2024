package aoc

import (
	lib "aoc/2024/lib"
)

func isRowSafe_P1(row []int) bool {
	var slope bool

	for i, num := range row {
		// last element
		if i == len(row)-1 {
			break
		}
		diff := row[i+1] - num

		// calculate initial slope
		if i == 0 {
			slope = diff > 0
		} else {
			// check if slope changes
			if diff > 0 != slope {
				return false
			}
		}
		//  check if diff between 1-3
		if diff == 0 || lib.Abs(diff) > 3 {
			return false
		}
	}

	return true
}

func isRowSafe_P2(row []int) bool {
	if isRowSafe_P1(row) {
		return true
	}

	for i := range row {
		newRow, err := lib.RemoveElementByIndex(row, i)
		if err != nil {
			panic(err)
		}

		if isRowSafe_P1(newRow) {
			return true
		}
	}

	return false
}

func Day2() int {
	rows := lib.ScanFileToRows("./day2/input")

	safeCount := 0

	for _, row := range rows {
		if isRowSafe_P1(row) {
			safeCount++
		}
	}
	return safeCount

}

func Day2Part2() int {
	rows := lib.ScanFileToRows("./day2/input")

	safeCount := 0

	for _, row := range rows {
		if isRowSafe_P2(row) {
			safeCount++
		}
	}
	return safeCount

}
