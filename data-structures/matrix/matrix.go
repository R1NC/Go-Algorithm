package stack

import (
	"fmt"
)

func spiralTraverse(m [][]string) {
	rowSize := len(m[0])
	columnSize := len(m)
	visited := 0
	dir := 0
	row := 0
	column := 0
	for visited < rowSize * columnSize {
		fmt.Printf("%s\n", m[row][column])
		visited++
		switch dir%4 {
		case 0: //TOP
			if column < rowSize - row - 1 {
				column++
			} else {
				row++
				dir++
			}
		case 1: //RIGHT
			if row < column {
				row++
			} else {
				column--
				dir++
			}
		case 2: //BOTTOM
			if column > rowSize - row - 1 {
				column--
			} else {
				row--
				dir++
			}
		case 3: //LEFT
			if row > column + 1 {
				row--
			} else {
				column++
				dir++
			}
		}
	}
}