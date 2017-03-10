package pawn

import (
	"encoding/json"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func MakeHandler(ctx context.Context, ps Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	pawnApplyHandler := kithttp.NewServer(
		ctx,
		makePawnApplyEndpoint(ps),
		decodePawnApplyRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/pawn/v2/apply", pawnApplyHandler).Methods("POST")
	return r
}

func decodePawnApplyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		FundNeeded int `json:"fund_needed"`
		Duration   int `json:"duration"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return pawnApplyRequest{
		FundNeeded: body.FundNeeded,
		Duration:   body.Duration,
	}, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
