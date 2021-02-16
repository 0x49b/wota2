package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Text    string `json:"text" form:"text" query:"text"`
	EventID uint   `gorm:"column:event_id" json:"event_id"`
	Event   Event  `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
