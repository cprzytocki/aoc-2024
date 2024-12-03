package aoc

import (
	lib "aoc/2024/lib"
	"slices"
)

func parseLists() ([]int, []int) {
	numbers := lib.ScanFileToNumbers("./day1/input")

	left := []int{}
	right := []int{}

	for i := 0; i < len(numbers); i++ {
		if i%2 == 0 {
			left = append(left, numbers[i])
		} else {
			right = append(right, numbers[i])
		}
	}

	return left, right
}

func Day1() int {
	left, right := parseLists()
	slices.Sort(left)
	slices.Sort(right)

	diffs := []int{}

	for i := 0; i < len(left); i++ {

		diff := lib.Abs(left[i] - right[i])
		diffs = append(diffs, diff)
	}

	return lib.SumInts(diffs)
}

func occuranceScore(num int, list []int) int {
	count := 0

	for i := 0; i < len(list); i++ {
		if list[i] == num {
			count++
		}
	}

	return count * num
}

func Day1Part2() int {
	left, right := parseLists()

	score := 0

	for i := 0; i < len(left); i++ {
		score += occuranceScore(left[i], right)
	}

	return score
}
