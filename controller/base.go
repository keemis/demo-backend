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

// Before 每次请求都执行
func (u *BaseController) Before() {
	u.Log.Debug("exec func: %v", "BaseController Before")

	u.srv = service.New(u.Log)
}
