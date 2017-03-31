package pawn

type Service interface {
	Apply(amount, pawnRate int) (PawnID, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}
