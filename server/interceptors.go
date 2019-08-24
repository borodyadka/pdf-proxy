package server

import (
	"github.com/borodyadka/pdf-proxy/prometheus"
	"net/http"
	"strconv"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	return w.ResponseWriter.Write(b)
}

func collectMetrics(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prometheus.HTTPRequestsTotal.Inc()
		sw := &statusWriter{ResponseWriter: w}
		start := time.Now()
		next.ServeHTTP(sw, r)
		prometheus.HTTPRequestProcessingTimeSummaryMs.Observe(
			float64(time.Since(start).Nanoseconds() / int64(time.Millisecond)))
		prometheus.HTTPRequestStatuses.WithLabelValues(strconv.Itoa(sw.status)).Inc()
	})
}
