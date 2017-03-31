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
