package models

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	cross       = "x"
	zero        = "o"
	emptySymbol = "*"
)

type Board struct {
	BoardPositions [][]string
	LatestX        int
	LatestY        int
}

func initBoard() *Board {

	return &Board{BoardPositions: [][]string{{emptySymbol, emptySymbol, emptySymbol}, {emptySymbol, emptySymbol, emptySymbol},
		{emptySymbol, emptySymbol, emptySymbol}}, LatestX: -1, LatestY: -1}

}

func (b *Board) UpdateBoard(symbol string, xcoord int, ycoord int) error {

	if b == nil || b.BoardPositions == nil {
		return errors.New("can't fill nil board")
	}

	if xcoord > len(b.BoardPositions) {
		return errors.New("max x value is greater than 2")
	}

	if ycoord > len((b.BoardPositions)[0]) {
		return errors.New("max y value is greater than 2")
	}

	if symbol != cross && symbol != zero {
		return errors.New("invalid symbol")
	}

	b.BoardPositions[xcoord][ycoord] = symbol
	b.LatestX = xcoord
	b.LatestY = ycoord
	return nil
}

func (b *Board) IsFull() bool {

	if b == nil || b.BoardPositions == nil {
		log.Fatal("can't check status of nil board")
	}
	for _, row := range b.BoardPositions {
		for _, symbol := range row {
			if symbol == emptySymbol {
				return false
			}
		}
	}
	return true

}

func (b *Board) HasResult() bool {

	if b == nil || b.BoardPositions == nil {
		log.Fatal("can't check result for nil board")
	}
	symbol := b.BoardPositions[b.LatestX][b.LatestY]
	return b.isRowFull(symbol, b.LatestX) || b.isColumnFull(symbol, b.LatestY) || b.areDiagonalsFull(symbol, b.LatestX, b.LatestY)
}

func (b *Board) isRowFull(symbol string, x int) bool {

	for _, currSymbol := range b.BoardPositions[x] {
		if currSymbol != symbol {
			return false
		}
	}

	return true

}

func (b *Board) isColumnFull(symbol string, y int) bool {

	for row := 0; row < len(b.BoardPositions); row++ {

		if b.BoardPositions[row][y] != symbol {
			return false
		}
	}
	return true

}

func (b *Board) areDiagonalsFull(symbol string, x, y int) bool {

	//case when we x,y correspond to center

	if x == 1 && y == 1 {
		return (b.BoardPositions[0][0] == symbol && b.BoardPositions[2][2] == symbol) ||
			(b.BoardPositions[2][0] == symbol && b.BoardPositions[0][2] == symbol)
	}

	if b.BoardPositions[1][1] != symbol {
		return false
	}
	if x == 0 && y == 0 {
		return b.BoardPositions[2][2] == symbol
	}

	if x == 0 && y == 2 {
		return b.BoardPositions[2][0] == symbol
	}

	if x == 2 && y == 0 {
		return b.BoardPositions[0][2] == symbol
	}

	if x == 2 && y == 2 {
		return b.BoardPositions[0][0] == symbol
	}

	return false

}

func (b *Board) IsCellInBounds(x, y int) bool {

	return 0 <= x && x <= len(b.BoardPositions)-1 && 0 <= y && y <= len(b.BoardPositions[0])-1
}
func (b *Board) IsCellEmpty(x, y int) bool {

	return b.BoardPositions[x][y] == emptySymbol
}

func (b *Board) IsMoveValid(x int, y int) bool {

	return b.IsCellInBounds(x, y) && b.IsCellEmpty(x, y)
}

func (b *Board) Print() {
	if b == nil || b.BoardPositions == nil {
		log.Fatal("Cannot show nil board")
	}
	for i := 0; i < len(b.BoardPositions); i++ {
		fmt.Println(b.BoardPositions[i])
	}
}
