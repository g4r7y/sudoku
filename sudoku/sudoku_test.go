package sudoku

import (
	"testing"
)


func TestVerifySudoku(t *testing.T) {
	sudoku := [][]int{
		{5,2,4,9,8,1,3,7,6},
		{1,3,8,6,2,7,4,5,9},
		{7,6,9,5,3,4,2,8,1},
		{4,7,2,3,1,8,9,6,5},
		{6,5,1,4,9,2,7,3,8},
		{8,9,3,7,5,6,1,4,2},
		{9,4,5,1,6,3,8,2,7},
		{3,8,6,2,7,9,5,1,4},
		{2,1,7,8,4,5,6,9,3},
	}
	if !VerifySudoku(sudoku) {
		t.Fatal("VerifySudoku: valid sudoku should return true, got false")
	}

	sudoku = [][]int{
		{5,2,4,9,8,1,3,7,6},
		{1,3,8,6,2,7,4,5,9},
		{7,6,9,5,3,4,2,8,1},
		{4,7,2,3,1,8,9,6,5},
		{6,5,1,4,9,2,7,3,8},
		{8,9,3,7,5,6,1,4,2},
		{9,4,5,1,6,3,8,2,7},
		{3,8,6,2,7,9,5,1,4},
		{2,1,7,8,4,5,6,9,10}, // 10!
	}
	if VerifySudoku(sudoku)==true {
		t.Fatal("VerifySudoku: sudoku with invalid number should return false, got true")
	}

	sudoku = [][]int{
		{5,5,4,9,8,1,3,7,6},//duplicate in row
		{1,3,8,6,2,7,4,5,9},
		{7,6,9,5,3,4,2,8,1},
		{4,7,2,3,1,8,9,6,5},
		{6,2,1,4,9,2,7,3,8},//duplicate in row
		{8,9,3,7,5,6,1,4,2},
		{9,4,5,1,6,3,8,2,7},
		{3,8,6,2,7,9,5,1,4},
		{2,1,7,8,4,5,6,9,3},
	}
	if VerifySudoku(sudoku)==true {
		t.Fatal("VerifySudoku: sudoku with duplicate in row should return false, got true")
	}

	sudoku = [][]int{
		{5,2,4,9,8,1,3,7,6},
		{1,3,8,6,2,7,4,5,9},
		{7,6,9,5,3,4,2,8,1},
		{4,7,2,3,1,8,9,6,5},
		{6,5,1,4,9,2,7,3,8},
		{8,9,3,7,5,6,1,4,2},
		{9,4,5,1,6,3,8,2,7},
		{3,8,6,2,7,9,5,1,4},
		{2,1,7,8,4,5,6,3,9},
		//dupe in cols ↑ ↑
	}
	if VerifySudoku(sudoku)==true {
		t.Fatal("VerifySudoku: sudoku with duplicate in column should return false, got true")
	}

	sudoku = [][]int{
		{5,2,4,9,8,1,3,7,6},
		{1,3,8,6,2,7,4,5,9},
		{7,6,9,5,3,4,2,8,1},
		{4,7,2,3,1,8,9,6,5},
		{6,5,1,4,9,2,7,3,8},
		{9,4,5,1,6,3,8,2,7}, // <- swapped these two rows so
		{8,9,3,7,5,6,1,4,2}, // <- there are duplicates in subboxes
		{3,8,6,2,7,9,5,1,4},
		{2,1,7,8,4,5,6,9,3},
	}
	if VerifySudoku(sudoku)==true {
		t.Fatal("VerifySudoku: sudoku with duplicate in subbox should return false, got true")
	}
}

func TestGenerateSudoko(t *testing.T) {

	for _ = range 1000 {
		sudoku := GenerateSudoku()

		if !VerifySudoku(sudoku) {
			for i := range len(sudoku) {
				t.Errorf("%v",sudoku[i])
			}
			t.Fatal("GenerateSudoku produced invalid sudoku.")
		}
	}
}
