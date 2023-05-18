package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type StartStopper struct {
	Server             *grpc.Server
	ReflectionDisabled bool
	Listener           net.Listener
	Port               int
}

func (s *StartStopper) Start() chan error {
	errCh := make(chan error, 1)
	if s.Server == nil {
		s.Server = grpc.NewServer()
	}
	if !s.ReflectionDisabled {
		reflection.Register(s.Server)
	}
	if s.Listener == nil {
		var err error
		s.Listener, err = net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
		if err != nil {
			errCh <- err
			return errCh
		}
	}
	go func() {
		err := s.Server.Serve(s.Listener)
		errCh <- err
	}()
	return errCh
}

func (s *StartStopper) Stop(_ context.Context) error {
	s.Server.Stop()
	return nil
}
