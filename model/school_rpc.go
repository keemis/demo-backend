package model

import (
	"github.com/keemis/library/validation"
)

// GetSchoolsReq 学校列表请求
type GetSchoolsReq struct {
	Page     int `json:"page" valid:"Min(0); Max(100)" `
	PageSize int `json:"page_size" valid:"Min(0); Max(50)" `
}

// Valid 自定义规则
func (u *GetSchoolsReq) Valid(v *validation.Validation) {
	if u.Page == 0 {
		u.Page = 1
	}
	if u.PageSize == 0 {
		u.PageSize = 50
	}
}
