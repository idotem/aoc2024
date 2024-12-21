package main

import (
	"fmt"
	"math"

	"aoc/util"
)

type Position struct {
	I int
	J int
}

func main() {
	lines := util.ReadFile("../input.txt")
	antenna_map := make([][]rune, len(lines))
	positions := make(map[Position]struct{})
	antennas := make(map[rune]int)

	for i, line := range lines {
		runes := []rune(line)
		antenna_map[i] = runes
	}

	for i := 0; i < len(antenna_map); i++ {
		for j := 0; j < len(antenna_map[i]); j++ {
			if antenna_map[i][j] != '.' {
				newPositions := findNextAntennas(
					antenna_map,
					i,
					j,
				)
				for k, v := range newPositions {
					positions[k] = v
				}
				antennas[antenna_map[i][j]] = antennas[antenna_map[i][j]] + 1
			}
		}
	}

	totalValidAntennasAsAntinodes := 0
	for _, v := range antennas {
		if v > 1 {
			totalValidAntennasAsAntinodes += v
		}
	}

	fmt.Println(len(antennas))
	fmt.Println(antennas)
	fmt.Println(len(positions))
	fmt.Println(totalValidAntennasAsAntinodes)
	fmt.Println(len(positions) + totalValidAntennasAsAntinodes)
}

func findNextAntennas(
	antenna_map [][]rune,
	i, j int,
) map[Position]struct{} {
	positions := make(map[Position]struct{})
	lenI := len(antenna_map)
	lenJ := len(antenna_map[0])
	flag := true
	var pos1 Position
	var pos2 Position
	for i2 := i; i2 < lenI; i2++ {
		for j2 := 0; j2 < len(antenna_map[i]); j2++ {
			if i == i2 && flag {
				j2 = j + 1
				flag = false
				continue
			}
			if antenna_map[i2][j2] == antenna_map[i][j] {
				if j < j2 {
					diffI := int(math.Abs(float64(i2 - i)))
					diffJ := int(math.Abs(float64(j2 - j)))
					pos1I := i - diffI
					pos1J := j - diffJ
					pos2I := i2 + diffI
					pos2J := j2 + diffJ
					pos1 = Position{pos1I, pos1J}
					pos2 = Position{pos2I, pos2J}
					for checkBounds(pos1, lenI, lenJ) {
						if antenna_map[pos1I][pos1J] == '.' {
							positions[pos1] = struct{}{}
						}
						pos1I = pos1I - diffI
						pos1J = pos1J - diffJ
						pos1 = Position{pos1I, pos1J}
					}
					for checkBounds(pos2, lenI, lenJ) {
						if antenna_map[pos2I][pos2J] == '.' {
							positions[pos2] = struct{}{}
						}
						pos2I = pos2I + diffI
						pos2J = pos2J + diffJ
						pos2 = Position{pos2I, pos2J}
					}

				} else {
					diffI := int(math.Abs(float64(i - i2)))
					diffJ := int(math.Abs(float64(j - j2)))
					pos1I := i - diffI
					pos1J := j + diffJ
					pos2I := i2 + diffI
					pos2J := j2 - diffJ
					pos1 = Position{pos1I, pos1J}
					pos2 = Position{pos2I, pos2J}
					for checkBounds(pos1, lenI, lenJ) {
						if antenna_map[pos1I][pos1J] == '.' {
							positions[pos1] = struct{}{}
						}
						pos1I = pos1I - diffI
						pos1J = pos1J + diffJ
						pos1 = Position{pos1I, pos1J}
					}
					for checkBounds(pos2, lenI, lenJ) {
						if antenna_map[pos2I][pos2J] == '.' {
							positions[pos2] = struct{}{}
						}
						pos2I = pos2I + diffI
						pos2J = pos2J - diffJ
						pos2 = Position{pos2I, pos2J}
					}
				}
			}
		}
	}
	return positions
}

func checkBounds(pos Position, lenI, lenJ int) bool {
	if pos.I < 0 || pos.J < 0 || pos.I >= lenI || pos.J >= lenJ {
		return false
	}
	return true
}
