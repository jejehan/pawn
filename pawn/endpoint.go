package pawn

import (
	"context"
	"pawn/pawn"

	"github.com/go-kit/kit/endpoint"
)

type pawnApplyRequest struct {
	FundNeeded int
	Duration   int
}

type pawnApplyResponse struct {
	PawnID pawn.PawnID `json:"pawn_id,omitempty"`
	Err    error       `jsos:"error,omitempty"`
}

func (r pawnApplyResponse) error() error { return r.Err }

func makePawnApplyEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(pawnApplyRequest)
		sim := s.Apply(req.FundNeeded, req.Duration)
		return pawnApplyResponse{ID: id, Err: err}, nil
	}
}
