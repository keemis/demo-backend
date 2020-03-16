package service

import (
	"github.com/keemis/library/curl"
	"github.com/keemis/library/logs"
	"github.com/keemis/library/mysql"
)

// Service 服务实现
type Service struct {
	mysql mysql.Orm
	log   logs.Logger
	curl  curl.Curl
}

// NEW 创建一个服务
func New(log logs.Logger) *Service {
	return &Service{
		log:   log,
		mysql: mysql.New(mysql.OptLogger(log)),
		curl:  curl.New(curl.OptLogger(log)),
	}

}
