package sudoku

func VerifySudoku(sudoku [][]int) bool {
	return verifySudoku(sudoku, false)
}

func VerifyPartialSudoku(sudoku [][]int) bool {
	return verifySudoku(sudoku, true)
}

func verifySudoku(sudoku [][]int, allowEmptyCells bool) bool {
	gridSize := len(sudoku)
	var subboxWidth, subboxHeight int
	switch gridSize {
	case 9:
		subboxWidth = 3
		subboxHeight = 3
	case 6:
		subboxWidth = 3
		subboxHeight = 2
	case 4:
		subboxWidth = 2
		subboxHeight = 2
	default:
		// invalid sudoku size
		return false
	}

	// array containing map for each subbox
	subBoxes := make([]map[int]bool, gridSize)
	for n := range subBoxes {
		subBoxes[n] = make(map[int]bool)
	}

	for r := 0; r < gridSize; r++ {
		rowVals := make(map[int]bool)
		colVals := make(map[int]bool)

		for c := 0; c < gridSize; c++ {
			val := sudoku[r][c]
			if !allowEmptyCells || val != 0 {
				// check range of val
				if val < 1 || val > gridSize {
					// invalid digit
					return false
				}

				// check current row
				if rowVals[val] {
					//already encountered this val in this row
					return false
				}
				rowVals[val] = true

				// check current subbox
				subBoxId := subboxHeight*(r/subboxHeight) + c/subboxWidth

				subBoxVals := subBoxes[subBoxId]
				if subBoxVals[val] {
					//already encountered this val in this subbox
					return false
				}
				subBoxVals[val] = true
			}

			// check column by flipping r and c indices
			val = sudoku[c][r]
			if !allowEmptyCells || val != 0 {
				if colVals[val] {
					//already encountered this val in this column
					return false
				}
				colVals[val] = true
			}
		}
	}

	return true
}
