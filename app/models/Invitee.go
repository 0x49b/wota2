package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Invitee struct {
	gorm.Model
	Name     string    `json:"name" form:"firstname" query:"name"`
	Email    string    `json:"email" form:"firstname" query:"email"`
	Uuid     uuid.UUID `gorm:"type:uuid,not null;" json:"uuid" form:"uuid" query:"uuid"`
	LoggedIn bool      `json:"logged_in" form:"logged_in" query:"logged_in"`
	EventID  uint      `gorm:"column:event_id" json:"event_id"`
	Event    Event     `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserID   uint      `gorm:"column:user_id" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
