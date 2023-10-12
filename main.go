package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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
const MOVE_CURSOR = "\x1B[H"

// position describes a positions state alive (true) or dead (false).
// day and night are used to clarify passing of time, day after night after day
type position struct {
	day   bool
	night bool
}

func makeNewBoard(initialBoard [][]bool) [][]position {
	newBoard := make([][]position, len(initialBoard))
	for row := range newBoard {
		newBoard[row] = make([]position, len(initialBoard[row]))
		for col := 0; col < len(initialBoard[row]); col++ {
			newBoard[row][col].day = initialBoard[row][col]
		}
	}
	return newBoard
}

func offBoard(board [][]position, row, col int) bool {
	if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) {
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

func gameOfLife(board [][]position) {
	day := true
	var generation uint

	for ; generation < 100; generation++ {
		printBoard(board, day, generation) ////

		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[row]); col++ {
				applyRules(board, row, col, day)
			}
		}

		day = !day
	}
}

func printPosition(alive bool) {
	if alive {
		fmt.Printf("%v  %v", BGREEN, RESET)
	} else {
		fmt.Printf("%v  %v", BRED, RESET)
	}
}

func printBoard(board [][]position, day bool, generation uint) {
	fmt.Printf("%v%v%vGame of Life%v\n\n", MOVE_CURSOR, BOLD, UNDERLINE, RESET)

	for row := 0; row < len(board); row++ {
		// fmt.Printf("%-3v ", row) // row index
		for col := 0; col < len(board[row]); col++ {
			if day {
				printPosition(board[row][col].day)
			} else {
				printPosition(board[row][col].night)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\ngeneration: %v\n\n", generation)
	time.Sleep(150 * time.Millisecond) // add -s flag!!!!!!!!!
}

func usageError(message string, err error) {
	fmt.Printf("%vERROR %v %v%v\n", RED, message, err, RESET)
	// printUsage()//!!
	os.Exit(1)
}

func errorExit(message string) {
	fmt.Printf("%vERROR %v %v\n", RED, message, RESET)
	// printUsage()//!!
	os.Exit(1)
}

func loadBoard(filename string) [][]position {
	filepath := "boards/" + filename + ".txt"

	readFile, err := os.Open(filepath)
	if err != nil {
		usageError("Invalid filepath: "+filepath, err)
	}

	var board [][]position
	var row int

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		board = append(board, []position{})
		line := fileScanner.Text()
		for col := 0; col < len(line); col++ {
			if line[col] == '0' {
				board[row] = append(board[row], position{false, false})
			} else if line[col] == '1' {
				board[row] = append(board[row], position{true, false})
			} else {
				errorExit("Invalid value in file: " + string(line[col]))
			}
		}
		row++
	}

	readFile.Close()
	return board
}

func randomBoard(size, percentAlive int) [][]position {
	var board [][]position
	// var row uint
	// var col uint

	rand.Seed(time.Now().UnixNano()) // -s flag!!!!!!

	for row := 0; row < size; row++ {
		board = append(board, []position{})
		for col := 0; col < size; col++ {
			random := rand.Intn(100)
			if random < percentAlive {
				board[row] = append(board[row], position{true, false})
			} else {
				board[row] = append(board[row], position{false, false})
			}
		}
	}
	return board
}

func main() {
	fmt.Printf("\033[H\033[2J") // Clear screen

	// filenme := "leetcode1"
	// filenme := "beacon"
	// board := loadBoard(filenme)
	board := randomBoard(42, 42) // flags!!!!!!!

	gameOfLife(board)
}

// ## To run enter:
// go run main.go
