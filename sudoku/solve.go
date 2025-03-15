package sudoku

import (
	"errors"
  "fmt"
	"math/rand"
)

func Solve(inputSudoku [][]int) ([][]int, error) {
	gridSize := len(inputSudoku)
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
		return nil, fmt.Errorf("Sudoku has unsupported size")
	}
	if !VerifyPartialSudoku(inputSudoku) {
		return nil, fmt.Errorf("Sudoku is invalid")
	}

	//initialise result
	solution := make([][]int, gridSize)
	for row,_ := range solution {
		solution[row] = make([]int, gridSize)
	}

	var backtrack bool
	var rowTries,subboxTries int
	
  for r := 0; r<gridSize; {
		backtrack = false
    
		// initialise set of candidate numbers for this row
		rowCandidates := NewNumSet(gridSize)
    rowCandidates.Fill()
		// exclude any values in the input sudoku for this row
		for _,val := range inputSudoku[r] {
			if val > 0 {
				rowCandidates.ClearNum(val)
			}
		}

    for c := 0; c<gridSize; c++ {
			if inputSudoku[r][c] > 0 {
				// copy value from the input sudoku
				solution[r][c] = inputSudoku[r][c]
				continue
			}
      
			// build set of candidate numbers for current column
      colCandidates := NewNumSet(gridSize)
      colCandidates.Fill()
      for rn, _ := range inputSudoku {
				// exclude any values in the input sudoku for in this column
				if inputSudoku[rn][c] > 0 {
					colCandidates.ClearNum(inputSudoku[rn][c])
				}
				if rn<r {
					// exclude any values in this column so far in the solution
        	colCandidates.ClearNum(solution[rn][c])
				}
      }

			// build set of candidate numbers for current subbox
      subboxCandidates := NewNumSet(gridSize)
      subboxCandidates.Fill()
      subboxTop := (r/subboxHeight) * subboxHeight
      subboxLeft := (c/subboxWidth) * subboxWidth
      for rb := subboxTop; rb<subboxTop+subboxHeight; rb++ {
        for cb := subboxLeft; cb < subboxLeft+subboxWidth; cb++ {
          if rb<r || rb ==r && cb<c{
						// exclude any values placed in this subbox so far (not looking ahead of r or c)
          	subboxCandidates.ClearNum(solution[rb][cb])
					}
					// exclude any values in the input sudoku for this subbox 
					if inputSudoku[rb][cb] > 0 {
						subboxCandidates.ClearNum(inputSudoku[rb][cb])
					}
        }
      }


      combinedSet := Union(subboxCandidates, Union(colCandidates,rowCandidates))
      if (combinedSet.Count() > 0) {
        val,_ := PickRandom(combinedSet)
        solution[r][c] = val
        rowCandidates.ClearNum(val)
      } else {
				if rowTries < 9 {
					// back to start of row
					rowTries++
				} else if subboxTries < 3 && r > subboxTop {
					// back to start of current subbox
					subboxTries++
					rowTries = 0
					r = subboxTop
				} else {
					// back to square one!
					subboxTries = 0
					rowTries = 0
					r = 0
				}
				backtrack = true
				break
			}
		}


		if !backtrack {
			r++
			rowTries = 0
			if r % subboxHeight == 0 {
				// made it to the next subbox
				subboxTries = 0
			}
		}
  }

	return solution,nil
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