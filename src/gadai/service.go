package gadai

import (
	gadai "pawn/src/gadai/domain"
	gadais "pawn/src/gadai/repository"
	"strconv"

	"github.com/jinzhu/gorm"
)

type GadaiService interface {
	Ajukan(fundNeeded, duration float64) float64
	Taksir(request taksiranGadaiRequest) string
}

type gadaiService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) GadaiService {
	return &gadaiService{
		db: db,
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

	repo := gadais.Taksir{
		TaksirID:         v.TaksirID,
		BarangID:         v.BarangID,
		Merk:             v.Merk,
		Tipe:             v.Tipe,
		Warna:            v.Warna,
		TahunPembelian:   v.TahunPembelian,
		KapasitasMemori:  v.KapasitasMemori,
		KapasitasHardisk: v.KapasitasHardisk,
		OperatingSistem:  v.OperatingSistem,
		KelengkapanLain:  v.KelengkapanLain,
		HargaBeli:        v.HargaBeli,
	}

	gs.db.Create(&repo)
	return v.TaksirID
}
