package server

import (
	"context"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	http *http.Server
}

func (s *Server) pingRoute(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("PONG\n"))
}

func (s *Server) renderRoute(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		_ = response(w, http.StatusBadRequest, "parameter `url` must be set")
		return
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		_ = response(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		log.WithField("error", err).Error("PDF generator error")
		return
	}

	page := wkhtmltopdf.NewPage(url)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		_ = response(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		log.WithField("error", err).Error("PDF creation error")
		return
	}
	// TODO: edit exif metadata to change author, producer and creator fields
	// see: https://github.com/dsoprea/go-exif

	log.WithField("url", url).Debug("rendering URL")
	headers := w.Header()
	headers["Content-Type"] = []string{"application/pdf"}
	_, _ = w.Write(pdfg.Bytes())
}

func (s *Server) ListenAndServe() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", s.pingRoute)
	mux.Handle("/render", collectMetrics(s.renderRoute))
	mux.Handle("/metrics", promhttp.Handler())

	s.http.Handler = mux
	return s.http.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}

func New(address string) *Server {
	return &Server{
		http: &http.Server{Addr: address},
	}
}
