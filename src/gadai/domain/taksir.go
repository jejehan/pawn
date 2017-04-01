package gadai

type TaksirResponse struct {
	TaksirID string
	ErrorMsg string
}

type TaksirElektronik struct {
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

func (t *TaksirElektronik) Taksir() TaksirResponse {
	return TaksirResponse{
		TaksirID: "236ae212-0e23-41cf-821f-e811c49a9d21",
		ErrorMsg: "",
	}
}
