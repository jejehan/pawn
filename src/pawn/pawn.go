package pawn

import "github.com/pborman/uuid"

type Pawn interface {
	Apply()
	PawnRate()
}

type PawnID string

type pawn struct {
	PawnID   PawnID
	MemberID string
	PawnRate int
	Amount   int
}

func (p *pawn) Apply() {
	pawnID := p.PawnID
}

func (p *pawn) PawnRates() int {
	return (p.Amount * p.PawnRate) / 100
}

func (p *pawn) NextPawnID() PawnID {
	return PawnID(uuid.New())
}
