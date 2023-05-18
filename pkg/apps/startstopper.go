package apps

import (
	"context"
	"sync"

	"github.com/dev-mockingbird/errors"
)

type Starter interface {
	Start() chan error
}

type Stopper interface {
	Stop(context.Context) error
}

type StartStopper interface {
	Starter
	Stopper
}

type StartStoppers []StartStopper

func (set StartStoppers) Start() chan error {
	errCh := make(chan error, len(set))
	for _, s := range set {
		go func(s Starter) {
			errch := s.Start()
			err := <-errch
			errCh <- err
		}(s)
	}
	return errCh
}

func (set StartStoppers) Stop(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(set))
	errs := errors.MultiError()
	for _, s := range set {
		go func(s Stopper) {
			defer wg.Done()
			if err := s.Stop(ctx); err != nil {
				errs.Join(err)
			}
		}(s)
	}
	wg.Wait()
	return errs
}
