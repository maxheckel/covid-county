package setup

import "time"

// Option is a functional option that permits mutating
// a server's port, timeouts and handler
type Option interface {
	Apply(*Server) error
}

// OptionFunc is a function implementing the Option interface
type OptionFunc func(*Server) error

// Apply will take the wrapped function and run it against the server
func (o OptionFunc) Apply(s *Server) error {
	return o(s)
}

// Port sets the server port to listen on
func Port(port int) Option {
	return OptionFunc(func(s *Server) error {
		s.Port = port
		return nil
	})
}

// Timeouts sets the server timeouts for read and write
func Timeouts(read time.Duration, write time.Duration) Option {
	return OptionFunc(func(s *Server) error {
		s.ReadTimeout = read
		s.WriteTimeout = write
		return nil
	})
}

// StopTimeout sets the server timeout for graceful shutdown
func StopTimeout(stop time.Duration) Option {
	return OptionFunc(func(s *Server) error {
		s.StopTimeout = stop
		return nil
	})
}

