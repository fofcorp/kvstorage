package httpserv

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Options for http web server.
type Options struct {
	Port string
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
	log.Printf("[httpserv] Start listening on %s\n", srv.instance.Addr)
	err := srv.instance.ListenAndServe()
	return err
}
