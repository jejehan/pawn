package gadai

import (
	gadai "pawn/src/gadai/domain"
	gadais "pawn/src/gadai/repository"
	"strconv"
)

type GadaiService interface {
	Ajukan(fundNeeded, duration float64) float64
	Taksir(request taksiranGadaiRequest) string
}

type gadaiService struct {
	taksirs gadais.Repository
}

func NewService(taksirs gadais.Repository) GadaiService {
	return &gadaiService{
		taksirs: taksirs,
	}
}

func (gs *gadaiService) Ajukan(fundNeeded, duration float64) float64 {
	return fundNeeded * (0.7 / 100) * duration
}

func (gs *gadaiService) Taksir(req taksiranGadaiRequest) string {
	hargaBeli, _ := strconv.Atoi(req.HargaBeli)
	rq := gadai.TaksirRequest{
		BarangID:        req.BarangID,
		Merk:            req.Merk,
		Tipe:            req.Tipe,
		Warna:           req.Warna,
		TahunPembelian:  req.TahunPembelian,
		KapasitasMemori: req.KapasitasMemori,
		HargaBeli:       hargaBeli,
	}
	v := gadai.Taksir(&rq)
	gs.taksirs.Save(v)
	return v.TaksirID
}
