package gadai

import "testing"

type taksiranElektronik struct {
	Taksir TaksirElektronik
}

func TestNasabahMelakukanTaksirElektronikOnline(t *testing.T) {
	var taksirCases = []taksiranElektronik{
		{TaksirElektronik{ //Handphone
			BarangID:        "236ae212-0e23-41cf-821f-e811c49a9d21",
			Merk:            "Apple",
			Tipe:            "iPhone 7 plus 256GB",
			Warna:           "Hitam",
			KapasitasMemori: "256",
			TahunPembelian:  "2016-01-01",
			HargaBeli:       15000000,
		}},
		{TaksirElektronik{ //Handphone transaksi ditolak
			BarangID:        "236ae212-0e23-41cf-821f-e811c49a9d21",
			Merk:            "Apple",
			Tipe:            "iPhone 7 plus 256GB",
			Warna:           "Hitam",
			KapasitasMemori: "256",
			TahunPembelian:  "2010-01-01",
			HargaBeli:       15000000,
		}},
	}

	for _, tt := range taksirCases {
		var gadai Gadai
		gadai = &tt.Taksir
		taksir := gadai.Taksir()
		taksirID := taksir.TaksirID
		errorMsg := taksir.ErrorMsg

		if taksirID == "" {
			if errorMsg == "" {
				t.Errorf("Taksir ID tidak ada, seharusnya ada pesan errornya")
			}
		}

		if taksirID != "" && len(taksirID) != 36 {
			t.Errorf("Taksir ID adalah uuid seharusnya jumlahnya 36, yang sekarang %d", len(taksirID))
		}
	}
}
