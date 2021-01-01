package models

const maxPlayers = 2
const (
	humanVersusHumanType = "human versus human"
	humanVersusAIType    = "human versus AI"
)

type Game struct {
	Board   *Board
	Type    string
	Players []Player
}

func (game *Game) Init(player1 Player, player2 Player) {

	game.Board = initBoard()
	game.Players = []Player{player1, player2}

}
