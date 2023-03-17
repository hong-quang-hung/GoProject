package utils

import (
	"fmt"
)

func Print(grid [][]int) {
	fmt.Println(fmt.Sprint(grid))
}

func PrintMatrix(grid [][]int) {
	for _, g := range grid {
		fmt.Println(fmt.Sprint(g))
	}
}
