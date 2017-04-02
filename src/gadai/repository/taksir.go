package gadai

import (
	"time"

	"github.com/jinzhu/gorm"
)

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
}
