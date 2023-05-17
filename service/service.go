package service

import (
	"github.com/bsm/redislock"
	"github.com/dev-mockingbird/logf"
	"gorm.io/gorm"
)

const (
	InvalidArguments = "invalid-arguments"
)

type Service struct {
	DB         *gorm.DB
	Logger     logf.Logger
	LockClient *redislock.Client
}

func (s Service) HelloWorld() {

}
