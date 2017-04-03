package gadai

import (
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	GadaiService
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s GadaiService) GadaiService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		GadaiService:   s,
	}
}

func (s *instrumentingService) Taksir(req taksiranGadaiRequest) string {
	defer func(begin time.Time) {
		s.requestCount.With("method", "taksir").Add(1)
		s.requestLatency.With("method", "taksir").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.GadaiService.Taksir(req)
}
