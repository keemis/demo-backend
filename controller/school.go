package controller

// SchoolController 学校接口
type SchoolController struct {
	BaseController
}

// List 学校列表
func (u *SchoolController) GetSchools() {
	u.log.Debug("exec func: %v", "Controller GetSchools")
	ret := u.srv.GetSchools()
	u.ApiSuccess(ret)
}
