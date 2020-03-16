package controller

import (
	"demo-backend/service"
	"github.com/keemis/library/beeControll"
)

// BaseController 基类
type BaseController struct {
	beeControll.BaseController
	srv *service.Service
}

// Before 请求前执行
func (u *BaseController) Before() {
	u.Log.Debug("exec func: %v", "Before")

	u.srv = service.New(u.Log)
}
