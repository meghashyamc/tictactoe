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

func ShowBoard(b *models.Board) {
	if b == nil || b.BoardPositions == nil {
		log.Fatal("Cannot show nil board")
	}
	fmt.Println("This is how the board now looks like:")
	for i := 0; i < len(b.BoardPositions); i++ {
		fmt.Println(b.BoardPositions[i])
	}

}

func ShowGameEnd(boardFull bool, hasResult bool, b *models.Board, p models.Player) {

	ShowBoard(b)
	fmt.Println("Game over")

	if boardFull {
		fmt.Println("The game ended in a tie")
	} else if hasResult {
		fmt.Println("The winner is", p.GetName()+". Congrats,"+p.GetName()+"!")
	}
}

func ShowPlayerMove(p models.Player, x, y int) {

	fmt.Println(p.GetName()+" just played their move:", x, y)
}

func ShowInvalidMove() {
	fmt.Println("Please enter valid coordinates. The coordinates must be from 0 to 3. Also, already occupied coordinates should not be entered.")

}
