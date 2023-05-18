package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"
	"time"
)

type handler struct{}

func (handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("hello world"))
}

func TestGrpc_Start(t *testing.T) {
	s := http.Server{Addr: ":8001", Handler: handler{}}
	rs := StartStopper{Server: s}
	errCh := rs.Start()
	time.Sleep(time.Second)

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8001", nil)
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if !(resp.StatusCode == 200 && bytes.Equal(body, []byte("hello world"))) {
		t.Fatal("received false message")
	}
	if err := rs.Stop(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := <-errCh; err != nil {
		t.Fatal(err)
	}
}
