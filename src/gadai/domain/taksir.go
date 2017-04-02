package gadai

import "time"

//TaksirResponse response yang diberikan untuk taksiran
type TaksirResponse struct {
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
	ErrorMsg         string
}

//TaksirRequest struct request untuk alat elektronik
type TaksirRequest struct {
	BarangID         string
	Merk             string
	Tipe             string
	Warna            string
	TahunPembelian   string
	KapasitasMemori  string
	KapasitasHardisk string
	OperatingSistem  string
	KelengkapanLain  string
	HargaBeli        int
}

//Taksir menghasilkan data untuk menyimpan data taksir
func Taksir(request interface{}) TaksirResponse {
	t := request.(*TaksirRequest)
	taksirID := TaksirID()
	tahunPembelian, err := ValidasiTahunPembelian(t.TahunPembelian, 5)
	if err != "" {
		return TaksirResponse{ErrorMsg: err}
	}
	return TaksirResponse{
		TaksirID:        taksirID,
		BarangID:        t.BarangID,
		Merk:            t.Merk,
		Tipe:            t.Tipe,
		Warna:           t.Warna,
		KapasitasMemori: t.KapasitasMemori,
		TahunPembelian:  tahunPembelian,
		HargaBeli:       t.HargaBeli,
		ErrorMsg:        "",
	}
}

//ValidasiTahunPembelian memvalidasi apakah tahun pembelian masih diterima
//return date, error
//string adalah error untuk mengisi error message
func ValidasiTahunPembelian(tahunPembelian string, tahunEkonomis int) (time.Time, string) {
	layout := "2006-01-02"
	date, _ := time.Parse(layout, tahunPembelian)
	now := time.Now()

	diffYears := DateDiffYear(now, date)
	if diffYears > tahunEkonomis {
		return date, "Barang anda sudah melewati batas tahun ekonomis"
	}

	return date, ""
}
