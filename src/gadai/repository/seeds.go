package gadai

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type SeedService interface {
	SeedAll()
	SeedsTaksir()
}

type seedService struct {
	db *gorm.DB
}

func Config(db *gorm.DB) SeedService {
	return &seedService{
		db: db,
	}
}

func (c *seedService) SeedAll() {
	c.SeedsTaksir()
}

func (c *seedService) SeedsTaksir() {
	tahunPembelian := time.Now()
	c.db.Create(&Taksir{
		TaksirID:        TaksirID(),
		BarangID:        "236ae212-0e23-41cf-821f-e811c49a9d21",
		Merk:            "Apple",
		Tipe:            "iPhone 7 Plus 256GB",
		Warna:           "Hitam",
		TahunPembelian:  tahunPembelian,
		KapasitasMemori: "256",
		KelengkapanLain: "Charger",
		HargaBeli:       10000000,
	})
}

func TaksirID() string {
	u := uuid.NewV4()
	return u.String()
}
