package sudoku

import (
	"errors"
  "math/rand"
)


func VerifySudoku(sudoku [][]int) bool {
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

  for r:=0; r<gridSize; r++ {
    rowVals := make(map[int]bool)
    colVals := make(map[int]bool)

    for c:=0; c<gridSize; c++ {
      val := sudoku[r][c]

      // check range of val
      if val<1 || val>gridSize {
        // invalid digit
        return false
      }

      // check current row
      if rowVals[val] == true {
        //already encountered this val in this row
        return false
      }
      rowVals[val] = true

     // check current subbox
      subBoxId := subboxHeight * (r / subboxHeight) +  c / subboxWidth

      subBoxVals := subBoxes[subBoxId]
      if subBoxVals[val] == true {
        //already encountered this val in this subbox
        return false
      }
      subBoxVals[val] = true

      // check column by flipping r and c indices
      val = sudoku[c][r]
      if colVals[val] == true {
        //already encountered this val in this column
        return false
      }
      colVals[val] = true
    }
  }

  return true
}


func GenerateSudoku(gridSize int) [][]int {
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
    panic("Invalid sudoku size")
  }
  
  // initialise grid
  sudoku := make([][]int, gridSize)
  for i := 0; i<gridSize; i++ {
    sudoku[i] = make([]int, gridSize)
    for j := 0; j<gridSize; j++ {
      sudoku[i][j] = 0
    }
  }

  rowCandidates := NewNumSet(gridSize)

  for r := 0; r<gridSize; r++ {
    //reset available values for this new row
    rowCandidates.Fill()
    for c := 0; c<gridSize; c++ {
      

      colCandidates := NewNumSet(gridSize)
      colCandidates.Fill()
      for rn := 0; rn<r; rn++ { 
        colCandidates.ClearNum(sudoku[rn][c])
      }

      subboxCandidates := NewNumSet(gridSize)
      subboxCandidates.Fill()
      subboxTop := (r/subboxHeight) * subboxHeight
      subboxLeft := (c/subboxWidth) * subboxWidth
      for rb := subboxTop; rb<subboxTop+subboxHeight; rb++ {
        for cb := subboxLeft; cb < subboxLeft+subboxWidth; cb++ {
          if rb>r || (rb == r && cb>=c) {
            // don't look ahead
            continue
          }
          subboxCandidates.ClearNum(sudoku[rb][cb])
        }
      }


      combinedSet := Union(subboxCandidates, Union(colCandidates,rowCandidates))
      if (combinedSet.Count() > 0) {
        val,_ := PickRandom(combinedSet)
        sudoku[r][c] = val
        rowCandidates.ClearNum(val)
        subboxCandidates.ClearNum(val)
      } else {
          if (r+1) % subboxHeight == 0 {
            // retry subbox, move row to start of current subbox
            r = subboxTop-1
            
          } else {
            // retry row
            r--
          }
          break
        }
      }
  }

  return sudoku
}


func Union(set1 NumSet, set2 NumSet) NumSet {
  maxNum := set1.MaxNum()
  if set2.MaxNum() > maxNum {
    maxNum = set2.MaxNum()
  }
  result := NewNumSet(maxNum)
  for num:=1; num<=maxNum; num++ {
    if (set1.Has(num) && set2.Has(num)) {
      result.SetNum(num)
    }
  }
  return result
}


func PickRandom(set NumSet) (int,error) {
  setSize := set.Count()
  if setSize == 0 {
    return 0, errors.New("Set is empty")
  }
  
  // randomly choose from the set members
  randItem := rand.Intn(setSize) + 1
  
  // find from the possible set members and return its val
  count := 0
  for num:=1; num<=set.MaxNum(); num++ {
    if set.Has(num) == true {
      count++
      if (count == randItem) {
        return num, nil
      }
    } 
  }
  panic("PickRandom: unexpected error")
}
