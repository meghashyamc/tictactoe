package models

type Player interface {
	Init()
	GetName() string
	GetID() int
	GetSymbol() string
	GetNextMove() (int, int, bool)
}
