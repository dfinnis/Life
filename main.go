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

// func offBoard(board [][]int, m, n int) bool {
// 	if m < 0 || n < 0 || m >= len(board) || n >= len(board[0]) {
// 		return true
// 	}
// 	return false
// }

// func countNeighbours(board [][]int, m, n int) int {
// 	neighbours := 0
// 	for mNeighbour := m - 1; mNeighbour <= m+1; mNeighbour++ {
// 		for nNeighbour := n - 1; nNeighbour <= n+1; nNeighbour++ {
// 			if (mNeighbour == m && nNeighbour == n) || offBoard(board, mNeighbour, nNeighbour) {
// 				continue
// 			}
// 			neighbours += board[mNeighbour][nNeighbour]
// 		}
// 	}
// 	return neighbours
// }

// func isAlive(wasAlive, neighbours int) int {
// 	if wasAlive == 1 {
// 		if neighbours == 2 || neighbours == 3 {
// 			return 1
// 		}
// 	} else {
// 		if neighbours == 3 {
// 			return 1
// 		}
// 	}
// 	return 0
// }

// func applyRules(board [][]int, m, n int) int {
// 	neighbours := countNeighbours(board, m, n)
// 	return isAlive(board[m][n], neighbours)
// }

func gameOfLife(initialBoard [][]int) {
	day := true
	board := makeNewBoard(initialBoard)

	// for m := 0; m < len(board); m++ {
	// 	for n := 0; n < len(board[0]); n++ {
	// 		newBoard[m][n] = applyRules(board, m, n)
	// 	}
	// }

	// for m := 0; m < len(board); m++ {
	// 	for n := 0; n < len(board[0]); n++ {
	// 		board[m][n] = newBoard[m][n]
	// 	}
	// }
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
	gameOfLife(initialBoard)

	// board = [][]int{{1, 1}, {1, 0}}
}

// ## To run enter:
// go run main.go
