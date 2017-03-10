package pawn

type Repository interface {
	Create(pawn)
	Save()
	Delete()
}
