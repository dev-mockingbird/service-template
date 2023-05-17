package main

import (
	"fmt"

	"github.com/bsm/redislock"
	"github.com/dev-mockingbird/logf"
	"github.com/gin-gonic/gin"
	"udious.com/mockingbird/channel/http"
	"udious.com/mockingbird/channel/pkg/config"
	"udious.com/mockingbird/channel/pkg/config/apps"
	"udious.com/mockingbird/channel/service"
)

func main() {
	app := apps.App{}
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
	http.RegisterGin(e, logger, &service.Service{
		DB:         db,
		Logger:     logger,
		LockClient: redislock.New(redis),
	})
	if err := e.Run(fmt.Sprintf(":%d", app.Cfg.Http.Port)); err != nil {
		panic(err)
	}
}
