package models

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

const (
	cross       = "x"
	zero        = "o"
	emptySymbol = "*"
)

type Board [][]string

func initBoard() Board {

	return Board{[]string{emptySymbol, emptySymbol, emptySymbol},
		[]string{emptySymbol, emptySymbol, emptySymbol},
		[]string{emptySymbol, emptySymbol, emptySymbol}}

}

func (b Board) UpdateBoard(symbol string, xcoord int, ycoord int) error {

	if b == nil {
		return errors.New("can't fill nil board")
	}

	if xcoord > len(b) {
		return errors.New("max x value is greater than 2")
	}

	if ycoord > len((b)[0]) {
		return errors.New("max y value is greater than 2")
	}

	if symbol != cross && symbol != zero {
		return errors.New("invalid symbol")
	}

	b[xcoord][ycoord] = symbol
	return nil
}

func (b Board) IsFull() bool {

	if b == nil {
		log.Fatal("can't check status of nil board")
	}
	for _, row := range b {
		for _, symbol := range row {
			if symbol == emptySymbol {
				return false
			}
		}
	}
	return true

}

func (b Board) HasResult(symbol string, x int, y int) bool {

	if b == nil {
		log.Fatal("can't check result for nil board")
	}
	return b.isRowFull(symbol, x) || b.isColumnFull(symbol, y) || b.areDiagonalsFull(symbol, x, y)
}

func (b Board) isRowFull(symbol string, x int) bool {

	for _, currSymbol := range b[x] {
		if currSymbol != symbol {
			return false
		}
	}

	return true

}

func (b Board) isColumnFull(symbol string, y int) bool {

	for row := 0; row < len(b); row++ {

		if b[row][y] != symbol {
			return false
		}
	}
	return true

}

func (b Board) areDiagonalsFull(symbol string, x, y int) bool {

	//case when we x,y correspond to center

	if x == 1 && y == 1 {
		return (b[0][0] == symbol && b[2][2] == symbol) ||
			(b[2][0] == symbol && b[0][2] == symbol)
	}

	if b[1][1] != symbol {
		return false
	}
	if x == 0 && y == 0 {
		return b[2][2] == symbol
	}

	if x == 0 && y == 2 {
		return b[2][0] == symbol
	}

	if x == 2 && y == 0 {
		return b[0][2] == symbol
	}

	if x == 2 && y == 2 {
		return b[0][0] == symbol
	}

	return false

}

func (b Board) IsCellInBounds(x, y int) bool {

	return 0 <= x && x <= len(b)-1 && 0 <= y && y <= len(b[0])-1
}
func (b Board) IsCellEmpty(x, y int) bool {

	return b[x][y] == emptySymbol
}

func (b Board) IsMoveValid(x int, y int) bool {

	return b.IsCellInBounds(x, y) && b.IsCellEmpty(x, y)
}
