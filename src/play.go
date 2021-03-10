package main

import "errors"

type Game struct {
	board      [9]string
	player     string
	turnNumber int
}

func (game *Game) SwitchPlayers() {
	if game.player == "O" {
		game.player = "X"
		return
	}
	game.player = "O"
}

func (game *Game) CheckMove(pos int) error {
	if game.board[pos-1] == "" {
		game.board[pos-1] = game.player
		game.SwitchPlayers()
		game.turnNumber += 1
		return nil
	}
	return errors.New("try another move")
}
