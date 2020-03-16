package service

import (
	"demo-backend/model"
)

// GetSchools 学校列表
func (u *Service) GetSchools() []model.School {
	u.log.Debug("exec func: %v", "Service GetSchools")
	var school []model.School
	db := u.mysql.DB()
	db.Find(&school)
	return school
}
