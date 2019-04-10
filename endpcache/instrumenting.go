package endpcache

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"

	"demo/cache"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           Service
}

func (mw instrumentingService) GetList(name string) (output *cache.Iata, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetList", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetList(name)
	return
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           s,
	}
}
