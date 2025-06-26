package sudoku

import (
	"testing"
)

func TestGenerateSudoko(t *testing.T) {
	checkSudoku := func(sudoku [][]int) {
		if !VerifyPartialSudoku(sudoku) {
			for i := range len(sudoku) {
				t.Errorf("%v", sudoku[i])
			}
			t.Fatal("GenerateSudoku produced invalid sudoku.")
		}
	}

	for range 500 {
		sudoku := GenerateSudoku(9, DifficultyMedium)
		checkSudoku(sudoku)
	}

	for range 200 {
		sudoku := GenerateSudoku(6, DifficultyMedium)
		checkSudoku(sudoku)
	}

	for range 100 {
		sudoku := GenerateSudoku(4, DifficultyMedium)
		checkSudoku(sudoku)
	}

}
