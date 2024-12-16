package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

type SimplePosition struct {
	row int
	col int
}

type DirPosition struct {
	row int
	col int
	dir string
}

type Grid struct {
	mat    [][]string
	rowLen int
	colLen int
	col    int
	row    int
	dir    string
}

func (grid *Grid) checkLimits(row, col int) bool {
	fmt.Println("checklimits(): ", row, col)
	if row < 0 || col < 0 || row >= grid.rowLen || col >= grid.colLen {
		return true
	}
	return false
}

func (grid *Grid) moveUp() (SimplePosition, error) {
	nextRow := grid.row - 1
	fmt.Println("MoveUp: ", nextRow, " ", grid.col)
	if grid.checkLimits(nextRow, grid.col) {
		return SimplePosition{}, errors.New("Out of limits for array")
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return SimplePosition{nextRow, grid.col}, nil
	}
	grid.row = nextRow
	grid.mat[nextRow][grid.col] = "^"
	return SimplePosition{nextRow, grid.col}, nil
}

func (grid *Grid) moveDown() (SimplePosition, error) {
	fmt.Println("moveDown(): ", grid.row, " ", grid.col)
	nextRow := grid.row + 1
	if grid.checkLimits(nextRow, grid.col) {
		return SimplePosition{}, errors.New("Out of limits for array")
	}
	if grid.checkBlockage(nextRow, grid.col) {
		grid.turn90Degrees()
		return SimplePosition{nextRow, grid.col}, nil
	}
	grid.row = nextRow
	grid.mat[nextRow][grid.col] = "v"
	return SimplePosition{nextRow, grid.col}, nil
}

func (grid *Grid) moveLeft() (SimplePosition, error) {
	nextCol := grid.col - 1
	fmt.Println("MoveLeft: ", grid.row, " col:", nextCol)
	if grid.checkLimits(grid.row, nextCol) {
		return SimplePosition{}, errors.New("Out of limits for array")
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return SimplePosition{grid.row, nextCol}, nil
	}
	grid.col = nextCol
	grid.mat[grid.row][nextCol] = "<"
	return SimplePosition{grid.row, nextCol}, nil
}

func (grid *Grid) moveRight() (SimplePosition, error) {
	nextCol := grid.col + 1
	fmt.Println("MoveRight: ", grid.row, " col:", nextCol)
	if grid.checkLimits(grid.row, nextCol) {
		return SimplePosition{}, errors.New("Out of limits for array")
	}
	if grid.checkBlockage(grid.row, nextCol) {
		grid.turn90Degrees()
		return SimplePosition{grid.row, nextCol}, nil
	}
	grid.col += 1
	grid.mat[grid.row][nextCol] = ">"
	return SimplePosition{grid.row, nextCol}, nil
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
		grid.dir = ">"
		*pos = ">"
	case ">":
		grid.dir = "v"
		*pos = "v"
	case "v":
		grid.dir = "<"
		*pos = "<"
	case "<":
		grid.dir = "^"
		*pos = "^"
	default:
		return
	}

}

func (grid *Grid) move() (SimplePosition, error) {
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
		return SimplePosition{}, errors.New("Invalid move for position, default switch activated")
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

	grid := Grid{mat, len(rows), len(rows[0]) - 1, startCol, startRow, "^"}
	fmt.Println("col start(): ", grid.col)
	fmt.Println("row start(): ", grid.row)
	fmt.Println("colLength start(): ", grid.colLen)
	fmt.Println("rowLenght start(): ", grid.rowLen)

	var pathPositions []SimplePosition
	for true {
		res, err := grid.move()
		if err == nil {
			fmt.Println(err)
			break
		}
		pathPositions = append(pathPositions, res)
	}
	secondGrid := Grid{mat, len(rows), len(rows[0]) - 1, startCol, startRow, "^"}

	total := 0
	startPos := SimplePosition{startRow, startCol}
	for r, row := range secondGrid.mat {
		for c, col := range row {
			pos := SimplePosition{c, r}
			if !positionContainedInPathPositions(pathPositions, pos) ||
				pos == startPos || col == "#" {
				continue
			}
			var alreadyPassedPos []DirPosition
			for true {
				if positionAlreadyPresent(alreadyPassedPos, DirPosition{secondGrid.row, secondGrid.col, secondGrid.dir}) {
					total++
				}
				res, err := secondGrid.move()
				if err == nil {
					secondGrid.mat[secondGrid.row][secondGrid.col] = "."
					secondGrid.mat[startRow][startCol] = "^"
					secondGrid.col = startCol
					secondGrid.row = startRow
					break
				}
				alreadyPassedPos = append(alreadyPassedPos, DirPosition{res.row, res.col, secondGrid.dir})
			}

		}
		fmt.Println()
	}
	fmt.Println("Total: ", total)
}

func positionAlreadyPresent(pathPos []DirPosition, pos DirPosition) bool {
	for _, p := range pathPos {
		if p == pos {
			return true
		}
	}
	return false

}

func positionContainedInPathPositions(pathPos []SimplePosition, pos SimplePosition) bool {
	for _, p := range pathPos {
		if p == pos {
			return true
		}
	}
	return false
}
