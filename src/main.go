package main

import (
	"fmt"
)

func main() {
	var game Game
	game.player = "O"

	gameOver := false
	var winner string

	for gameOver != true {
		PrintBoard(game.board)
		move := AskForPlay()
		err := game.CheckMove(move)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gameOver, winner = CheckForWinner(game.board, game.turnNumber)
	}
	PrintBoard(game.board)
	PrintVictory(winner)
}
