package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {

	gameOver := false
	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	turn := 1

	for gameOver != true {
		drawBoard(board)
		player := 0

		if turn%2 == 1 {
			player = 1
		} else {
			player = 2
		}

		selectedMove := 9
		if player == 1 {
			selectedMove = botMove(turn, player, board)
		} else {
			fmt.Println("Player " + strconv.Itoa(player) + " turn")
			selectedMove = promptForMove()
		}
		if selectedMove == 9 {
			return
		}

		board = executeMove(selectedMove, player, board)

		result := checkBoardForAWinner(board)
		if result > 0 {
			fmt.Printf("Player %d wins!\n\n", result)
			gameOver = true
		} else if turn == 9 {
			// Tie game example: 0 2 1 3 4 7 5 8 6
			fmt.Printf("Tie game!\n\n")
			gameOver = true
		} else {
			turn++
			cmd := exec.Command("clear") //Linux example, its tested
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	}
}

func drawBoard(board [9]int) {

	for i, v := range board {
		if v == 0 {
			// empty space. Display number
			fmt.Printf("%d", i)
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 10 {
			fmt.Printf("O")
		}
		// And now the decorations
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}

func promptForMove() int {
	fmt.Println("Select a move")
	var move int
	fmt.Scan(&move)

	return move
}

func executeMove(currentMove int, player int, board [9]int) [9]int {
	if board[currentMove] != 0 {
		fmt.Println("That tile (" + strconv.Itoa(currentMove) + ") is already selected. Please select another one.")
		currentMove := promptForMove()
		board = executeMove(currentMove, player, board)
	} else if player == 1 {
		board[currentMove] = 1
	} else if player == 2 {
		board[currentMove] = 10
	}

	for currentMove > 9 {
		fmt.Println("Please select a tile that exists. There are only 9 tiles.")
		currentMove = promptForMove()
	}

	return board
}

func checkBoardForAWinner(board [9]int) int {

	winnings := winningPossibilities(board)

	for _, check := range winnings {
		if check == 3 {
			return 1
		} else if check == 30 {
			return 2
		}
	}

	return 0
}

func winningPossibilities(board [9]int) [8]int {
	winnings := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	// Vertical winnings
	winnings[0] = board[0] + board[3] + board[6]
	winnings[1] = board[1] + board[4] + board[7]
	winnings[2] = board[2] + board[5] + board[8]

	// Horizontal winnings
	winnings[3] = board[0] + board[1] + board[2]
	winnings[4] = board[3] + board[4] + board[5]
	winnings[5] = board[6] + board[7] + board[8]

	// Diagonal winnings
	winnings[6] = board[0] + board[4] + board[8]
	winnings[7] = board[2] + board[4] + board[6]

	return winnings
}

func botMove(currentTurn int, player int, board [9]int) int {

	fmt.Println("Botmove!")

	if board[4] == 0 {
		executeMove(4, player, board)
		return 4
	}

	if player == 2 {
		winnings := winningPossibilities(board)
		for i, _ := range winnings {
			if winnings[i] == 20 {
				possible := possibleSelections(i)

				//fmt.Println(possible)

				for _, value := range possible {
					if board[value] == 0 {
						executeMove(value, player, board)
						return value
					}
				}

				// select the third to win
			}

			if winnings[i] == 2 {
				possible := possibleSelections(i)

				for _, value := range possible {
					if board[value] == 0 {
						executeMove(value, player, board)
						return value
					}
				}

				// select the third to prevent a loss
			}
		}
	}

	if player == 1 {
		canWeWin := false

		winnings := winningPossibilities(board)
		for i, _ := range winnings {
			if winnings[i] == 2 {

				fmt.Println("We have a possibillity to win")

				// select the third to win
				possible := possibleSelections(i)

				//fmt.Println(possible)

				for _, value := range possible {
					if board[value] == 0 {
						canWeWin = true
						executeMove(value, player, board)
						return value
					}
				}
			}
			if canWeWin == false && winnings[i] == 20 {

				// select the third to prevent a loss
				possible := possibleSelections(i)

				for _, value := range possible {
					if board[value] == 0 {
						executeMove(value, player, board)
						return value
					}
				}
			}
		}
	}

	for i, _ := range board {
		if board[i] == 0 {
			executeMove(i, player, board)
			return i
		}
	}

	return 9
}

func possibleSelections(iteration int) [3]int {

	if iteration < 3 {
		return [3]int{iteration, iteration + 3, iteration + 6}
	}
	if iteration == 3 {
		return [3]int{0, 1, 2}
	}
	if iteration == 4 {
		return [3]int{3, 4, 5}
	}
	if iteration == 5 {
		return [3]int{6, 7, 8}
	}
	if iteration == 6 {
		return [3]int{0, 4, 8}
	}
	if iteration == 7 {
		return [3]int{2, 4, 6}
	}

	return [3]int{0, 0, 0}
}