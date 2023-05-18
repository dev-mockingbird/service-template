package http

import (
	"context"
	"net/http"

	"udious.com/mockingbird/channel/pkg/apps"
)

type StartStopper struct {
	Server http.Server
}

var _ apps.StartStopper = StartStopper{}

func (h StartStopper) Start() chan error {
	errCh := make(chan error, 1)
	go func() {
		err := h.Server.ListenAndServe()
		errCh <- err
	}()
	return errCh
}

func (h StartStopper) Stop(ctx context.Context) error {
	return h.Server.Shutdown(ctx)
}
