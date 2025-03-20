package sudoku

import (
  "math/rand"
)

type Difficulty int

const (
  DifficultyEasy Difficulty = iota
  DifficultyMedium
  DifficultyHard
)

func GenerateSudoku(gridSize int, difficulty Difficulty) [][]int {
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


  var cellsToHide int
  switch gridSize {
  case 4: 
    cellsToHide = 10 + rand.Intn(3)//show 4-6 numbers
  case 6: 
    cellsToHide = 22 + rand.Intn(3)//show 12-14 numbers
  case 9: 
    cellsToHide = 40 + rand.Intn(7)//show 41-47 numbers
    switch difficulty {
      case DifficultyEasy: cellsToHide -= 5
      case DifficultyHard: cellsToHide += 5
    }
  }

  for range cellsToHide {
    for ;; {
      r := rand.Intn(gridSize)
      c := rand.Intn(gridSize)
      if result[r][c] != 0 {
        result[r][c] = 0
        break;
      }
    }
  }
  
  return result
}

