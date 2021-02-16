package models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Name      string    `json:"name" form:"name" query:"name"`
	Date      time.Time `json:"date" form:"date" query:"date"`
	Autoclose bool      `json:"autoclose" form:"autoclose" query:"autoclose"`
	UserID    uint      `json:"user_id" form:"user_id" query:"user_id"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
