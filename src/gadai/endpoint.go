package gadai

import (
	"context"

	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func makeAjukanGadaiEndpoint(svc GadaiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ajukanGadaiRequest)

		fundNeededFloat, _ := strconv.ParseFloat(req.fundNeeded, 64)
		durationFloat, _ := strconv.ParseFloat(req.duration, 64)

		v := svc.Ajukan(fundNeededFloat, durationFloat)

		return ajukanGadaiResponse{v}, nil
	}
}

type ajukanGadaiRequest struct {
	fundNeeded string `json:"fund_needed"`
	duration   string `json:"duration"`
}

type ajukanGadaiResponse struct {
	AdminCost float64 `json:"admin_cost"`
}

func makeTaksirGadaiEndPoint(svc GadaiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(taksiranGadaiRequest)
		v := svc.Taksir(req)
		return taksiranGadaiResponse{v}, nil
	}
}

type taksiranGadaiRequest struct {
	BarangID         string `json:"barang_id"`
	Merk             string `json:"merk"`
	Tipe             string `json:"tipe"`
	Warna            string `json:"warna"`
	TahunPembelian   string `json:"tahun_pembelian"`
	KapasitasMemori  string `json:"kapasitas_memori"`
	KapasitasHardisk string `json:"kapasitas_hardisk"`
	OperatingSistem  string `json:"operating_sistem"`
	KelengkapanLain  string `json:"kelengkapan_lain"`
	HargaBeli        string `json:"harga_beli"`
}

type taksiranGadaiResponse struct {
	TaksirID string `json:"taksir_id"`
}
