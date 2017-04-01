package gadai

import (
	"pawn/src/gadai/domain"
)

type GadaiService interface {
	Ajukan(fundNeeded, duration float64) float64
	Taksir(request taksiranGadaiRequest) string
}

type gadaiService struct{}

func NewService() GadaiService {
	return &gadaiService{}
}

func (gs *gadaiService) Ajukan(fundNeeded, duration float64) float64 {
	return fundNeeded * (0.7 / 100) * duration
}

func (gs *gadaiService) Taksir(req taksiranGadaiRequest) string {
	var gadai gadai.Gadai
	v := gadai.Taksir()
	return v.TaksirID
}
