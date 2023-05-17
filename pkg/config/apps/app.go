package apps

import (
	"errors"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"udious.com/mockingbird/channel/pkg/config"
)

type App struct {
	Cfg       config.Config
	db        *gorm.DB
	dbinit    sync.Once
	dberr     error
	redis     *redis.Client
	rediserr  error
	redisinit sync.Once
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
