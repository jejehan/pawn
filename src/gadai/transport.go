package gadai

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func MakeHandler(ctx context.Context, gd GadaiService) http.Handler {
	r := mux.NewRouter()

	taksiranGadaiHandler := kithttp.NewServer(
		ctx,
		makeTaksirGadaiEndPoint(gd),
		decodeTaksirGadaiRequest,
		encodeResponse,
	)
	ajukanGadaiHandler := kithttp.NewServer(
		ctx,
		makeAjukanGadaiEndpoint(gd),
		decodeAjukanGadaiRequest,
		encodeResponse,
	)
	r.Handle("/gadai/v2/taksiran", taksiranGadaiHandler).Methods("POST")
	r.Handle("/gadai/v2/ajukan", ajukanGadaiHandler).Methods("POST")

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeAjukanGadaiRequest(_ context.Context, r *http.Request) (interface{}, error) {

	fundNeeded := r.FormValue("fund_needed")
	duration := r.FormValue("duration")

	return ajukanGadaiRequest{
		fundNeeded: fundNeeded,
		duration:   duration,
	}, nil
}

func decodeTaksirGadaiRequest(_ context.Context, r *http.Request) (interface{}, error) {

	barangID := r.FormValue("barang_id")
	merk := r.FormValue("merk")
	tipe := r.FormValue("tipe")
	kapasitasMemori := r.FormValue("kapasitas_memori")
	tahunPembelian := r.FormValue("tahun_pembelian")
	hargaBeli := r.FormValue("harga_beli")

	return taksiranGadaiRequest{
		BarangID:        barangID,
		Merk:            merk,
		Tipe:            tipe,
		KapasitasMemori: kapasitasMemori,
		TahunPembelian:  tahunPembelian,
		HargaBeli:       hargaBeli,
	}, nil
}
