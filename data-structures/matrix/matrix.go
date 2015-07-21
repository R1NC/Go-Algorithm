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
		case 0:
			if column < rowSize - row - 1 {
				column++
			} else {
				row++
				dir++
			}
		case 1:
			if row < column {
				row++
			} else {
				column--
				dir++
			}
		case 2:
			if column > rowSize - row - 1 {
				column--
			} else {
				row--
				dir++
			}
		case 3:
			if row > column + 1 {
				row--
			} else {
				column++
				dir++
			}
		}
	}
}