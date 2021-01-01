package models

const humanPlayerType = "human"

type HumanPlayer struct {
	Name   string
	ID     int
	Symbol string
}

func (p *HumanPlayer) Init(id int) {
	return
}
func (p *HumanPlayer) GetName() string {
	return p.Name
}
func (p *HumanPlayer) GetID() int {
	return p.ID
}

func (p *HumanPlayer) GetSymbol() string {
	return p.Symbol
}

func (p *HumanPlayer) GetOpponentSymbol() string {

	if p.Symbol == cross {
		return zero
	}
	return cross
}

func (p *HumanPlayer) GetNextMove(*Board) (int, int, bool) {
	return 0, 0, false
}
