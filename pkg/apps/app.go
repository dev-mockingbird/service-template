package apps

import (
	"context"
	"fmt"
	"sync"

	"github.com/dev-mockingbird/errors"
	"github.com/dev-mockingbird/events"
	"github.com/dev-mockingbird/logf"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"udious.com/mockingbird/channel/pkg/config"
)

type App struct {
	Cfg          config.Config
	Name         string
	Version      string
	logger       logf.Logger
	initlogger   sync.Once
	db           *gorm.DB
	dbinit       sync.Once
	dberr        error
	redis        *redis.Client
	rediserr     error
	redisinit    sync.Once
	servers      StartStoppers
	eventbuses   map[string]events.EventBus
	eventbuslock sync.Mutex
}

func (app *App) DB() (*gorm.DB, error) {
	app.dbinit.Do(func() {
		if app.Cfg.DB.Disable {
			app.dberr = errors.New("db is disabled")
			return
		}
		switch app.Cfg.DB.DBMS {
		case config.Mysql:
			app.db, app.dberr = gorm.Open(mysql.Open(app.Cfg.DB.DSN()), &gorm.Config{})
		default:
			app.dberr = fmt.Errorf("not supported dbms [%s]", app.Cfg.DB.DBMS)
		}
	})
	return app.db, app.dberr
}

func (app *App) Logger() logf.Logger {
	app.initlogger.Do(func() {
		app.logger = logf.New()
	})
	return app.logger
}

func (app *App) Redis() (*redis.Client, error) {
	app.redisinit.Do(func() {
		if app.Cfg.Redis.Disable {
			app.rediserr = errors.New("redis is disabled")
			return
		}
		app.redis = redis.NewClient(&redis.Options{
			Addr:     app.Cfg.Redis.Addr,
			Username: app.Cfg.Redis.Username,
			Password: app.Cfg.Redis.Password,
			DB:       app.Cfg.Redis.DB,
		})
	})
	return app.redis, app.rediserr
}

func (app *App) With(s StartStopper) *App {
	app.servers = append(app.servers, s)
	return app
}

func (app *App) Start() chan error {
	return app.servers.Start()
}

func (app *App) Stop(ctx context.Context) error {
	return app.servers.Stop(ctx)
}
