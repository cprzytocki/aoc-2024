package aoc

import (
	lib "aoc/2024/lib"
	"regexp"
	"strings"
)

func getDimensions(matrix [][]string) (width, height int) {
	width, height = len(matrix[0]), len(matrix)
	return
}

func getHorizontalLines(matrix [][]string) []string {
	_, height := getDimensions(matrix)

	horizontalLines := []string{}
	for i := 0; i < height; i++ {
		horizontalLines = append(horizontalLines, strings.Join(matrix[i], ""))
	}
	return horizontalLines
}

func getVerticalLines(matrix [][]string) []string {
	width, height := getDimensions(matrix)
	verticalLines := []string{}

	for i := 0; i < width; i++ {
		verticalLine := ""
		for j := 0; j < height; j++ {
			verticalLine += matrix[j][i]
		}
		verticalLines = append(verticalLines, verticalLine)
	}
	return verticalLines
}

func rotateMatrix45(matrix [][]string, clockwise bool) [][]string {
	width, height := getDimensions(matrix)
	newSize := height + width

	// allocate memory for the new matrix
	newMatrix := make([][]string, newSize)
	for i := 0; i < newSize; i++ {
		newMatrix[i] = make([]string, newSize)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if clockwise {
				newMatrix[i+j][j] = matrix[i][j]
			} else {
				newMatrix[width-i+j][j] = matrix[i][j]
			}
		}
	}

	return newMatrix
}

func getDiagonalLines(matrix [][]string, minLength int) (diagnalLines1 []string, diagnalLines2 []string) {
	matrix1 := rotateMatrix45(matrix, true)
	matrix2 := rotateMatrix45(matrix, false)

	for i := 0; i < len(matrix1); i++ {
		str := strings.Join(matrix1[i], "")
		if len(str) >= minLength {
			diagnalLines1 = append(diagnalLines1, str)
		}
	}

	for i := 0; i < len(matrix2); i++ {
		str := strings.Join(matrix2[i], "")
		if len(str) >= minLength {
			diagnalLines2 = append(diagnalLines2, str)
		}
	}

	return diagnalLines1, diagnalLines2
}

func countMatches(lines []string, patterns []*regexp.Regexp) int {
	count := 0
	for _, line := range lines {
		for _, pattern := range patterns {
			match := pattern.FindAllStringSubmatch(line, -1)
			count += len(match)
		}
	}
	return count
}

func Day4() int {
	matrix := lib.ScanFileToMatrix("day4/input")
	minLength := 4

	horizontalLines := getHorizontalLines(matrix)
	verticalLines := getVerticalLines(matrix)
	diagnalLines1, diagnalLines2 := getDiagonalLines(matrix, minLength)

	// DEBUG
	// fmt.Println("horizontalLines", horizontalLines)
	// fmt.Println("verticalLines", verticalLines)
	// fmt.Println("diagnalLines1", diagnalLines1)
	// fmt.Println("diagnalLines2", diagnalLines2)

	regexAll := []*regexp.Regexp{regexp.MustCompile(`XMAS`), regexp.MustCompile(`SAMX`)}

	matches := countMatches(horizontalLines, regexAll) +
		countMatches(verticalLines, regexAll) +
		countMatches(diagnalLines1, regexAll) +
		countMatches(diagnalLines2, regexAll)

	return matches
}
