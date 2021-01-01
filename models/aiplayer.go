package models

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
)

const (
	aiPlayerType = "ai"
	aiPlayerName = "Computer"
)

type AIPlayer struct {
	Name   string
	ID     int
	Symbol string
}

func (p *AIPlayer) Init(id int) {

	p.Name = aiPlayerName
	p.ID = id
}
func (p *AIPlayer) GetName() string {
	return p.Name
}
func (p *AIPlayer) GetID() int {
	return p.ID
}

func (p *AIPlayer) GetSymbol() string {
	return p.Symbol
}
func (p *AIPlayer) GetOpponentSymbol() string {
	if p.Symbol == cross {
		return zero
	}
	return cross
}

func (p *AIPlayer) GetNextMove(b *Board) (int, int, bool) {

	if b == nil || b.BoardPositions == nil {
		log.Fatal("Can't get next move for nil board")
	}
	latestX := b.LatestX
	latestY := b.LatestY

	type pos struct {
		x int
		y int
	}

	availablePositions := make([]pos, 0)

	for x, row := range b.BoardPositions {
		for y, symbol := range row {
			if symbol == emptySymbol {
				b.BoardPositions[x][y] = p.GetOpponentSymbol()
				b.LatestX = x
				b.LatestY = y
				if b.HasResult() {
					b.BoardPositions[x][y] = emptySymbol
					b.restore(latestX, latestY)
					return x, y, true
				}
				b.BoardPositions[x][y] = p.GetSymbol()
				if b.HasResult() {
					b.BoardPositions[x][y] = emptySymbol
					b.restore(latestX, latestY)
					return x, y, true
				}
				b.BoardPositions[x][y] = emptySymbol
				b.restore(latestX, latestY)
				availablePositions = append(availablePositions, pos{x, y})
			}
		}
	}

	//reaching here means we can't win in this turn, nor can the opponent win in the next turn
	randomIndex := rand.Intn(len(availablePositions))
	return availablePositions[randomIndex].x, availablePositions[randomIndex].y, true

}

func (b *Board) restore(restoreX, restoreY int) {
	b.LatestX = restoreX
	b.LatestY = restoreY
}
