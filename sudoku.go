package main

import (
  "math/rand"
  "math"
  "errors"
)

type NumSet struct {
  a []bool
  cnt int
}

func NewNumSet(size int) NumSet {
  return NumSet {
    a: make([]bool, size),
    cnt: 0,
  }
}

func (set* NumSet) Fill() {
  for i := range len(set.a) {
    set.a[i] = true
  }
  set.cnt = len(set.a)
}

func (set* NumSet) SetVal(num int) {
  if num < 1 || num > len(set.a) {
    panic("SetVal called with number out of range")
  }
  set.a[num-1] = true
  set.cnt++
}

func (set* NumSet) ClearVal(num int) {
  if num < 1 || num > len(set.a) {
    panic("ClearVal called with number out of range")
  }
  set.a[num-1] = false
  set.cnt--
}

func (set *NumSet) Has(num int) bool {
  return num>0 && num<len(set.a)+1 && set.a[num-1] == true
}

func (set *NumSet) Count() int {
  return set.cnt
}

func (set *NumSet) PickRandom() (int,error) {
  if set.Count() == 0 {
    return 0, errors.New("Set is empty")
  }
  
  // randomly choose item out of the set
  randItem := rand.Intn(set.Count()) + 1
  
  // find the matching item, return the val
  count := 0
  for i := range len(set.a) {
    if set.a[i] == true {
      count++
      if (count == randItem) {
        return i + 1, nil
      }
    } 
  }
  
  panic("PickRandom: unexpected error")
}

func Union(set1 NumSet, set2 NumSet) NumSet {
  // just in case sets are different lengths (they shouldn't be in this though)
  biggest := int(math.Max(float64(len(set1.a)), float64(len(set2.a))))
  result := NewNumSet(biggest)
  for i := range biggest {
    if (set1.Has(i+1) && set2.Has(i+1)) {
      result.SetVal(i+1)
    }
  }
  return result
}

// todo make this an arg
const gridSize int = 9
const subboxWidth int = 3
const subboxHeight int = 3

func GenerateSudoku() [][]int {
  
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
        colCandidates.ClearVal(sudoku[rn][c])
      }

      subboxCandidates := NewNumSet(gridSize)
      subboxCandidates.Fill()
      subboxTop := int(math.Floor(float64(r/subboxHeight))) * subboxHeight
      subboxLeft := int(math.Floor(float64(c/subboxWidth))) * subboxWidth
      for rb := subboxTop; rb<subboxTop+subboxHeight; rb++ {
        for cb := subboxLeft; cb < subboxLeft+subboxWidth; cb++ {
          if rb>r || (rb == r && cb>=c) {
            // don't look ahead
            continue
          }
          subboxCandidates.ClearVal(sudoku[rb][cb])
        }
      }


      union := Union(Union(colCandidates, rowCandidates), subboxCandidates)
      if (union.Count() > 0) {
        val,_ := union.PickRandom()
        sudoku[r][c] = val
        rowCandidates.ClearVal(val)
        subboxCandidates.ClearVal(val)
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



func VerifySudoku(sudoku [][]int) bool {
  gridSize := len(sudoku)

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
      subBoxId := int(3 * math.Floor(float64(r) / 3) + math.Floor(float64(c) / 3))
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
