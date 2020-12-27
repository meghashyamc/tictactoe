package core

import (
	models "github.com/meghashyamc/tictactoe/models"
	ui "github.com/meghashyamc/tictactoe/ui"
)

func InitializeNewGame(gameType string) *models.Game {

	game := &models.Game{Type: gameType}

	player1 := &models.HumanPlayer{}
	player1 = ui.GetHumanPlayer(1)
	player1.Symbol = cross

	var player2 models.Player
	if gameType == humanVersusHumanType {
		player := ui.GetHumanPlayer(2)
		player.Symbol = zero
		player2 = player
	} else {
		player := &models.AIPlayer{}
		player.Init()
		player.Symbol = zero
		player2 = player
	}

	game.Init(player1, player2)
	ui.ShowBoard(game.Board)
	return game

}
