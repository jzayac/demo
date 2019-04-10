package endpcache

import (
	"time"

	"demo/cache"
	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	next   Service
}

func (ls loggingService) GetList(name string) (output *cache.Iata, err error) {
	defer func(begin time.Time) {
		_ = ls.logger.Log(
			"method", "GetList",
			"cache_name", name,
			"output_length", len(*output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return ls.next.GetList(name)
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}
