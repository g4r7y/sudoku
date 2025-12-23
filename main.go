package main

import (
	"errors"
	"fmt"
	"github.com/g4r7y/sudoku/sudoku"
	"io"
	"os"
	"regexp"
	"strconv"
)

const emptyCell = " "

func printSudoku(grid [][]int) {
	topBorder := "╔═"
	bottomBorder := "╚═"
	horizGrid := "╟─"
	for x := range grid[0] {
		if x>0 {
			topBorder += "══╤═"
			bottomBorder += "══╧═"
			horizGrid += "──┼─"
		}
	}
	topBorder += "══╗"
	bottomBorder += "══╝"
	horizGrid += "──╢"

	fmt.Println(topBorder)

	for r := range grid {
		if r > 0 {
			fmt.Println(horizGrid)
		}
		fmt.Print("║ ")
		for c := range grid[r] {
			val := strconv.Itoa(grid[r][c])
			if val == "0" {
				val = emptyCell
			}
			spacer := ""
			if c>0 {
				spacer = " │ "
			}
			fmt.Printf("%v%v", spacer, val)
		}
		fmt.Println(" ║")
	}

	for range grid[0] {
	}
	fmt.Println(bottomBorder)
}

func generate(size int, difficulty sudoku.Difficulty) {
	grid := sudoku.GenerateSudoku(size, difficulty)
	printSudoku(grid)
}

func solve(grid [][]int) {
	fmt.Println("Sudoku:")
	printSudoku(grid)
	solution, err := sudoku.Solve(grid)
	if err != nil {
		fatal(err)
	}
	fmt.Println("Solution:")
	printSudoku(solution)
}

func main() {

	args := os.Args[1:]

	fi, err := os.Stdin.Stat()
	if err != nil {
		fatal(err)
	}
	// read from piped command e.g. cat a.txt | sudoko)
	//  or from piped file (e.g. sudoku < a.txt)
	if (fi.Mode()&os.ModeNamedPipe != 0) || fi.Size() > 0 {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			fatal(err)
		}
		args = append(args, string(stdin))
	}

	if len(args) == 0 || args[0] == "generate" {
		size := 9
		difficulty := sudoku.DifficultyEasy
		if len(args) > 1 {
			mode := args[1]
			val, err := strconv.Atoi(mode)
			if err == nil && (val == 4 || val == 6 || val == 9) {
				size = val
			} else {
				switch mode {
				case "easy":
					difficulty = sudoku.DifficultyEasy
				case "medium":
					difficulty = sudoku.DifficultyMedium
				case "hard":
					difficulty = sudoku.DifficultyHard
				default:
					fatal(errors.New("Sudoku generate: invalid argument"))
				}
			}
		}
		generate(size, difficulty)
	} else if args[0] == "solve" {
		if len(args) < 2 {
			fatal(errors.New("Sudoku solve: invalid arguments - sudoku string is required"))
		}

		grid, err := parseSudoku(args[1])
		if err != nil {
			fatal(errors.New("Sudoku solve: invalid arguments - sudoku string could not be recognised"))
		}
		solve(grid)
	} else {
		fatal(errors.New(""))
	}
}

func parseSudoku(input string) ([][]int, error) {
	re, _ := regexp.Compile(`([0-9\.` + emptyCell + `])`)
	matches := re.FindAllString(input, -1)
	var sudokuSize int
	switch len(matches) {
	case 81:
		sudokuSize = 9
	case 36:
		sudokuSize = 6
	case 16:
		sudokuSize = 4
	default:
		return nil, errors.New("Wrong number of digits")
	}

	grid := make([][]int, sudokuSize)
	for r := range sudokuSize {
		grid[r] = make([]int, sudokuSize)
		for c := range sudokuSize {
			numStr := matches[r*sudokuSize+c]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				num = 0
			}
			grid[r][c] = num
		}
	}
	return grid, nil
}

func fatal(err error) {
	fmt.Println(err)
	fmt.Println(`
╔═════════════╗
║ S U D O K U ║
╚═════════════╝

Usage:

sudoku generate [type]

Generates a random sudoko.
type: easy | medium | hard | 4 | 6 | 9 
Optional type specifies the difficulty or the size of the sudoku. Defaults to easy 9x9. If number is specified, an easy sudoku of size 4x4, 6x6 or 9x9 will be generated accordingly. If difficulty is specified, a 9x9 sudoku will be generated with the specified difficulty.
Example:
  sudoku generate hard

sudoku solve <sudoku-grid>

Solves a sudoko.
sudoku-grid: sequence of numbers that form the grid, reading left to right, row by row. Numbers can be delimited with punctuation or spaces. Empty cells should be a '0' or '.' character
Example:
sudoku solve "4 3 . 1, 1 . 3 ., . 4 1 2, 2 1 . 3"
sudoku solve "4301,1030,0412,2103"

Piping of arguments is supported, for example:
cat sudoku.txt | sudoku solve
sudoku solve < sudoku.txt
sudoku generate | sudoku solve`)

	os.Exit(1)
}
