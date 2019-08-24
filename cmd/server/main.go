package main

import (
	"context"
	"fmt"
	"github.com/borodyadka/pdf-proxy/config"
	"github.com/borodyadka/pdf-proxy/server"
	"time"

	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		DisableSorting: false,
	})
	log.SetOutput(os.Stdout)
	logLevel, _ := log.ParseLevel(config.LogLevel)
	log.SetLevel(logLevel)

	srv := server.New(config.Address)

	go func() {
		log.WithFields(log.Fields{
			"transport": "http",
			"address":   config.Address,
		}).Info("server listening")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("unable to start http server: %s", err)
		}
	}()

	sigChan := make(chan error, 2)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
		sigChan <- fmt.Errorf("received signal %s", <-c)

		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()
		// TODO: wait for requests is finished
		if err := srv.Stop(ctx); err != nil {
			log.Fatalf("unable to stop http server: %s", err)
		}
	}()

	log.Info("terminated ", <-sigChan)
}
