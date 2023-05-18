package main

import (
	"fmt"
	"net/http"

	"github.com/bsm/redislock"
	"github.com/dev-mockingbird/logf"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	grpch "udious.com/mockingbird/channel/grpc"
	pb "udious.com/mockingbird/channel/grpc/proto"
	http1 "udious.com/mockingbird/channel/http"
	"udious.com/mockingbird/channel/pkg/apps"
	"udious.com/mockingbird/channel/pkg/config"
	grpc2 "udious.com/mockingbird/channel/pkg/grpc"
	http2 "udious.com/mockingbird/channel/pkg/http"
	"udious.com/mockingbird/channel/service"
)

func main() {
	app := apps.App{
		Name:    "channel",
		Version: "v0.0.1",
	}
	logger := logf.New()
	if err := config.ReadInput(&app.Cfg, logger); err != nil {
		panic(err)
	}
	e := gin.Default()
	db, err := app.DB()
	if err != nil {
		panic(err)
	}
	redis, err := app.Redis()
	if err != nil {
		panic(err)
	}
	if !app.Cfg.Http.Disable {
		http1.RegisterGin(e, logger, &service.Service{
			DB:         db,
			Logger:     logger,
			LockClient: redislock.New(redis),
		})
		app.With(&http2.StartStopper{Server: http.Server{
			Addr:    fmt.Sprintf(":%d", app.Cfg.Http.Port),
			Handler: e,
		}})
	}
	if !app.Cfg.Grpc.Disable {
		s := grpc.NewServer()
		pb.RegisterGreeterServer(s, grpch.Handler{})
		app.With(&grpc2.StartStopper{Server: s, Port: 8000})
	}
	if err := <-app.Start(); err != nil {
		panic(err)
	}
}
