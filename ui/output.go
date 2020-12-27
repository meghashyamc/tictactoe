package ui

import (
	"fmt"

	models "github.com/meghashyamc/tictactoe/models"
	log "github.com/sirupsen/logrus"
)

func ShowWelcome() {
	fmt.Println("Welcome to Tic Tac Toe")

}
func ShowPlayer(p models.Player) {

	if p == nil {
		log.Fatal("can't render nil player")
	}
	fmt.Println("Player number", p.GetID(), ":", p.GetName())

}

func ShowBoard(b models.Board) {

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

}

func ShowGameEnd(boardFull bool, hasResult bool, b models.Board, p models.Player) {

	ShowBoard(b)
	fmt.Println("Game over")

	if boardFull {
		fmt.Println("The game ended in a tie")
	} else if hasResult {
		fmt.Println("The winner is", p.GetName()+". Congrats,"+p.GetName()+"!")
	}
}

func ShowInvalidMove() {
	fmt.Println("Please enter valid coordinates. The coordinates must be from 0 to 3. Also, already occupied coordinates should not be entered.")

}
