package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Grid struct {
	mat    [][]string
	rowLen int
	colLen int
	col    int
	row    int
}

func (grid *Grid) checkLimits(row, col int) bool {
	fmt.Println("checklimits(): ", row, col)
	if row < 0 || col < 0 || row >= grid.rowLen || col >= grid.colLen {
		return true
	}
	return false
}

func (grid *Grid) moveUp() bool {
	nextRow := grid.row - 1
	fmt.Println("MoveUp: ", nextRow, " ", grid.col)
	if grid.checkLimits(nextRow, grid.col) {
		grid.mat[grid.row][grid.col] = "X"
		return false
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return true
	}
	grid.mat[grid.row][grid.col] = "X"
	grid.row = nextRow
	grid.mat[nextRow][grid.col] = "^"
	return true
}

func (grid *Grid) moveDown() bool {
	fmt.Println("moveDown(): ", grid.row, " ", grid.col)
	nextRow := grid.row + 1
	if grid.checkLimits(nextRow, grid.col) {
		grid.mat[grid.row][grid.col] = "X"
		return false
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return true
	}
	grid.mat[grid.row][grid.col] = "X"
	grid.row = nextRow
	grid.mat[nextRow][grid.col] = "v"
	return true
}

func (grid *Grid) moveLeft() bool {
	nextCol := grid.col - 1
	fmt.Println("MoveLeft: ", grid.row, " col:", nextCol)
	if grid.checkLimits(grid.row, nextCol) {
		grid.mat[grid.row][grid.col] = "X"
		return false
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return true
	}
	grid.mat[grid.row][grid.col] = "X"
	grid.col = nextCol
	grid.mat[grid.row][nextCol] = "<"
	return true
}

func (grid *Grid) moveRight() bool {
	nextCol := grid.col + 1
	fmt.Println("MoveRight: ", grid.row, " col:", nextCol)
	if grid.checkLimits(grid.row, nextCol) {
		grid.mat[grid.row][grid.col] = "X"
		return false
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return true
	}
	grid.mat[grid.row][grid.col] = "X"
	grid.col += 1
	grid.mat[grid.row][nextCol] = ">"
	return true
}

func getStartingPosIfPresent(cols []string) int {
	for i, c := range cols {
		fmt.Println(c)
		fmt.Println(c == "^")
		if c == "^" || c == ">" || c == "<" || c == "v" {
			fmt.Println(c == "^")
			fmt.Println(i)
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
	pos := &grid.mat[grid.row][grid.col]
	fmt.Println("Turn90Deg", pos)
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
	pos := grid.mat[grid.row][grid.col]
	fmt.Println("POS move(): ", pos)
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
		fmt.Printf("Characters: %q\n", cols)
		mat[i] = cols
	}

	grid := Grid{mat, len(rows), len(rows[0]) - 1, startCol, startRow}
	fmt.Println("col start(): ", grid.col)
	fmt.Println("row start(): ", grid.row)
	fmt.Println("colLength start(): ", grid.colLen)
	fmt.Println("rowLenght start(): ", grid.rowLen)

	for true {
		res := grid.move()
		fmt.Println(res)
		if res == false {
			break
		}
	}

	total := 0
	for _, row := range grid.mat {
		for _, col := range row {
			fmt.Print(col)
			if col == "X" {
				total++
			}
		}
		fmt.Println()
	}
	fmt.Println("Total: ", total)
}
