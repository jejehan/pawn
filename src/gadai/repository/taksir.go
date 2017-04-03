package gadai

import (
	gadai "pawn/src/gadai/domain"
	"time"

	"github.com/jinzhu/gorm"
)

type Repository interface {
	Save(req gadai.TaksirResponse)
}

type repository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type Taksir struct {
	gorm.Model
	TaksirID         string
	BarangID         string
	Merk             string
	Tipe             string
	Warna            string
	TahunPembelian   time.Time
	KapasitasMemori  string
	KapasitasHardisk string
	OperatingSistem  string
	KelengkapanLain  string
	HargaBeli        int
	HargaTaksirAtas  int
	HargaTaksirBawah int
}

func (r *repository) Save(req gadai.TaksirResponse) {
	repo := &Taksir{
		TaksirID:         req.TaksirID,
		BarangID:         req.BarangID,
		Merk:             req.Merk,
		Tipe:             req.Tipe,
		Warna:            req.Warna,
		TahunPembelian:   req.TahunPembelian,
		KapasitasMemori:  req.KapasitasMemori,
		KapasitasHardisk: req.KapasitasHardisk,
		OperatingSistem:  req.OperatingSistem,
		KelengkapanLain:  req.KelengkapanLain,
		HargaBeli:        req.HargaBeli,
	}

	r.db.Create(repo)
}
