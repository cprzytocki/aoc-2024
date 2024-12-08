package aoc

import (
	lib "aoc/2024/lib"
	"fmt"
)

type Direction int

const (
	Up Direction = iota
	Down
	Right
	Left
)

type Player struct {
	y int
	x int
	d Direction
}

func (p *Player) turn() {
	switch p.d {
	case Up:
		p.d = Right
	case Right:
		p.d = Down
	case Down:
		p.d = Left
	case Left:
		p.d = Up
	default:
		panic("Invalid direction")
	}
}

func (p *Player) next() (y, x int) {
	switch p.d {
	case Up:
		return p.y - 1, p.x
	case Down:
		return p.y + 1, p.x
	case Right:
		return p.y, p.x + 1
	case Left:
		return p.y, p.x - 1
	}
	panic("Invalid direction")
}

type Square int

const (
	Untravelled Square = iota
	Travelled
	Wall
	DONE
)

func checkSquare(m [][]string, y, x int) Square {
	if y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
		return DONE
	}

	str := m[y][x]

	switch str {
	case ".":
		return Untravelled
	case "#":
		return Wall
	case "X":
		return Travelled
	}
	panic("Invalid square: " + str + " " + fmt.Sprint(y) + " " + fmt.Sprint(x))
}

func printMatrix(m [][]string) {
	for _, row := range m {
		fmt.Println(row)
	}
}

func getInitialDirection(str string) Direction {
	switch str {
	case "^":
		return Up
	case "v":
		return Down
	case ">":
		return Right
	case "<":
		return Left
	}
	panic("Invalid initial direction: " + str)
}

func countTravelled(rows [][]string) int {
	count := 0
	for _, row := range rows {
		for _, cell := range row {
			if cell == "X" {
				count++
			}
		}
	}
	return count
}

func Day6() int {
	rows, py, px := lib.ScanFileToMatrix2("day6/input")
	player := Player{y: py, x: px, d: getInitialDirection(rows[py][px])}
	fmt.Println(player)

	rows[py][px] = "X"
	done := false
	for !done {
		y, x := player.next()
		res := checkSquare(rows, y, x)
		if res == DONE {
			done = true
		} else if res == Untravelled || res == Travelled {
			player.y = y
			player.x = x
			if res == Untravelled {
				rows[y][x] = "X"
			}
		} else if res == Wall {
			player.turn()
		} else {
			panic("Unahandeld res: " + fmt.Sprint(res))
		}
	}

	printMatrix(rows)
	// write to file
	// lib.WriteMatrixToFile("day6/output", rows)
	return countTravelled(rows)
}
