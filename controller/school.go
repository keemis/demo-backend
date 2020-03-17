package controller

import (
	"demo-backend/model"
)

// SchoolController 学校接口
type SchoolController struct {
	BaseController
}

// List 学校列表
func (u *SchoolController) GetSchools() {
	u.Log.Debug("exec func: %v", "Controller GetSchools")

	po := &model.GetSchoolsReq{}
	u.ValidQuery(po)
	u.Log.Debug("request params: %+v", po)

	ret := u.srv.GetSchools(*po)
	u.ApiSuccess(ret)
}
