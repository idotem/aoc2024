package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Position struct {
	col       int
	row       int
	direction string
}

type Grid struct {
	mat    [][]string
	rowLen int
	colLen int
	col    int
	row    int
	dir    string
}

func (pos *Position) Ordered() {

}

func (grid *Grid) checkLimits(row, col int) bool {
	if row < 0 || col < 0 || row >= grid.rowLen || col >= grid.colLen {
		return true
	}
	return false
}

func (grid *Grid) moveUp() bool {
	nextRow := grid.row - 1
	if grid.checkLimits(nextRow, grid.col) {
		return false
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return true
	}
	grid.row = nextRow
	return true
}

func (grid *Grid) moveDown() bool {
	nextRow := grid.row + 1
	if grid.checkLimits(nextRow, grid.col) {
		return false
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return true
	}
	grid.row = nextRow
	return true
}

func (grid *Grid) moveLeft() bool {
	nextCol := grid.col - 1
	if grid.checkLimits(grid.row, nextCol) {
		return false
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return true
	}
	grid.col = nextCol
	return true
}

func (grid *Grid) moveRight() bool {
	nextCol := grid.col + 1
	if grid.checkLimits(grid.row, nextCol) {
		return false
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return true
	}
	grid.col = nextCol
	return true
}

func getStartingPosIfPresent(cols []string) int {
	for i, c := range cols {
		if c == "^" || c == ">" || c == "<" || c == "v" {
			return i
		}
	}
	return -1
}

func (grid *Grid) checkBlockage(row, col int) bool {
	if grid.mat[row][col] == "#" {
		return true
	}
	return false
}

func (grid *Grid) turn90Degrees() {
	pos := &grid.dir
	switch *pos {
	case "^":
		*pos = ">"
	case ">":
		*pos = "v"
	case "v":
		*pos = "<"
	case "<":
		*pos = "^"
	default:
		return
	}

}

func (grid *Grid) move() bool {
	pos := grid.dir
	switch pos {
	case "^":
		return grid.moveUp()
	case ">":
		return grid.moveRight()
	case "v":
		return grid.moveDown()
	case "<":
		return grid.moveLeft()
	default:
		return false
	}
}

func main() {
	stream, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	startCol := -1
	startRow := -1
	flag := false
	rows := strings.Split(string(stream), "\n")
	mat := make([][]string, len(rows))
	for i, row := range rows {
		row = strings.Trim(row, "\r")
		cols := strings.Split(row, "")
		if flag == false {
			if startCol = getStartingPosIfPresent(cols); startCol != -1 {
				startRow = i
				flag = true
			}
		}
		mat[i] = cols
	}

	grid := Grid{mat, len(rows), len(rows[0]) - 1, startCol, startRow, mat[startRow][startCol]}
	var positions []Position
	firstFlag := true
	for true {
		if firstFlag {
			fmt.Println("FIRST")
			firstFlag = false
			grid.move()
			continue
		}
		oldPos := Position{grid.col, grid.row, grid.dir}
		if checkIfNewBlockageCreatesLoop(&grid) {
			positions = append(positions, oldPos)
		}
		grid.col = oldPos.col
		grid.row = oldPos.row
		grid.dir = oldPos.direction
		res := grid.move()
		if res == false {
			break
		}

	}

	fmt.Println("Positions: ", len(positions))
}

func checkIfNewBlockageCreatesLoop(grid *Grid) bool {
	var otherPos []Position
	posOfNewBlockage := addNewBlockageToGrid(grid)
	for true {
		newPosition := Position{grid.row, grid.col, grid.dir}
		res := grid.move()
		if res == false {
			return false
		}
		if posOfNewBlockage == newPosition || positionAlreadyPresent(otherPos, newPosition) {
			return true
		}
		otherPos = append(otherPos, newPosition)
	}
	return false

}

func addNewBlockageToGrid(grid *Grid) Position {
	position := Position{grid.row, grid.col, grid.dir}
	// pos := &grid.dir
	// switch *pos {
	// case "^":
	// 	grid.mat[grid.row-1][grid.col] = "#"
	// case ">":
	// 	grid.mat[grid.row][grid.col+1] = "#"
	// case "v":
	// 	grid.mat[grid.row+1][grid.col] = "#"
	// case "<":
	// 	grid.mat[grid.row][grid.col-1] = "#"
	// default:

	// }
	grid.turn90Degrees()
	return position
}

func positionAlreadyPresent(positions []Position, position Position) bool {
	// fmt.Println(positions)
	// fmt.Println(position)
	for _, pos := range positions {
		if pos == position {
			return true
		}
	}
	return false
}
