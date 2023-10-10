package main

import (
	"fmt"
)

// Print Color
const RESET = "\x1B[0m"
const BOLD = "\x1B[1m"
const UNDERLINE = "\x1B[4m"
const RED = "\x1B[31m"
const GREEN = "\x1B[32m"
const YELLOW = "\x1B[33m"
const BRED = "\x1B[41m"
const BGREEN = "\x1B[42m"

// position describes a positions state alive (true) or dead (false).
// day and night are used to clarify passing of time, day after night after day
type position struct {
	day   bool
	night bool
}

func makeNewBoard(initialBoard [][]int) [][]position {
	newBoard := make([][]position, len(initialBoard))
	for row := range newBoard {
		newBoard[row] = make([]position, len(initialBoard[0]))
		for col := 0; col < len(initialBoard[0]); col++ {
			if initialBoard[row][col] == 1 {
				newBoard[row][col].day = true
			}
		}
	}
	return newBoard
}

func offBoard(board [][]position, row, col int) bool {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[0]) {
		return true
	}
	return false
}

func countNeighbours(board [][]position, row, col int, day bool) uint {
	var neighbours uint
	for rowNeighbour := row - 1; rowNeighbour <= row+1; rowNeighbour++ {
		for colNeighbour := col - 1; colNeighbour <= col+1; colNeighbour++ {
			if (rowNeighbour == row && colNeighbour == col) || offBoard(board, rowNeighbour, colNeighbour) {
				continue
			}
			if day {
				if board[rowNeighbour][colNeighbour].day {
					neighbours++
				}
			} else {
				if board[rowNeighbour][colNeighbour].night {
					neighbours++
				}
			}

		}
	}
	return neighbours
}

func deadOrAlive(board [][]position, row, col int, day bool, neighbours uint) {
	if day {
		if board[row][col].day {
			if neighbours == 2 || neighbours == 3 {
				board[row][col].night = true
			} else {
				board[row][col].night = false
			}
		} else { // was dead
			if neighbours == 3 {
				board[row][col].night = true
			} else {
				board[row][col].night = false
			}
		}
	} else {
		if board[row][col].night {
			if neighbours == 2 || neighbours == 3 {
				board[row][col].day = true
			} else {
				board[row][col].day = false
			}
		} else { // was dead
			if neighbours == 3 {
				board[row][col].day = true
			} else {
				board[row][col].day = false
			}
		}
	}
}

func applyRules(board [][]position, row, col int, day bool) {
	neighbours := countNeighbours(board, row, col, day)
	deadOrAlive(board, row, col, day, neighbours)
}

func gameOfLife(initialBoard [][]int) {
	day := true
	board := makeNewBoard(initialBoard)

	printBoard(board, day) ////

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			applyRules(board, row, col, day)
		}
	}

	if day {
		day = false
	} else {
		day = true
	}
	printBoard(board, day) ////
}

func printPosition(alive bool) {
	if alive {
		fmt.Printf("%v  %v", BGREEN, RESET)
	} else {
		fmt.Printf("%v  %v", BRED, RESET)
	}
}

func printBoard(board [][]position, day bool) {
	for row := 0; row < len(board); row++ {
		// fmt.Printf("%-3v ", row)//
		for col := 0; col < len(board[0]); col++ {
			if day {
				printPosition(board[row][col].day)
			} else {
				printPosition(board[row][col].night)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Printf("\033[H\033[2J") // Clear screen
	fmt.Printf("%v%vGame of Life%v\n\n", BOLD, UNDERLINE, RESET)

	initialBoard := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	// initialBoard := [][]int{{1, 1}, {1, 0}}
	gameOfLife(initialBoard)

}

// ## To run enter:
// go run main.go
