package gadai

type GadaiService interface {
	Ajukan(fundNeeded, duration float64) float64
}

type gadaiService struct{}

func NewService() GadaiService {
	return &gadaiService{}
}

func (gadaiService) Ajukan(fundNeeded, duration float64) float64 {
	return fundNeeded * (0.7 / 100) * duration
}
