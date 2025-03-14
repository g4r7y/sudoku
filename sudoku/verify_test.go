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


func TestVerifyPartialSudoku(t *testing.T) {
	sudoku := [][]int{
		{0,2,4,0,0,0,3,0,0},
		{1,0,8,0,2,0,4,0,9},
		{0,6,0,5,0,4,0,0,1},
		{0,7,2,3,1,0,0,0,5},
		{0,0,0,0,0,0,0,0,0},
		{8,0,0,0,5,0,1,0,0},
		{9,0,0,1,0,3,0,0,0},
		{3,8,6,2,7,9,0,0,0},
		{0,0,0,8,0,5,0,0,0},
	}
	if !VerifyPartialSudoku(sudoku) {
		t.Fatal("VerifyPartialSudoku: valid sudoku should return true, got false")
	}
	
	sudoku = [][]int{
		{0,2,4,0,0,0,3,0,-10},// invalid number
		{1,0,8,0,2,0,4,0,9},
		{0,6,0,5,0,4,0,0,1},
		{0,7,2,3,1,0,0,0,5},
		{0,0,0,0,0,0,0,0,0},
		{8,0,0,0,5,0,1,0,0},
		{9,0,0,1,0,3,0,0,0},
		{3,8,6,2,7,9,0,0,0},
		{0,0,0,8,0,5,0,0,0},
	}
	if VerifyPartialSudoku(sudoku)==true {
		t.Fatal("VerifyPartialSudoku: sudoku with invalid number should return false, got true")
	}

	sudoku = [][]int{
		{0,2,4,0,0,0,3,0,0},
		{1,0,8,0,2,0,0,0,9},
		{0,6,0,5,0,4,4,0,1},//duplicate 4 in row
		{0,7,2,3,1,0,0,0,5},
		{0,0,0,0,0,0,0,0,0},
		{8,0,0,0,5,0,1,0,0},
		{9,0,0,1,0,3,0,0,0},
		{3,8,6,2,7,9,0,0,0},
		{0,0,0,8,0,5,0,0,0},
	}
	if VerifyPartialSudoku(sudoku)==true {
		t.Fatal("VerifyPartialSudoku: sudoku with duplicate in row should return false, got true")
	}

	sudoku = [][]int{
		{0,2,4,0,0,0,3,0,0},
		{1,0,8,0,2,0,4,0,9},
		{0,6,0,5,0,4,0,0,1},
		{0,7,2,3,1,0,0,0,5},
		{0,0,0,0,0,0,0,0,0},
		{8,0,0,0,5,0,1,0,0},
		{9,0,0,1,0,3,0,0,0},
		{3,8,6,2,7,9,0,0,0},
		{0,0,0,8,0,5,0,0,1},//duplicate 1 in last column
	}
	if VerifyPartialSudoku(sudoku)==true {
		t.Fatal("VerifyPartialSudoku: sudoku with duplicate in column should return false, got true")
	}


	sudoku = [][]int{
		{0,2,4,0,0,0,3,0,0},
		{1,0,8,0,2,0,4,0,9},
		{0,6,0,5,0,4,0,0,3},//duplicate 3 in top right subblock
		{0,7,2,3,1,0,0,0,5},
		{0,0,0,0,0,0,0,0,0},
		{8,0,0,0,5,0,1,0,0},
		{9,0,0,1,0,3,0,0,0},
		{3,8,6,2,7,9,0,0,0},
		{0,0,0,8,0,5,0,0,0},
	}
	if VerifyPartialSudoku(sudoku)==true {
		t.Fatal("VerifyPartialSudoku: sudoku with duplicate in subblock should return false, got true")
	}
}