package core

import (
	models "github.com/meghashyamc/tictactoe/models"
	ui "github.com/meghashyamc/tictactoe/ui"
)

func Play(game *models.Game) {

	for {
		for _, player := range game.Players {
			x, y, gotMove := player.GetNextMove()
			if !gotMove {
				x, y = ui.GetNextMove(player)
			}
			if game.Board.IsMoveValid(x, y) {
				game.Board.UpdateBoard(player.GetSymbol(), x, y)
				boardFull := game.Board.IsFull()
				hasResult := game.Board.HasResult(player.GetSymbol(), x, y)
				if boardFull || hasResult {
					ui.ShowGameEnd(boardFull, hasResult, game.Board, player)
					return
				}
			} else {
				ui.ShowInvalidMove()
				break
			}

			ui.ShowBoard(game.Board)
		}
	}
}
