package main

import (
	"github.com/meghashyamc/tictactoe/core"
	"github.com/meghashyamc/tictactoe/ui"
)

func main() {

	ui.ShowWelcome()
	gameType := ui.GetGameType()
	newGame := core.InitializeNewGame(gameType)
	core.Play(newGame)

}
