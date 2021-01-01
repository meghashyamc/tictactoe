package models

type Player interface {
	Init(id int)
	GetName() string
	GetID() int
	GetSymbol() string
	GetOpponentSymbol() string
	GetNextMove(*Board) (int, int, bool)
}
