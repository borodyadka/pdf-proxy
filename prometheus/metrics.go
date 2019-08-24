package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

var HTTPRequestsTotal prometheus.Counter
var HTTPRequestProcessingTimeSummaryMs prometheus.Summary
var HTTPRequestStatuses *prometheus.CounterVec

func init() {
	HTTPRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of handled requests",
		})
	prometheus.MustRegister(HTTPRequestsTotal)

	HTTPRequestProcessingTimeSummaryMs = prometheus.NewSummary(
		prometheus.SummaryOpts{
			Name:       "http_request_processing_time_summary_ms",
			Help:       "Average time per request",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		})
	prometheus.MustRegister(HTTPRequestProcessingTimeSummaryMs)

	HTTPRequestStatuses = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_request_statuses",
		Help: "Total number of requests by statuses",
	}, []string{"code"})
	prometheus.MustRegister(HTTPRequestStatuses)
}
