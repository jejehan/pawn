package gadai

import (
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	GadaiService
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s GadaiService) GadaiService {
	return &loggingService{logger, s}
}

func (s *loggingService) Taksir(req taksiranGadaiRequest) string {
	reqJson, _ := json.Marshal(req)
	defer func(begin time.Time) {
		s.logger.Log("method", "taksir", "req", string(reqJson), "took", time.Since(begin))
	}(time.Now())
	return s.GadaiService.Taksir(req)
}
