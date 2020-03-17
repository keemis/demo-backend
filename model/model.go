package model

import (
	"time"
)

// School 学校
type School struct {
	ID         int64     `json:"id" gorm:"column:id"`
	Name       string    `json:"name" gorm:"column:name"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time"`
}
