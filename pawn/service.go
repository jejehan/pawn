package pawn

import "pawn/pawn"

type Service interface {
	Apply(memberid, amount, pawnRate) (pawn.PawnID, error)
}

type service struct {
	pawn pawn.Repository
}

func NewService() Service {
	return &service{}
}
