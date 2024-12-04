package aoc

import (
	lib "aoc/2024/lib"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func scanFileToIntPair(filepath string, regex *regexp.Regexp) []lib.Pair {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := []lib.Pair{}

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(match[2])
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, lib.Pair{Num1: num1, Num2: num2})
		}
	}

	return results
}

func scanFileToIntPairWithDoDont(filepath string, regex *regexp.Regexp) []lib.Pair {
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	results := []lib.Pair{}

	do := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if match[0] == "do()" {
				do = true
			} else if match[0] == "don't()" {
				do = false
			} else if do {
				num1, err := strconv.Atoi(match[1])
				if err != nil {
					log.Fatal(err)
				}
				num2, err := strconv.Atoi(match[2])
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, lib.Pair{Num1: num1, Num2: num2})
			}
		}
	}
	return results
}

func Day3() int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	pairs := scanFileToIntPair("./day3/input", regex)
	total := 0
	for _, pair := range pairs {
		total += lib.MultiplyPair(pair)
	}

	return total
}

func Day3Part2() int {
	// if go supported lookbehinds we could do this in one regex
	// (?:(?<!don't\(\).+)|(?<=do\(\).+))(?:mul\((\d+),(\d+)\))
	regex := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	pairs := scanFileToIntPairWithDoDont("./day3/input", regex)
	total := 0
	for _, pair := range pairs {
		total += lib.MultiplyPair(pair)
	}

	return total
}
