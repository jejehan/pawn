package gadai

import "testing"

//Test Nasabah Melakukan Taksir Elektronik Online
func TestNasabahMelakukanTaksirElektronikOnline(t *testing.T) {

	//var gadai Gadai
	request := TaksirRequest{
		BarangID:        "236ae212-0e23-41cf-821f-e811c49a9d21",
		Merk:            "Apple",
		Tipe:            "iPhone 7 plus",
		Warna:           "Hitam",
		KapasitasMemori: "256",
		TahunPembelian:  "2016-01-01",
		HargaBeli:       15000000,
	}
	//taksir := gadai.Taksir(&request)
	taksir := Taksir(&request)
	taksirID := taksir.TaksirID
	errorMsg := taksir.ErrorMsg

	if taksirID == "" {
		if errorMsg == "" {
			t.Errorf("Taksir ID tidak ada, seharusnya ada pesan errornya")
		}
	}

	lenOfUUID := 36
	if taksirID != "" && len(taksirID) != lenOfUUID {
		t.Errorf("Taksir ID adalah uuid seharusnya jumlahnya %d, yang sekarang %d", lenOfUUID, len(taksirID))
	}

}
