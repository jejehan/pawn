package gadai

import "testing"

func TestNasabahMelakukanTaksirElektronikOnline(t *testing.T) {

	//var gadai Gadai
	request := TaksirRequest{}
	//taksir := gadai.Taksir(&request)
	taksir := Taksir(&request)
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
