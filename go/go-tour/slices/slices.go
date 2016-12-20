package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	matrix := make([][]uint8, dy)
	for i := range matrix {
		matrix[i] = make([]uint8, dx)
	}
	for i, row := range matrix {
		for j := range row {
			row[j] = uint8((i + j) / 2)
		}
	}
	return matrix
}

func main() {
	pic.Show(Pic)
}
