package setup

import (
	"fmt"
	"github.com/pseidemann/finish"
	"log"
	"net/http"
	"time"
)



// Server represents the basic struct
// to implement a listening server that supports graceful shutdown
type Server struct {
	Router       http.Handler
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	StopTimeout  time.Duration
	finisher     *finish.Finisher
}

// Run starts the server with a finish handler
func (s *Server) Run() error {
	srv := &http.Server{
		Handler:      s.Router,
		Addr:         fmt.Sprintf(":%d", s.Port),
		WriteTimeout: s.WriteTimeout,
		ReadTimeout:  s.ReadTimeout,
	}

	s.finisher = &finish.Finisher{
		Timeout: s.StopTimeout,
	}

	s.finisher.Add(srv)

	errs := make(chan error, 1)

	go func() {
		log.Printf("Server started on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errs <- err
			s.finisher.Trigger()
		}
		close(errs)
	}()

	s.finisher.Wait()

	return <-errs
}

// New creates a new server with options
func New(router http.Handler) (*Server, error) {
	s := &Server{Router: router}
	s.Port = 8000
	return s, nil
}
