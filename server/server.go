package server

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type server struct{}

func (s *server) Run() error {
	g := new(multierror.Group)
	g.Go(
		func() error {
			// do some work in the background
			// ... work code here ...
			return nil
		},
	)

var e *multierror.Error
if errors.As(g.Wait(), &e) {
	// err is a *multierror.Error, and e is set to the error's value
	errs := e.WrappedErrors()
	...
}

	errs := err.WrappedErrors()

	_ = errs

	return err
}

func main() {
	s := server{}
	if err := s.Run(); err != nil {
		panic("error while running")
	}
}
