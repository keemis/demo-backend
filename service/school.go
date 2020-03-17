package service

import (
	"demo-backend/model"
)

// GetSchools 学校列表
func (u *Service) GetSchools(po model.GetSchoolsReq) []model.School {
	u.log.Debug("exec func: %v", "Service GetSchools")
	db := u.mysql.DB()
	var school []model.School
	db.Limit(po.PageSize).Offset((po.Page - 1) * po.PageSize).Find(&school)
	return school
}
