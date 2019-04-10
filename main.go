package main

import (
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	ec "demo/endpcache"

	"demo/model"
)

const (
	defaultPort = ":7000"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	if err := model.InitDatabaseModel(); err != nil {
		logger.Log("fatal connection to db: ", err)
		os.Exit(1)
	}

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "cache",
		Subsystem: "cache_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "cache",
		Subsystem: "cache_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	var svc ec.Service
	svc = ec.NewService()
	svc = ec.NewLoggingService(logger, svc)
	svc = ec.NewInstrumentingService(requestCount, requestLatency, svc)

	httpLogger := log.With(logger, "component", "http")
	mux := http.NewServeMux()

	mux.Handle("/v1/", ec.MakeHandler(svc, httpLogger))

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", accessControl(mux))

	logger.Log("msg", "HTTP", "addr", defaultPort)
	logger.Log("err", http.ListenAndServe(defaultPort, nil))
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
