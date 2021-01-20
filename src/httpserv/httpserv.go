package httpserv

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fofcorp/kvstorage/src/storage"
	log "github.com/sirupsen/logrus"
)

// Options for http web server.
type Options struct {
	Port  string
	Store storage.Storage
}

// Server represents http web server.
type Server struct {
	instance *http.Server
	options  *Options
}

// Init constructs server.
func Init(options *Options) (*Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	middlewares := []Middleware{
		AccessLog,
	}

	mux.Handle("/api/v0/get",
		MiddlewareChain(
			http.HandlerFunc(GetHandler(options.Store)),
			middlewares...,
		),
	)
	mux.Handle("/api/v0/put",
		MiddlewareChain(
			http.HandlerFunc(PutHandler(options.Store)),
			middlewares...,
		),
	)
	mux.Handle("/api/v0/delete",
		MiddlewareChain(
			http.HandlerFunc(DeleteHandler(options.Store)),
			middlewares...,
		),
	)
	addr := fmt.Sprintf("localhost:%s", options.Port)
	instance := &http.Server{
		Handler:           mux,
		Addr:              addr,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	return &Server{options: options, instance: instance}, nil
}

// Run http web server.
func (srv *Server) Run() error {
	log.WithFields(log.Fields{
		"module": "httpserv",
		"host":   srv.instance.Addr,
	}).Info("httpserv_start")
	err := srv.instance.ListenAndServe()
	return err
}
