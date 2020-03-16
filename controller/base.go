package controller

import (
	"demo-backend/service"
	"github.com/keemis/library/beeControll"
	"github.com/keemis/library/logs"
)

// BaseController 基类
type BaseController struct {
	beeControll.BaseController
	log logs.Logger
	srv *service.Service
}

// Before 请求前执行
func (u *BaseController) Before() {
	u.log = logs.New()
	u.srv = service.New(u.log)
	u.log.Debug("exec func: %v", "Before")
}
