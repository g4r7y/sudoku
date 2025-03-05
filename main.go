package main

import (
	"codeberg.org/grty/sudoku/sudoku"
	"fmt"
)

func main() {
	sudoku := sudoku.GenerateSudoku(9)
	for r := range sudoku {
		fmt.Printf("%v\n", sudoku[r])
	}
}
