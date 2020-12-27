package models

const aiPlayerType = "ai"

type AIPlayer struct {
	Name   string
	ID     int
	Symbol string
}

func (p *AIPlayer) Init() {

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

func (p *AIPlayer) GetNextMove() (int, int, bool) {
	return 0, 0, false
}
