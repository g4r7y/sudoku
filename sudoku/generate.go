package sudoku


func GenerateSudoku(gridSize int) [][]int {
  switch gridSize {
  case 9: 
  case 6:
  case 4:
  default: 
    panic("Invalid sudoku size")
  }
  
  // initialise grid with zeros
  sudoku := make([][]int, gridSize)
  for r := range gridSize{ 
    sudoku[r] = make([]int, gridSize)
  }

  // passing empty grid to solver will generate all cells of the sudoku
  result,_ := Solve(sudoku)

  return result
}

