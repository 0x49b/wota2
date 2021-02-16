package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Name      string `json:"name" form:"name" query:"name" gorm:"unique,not null"`
	FirstName string `json:"firstname" form:"firstname" query:"firstname"`
	Password  string `json:"-" xml:"-" form:"-" query:"-" gorm:"not null"`
	Email     string `json:"email" query:"email" gorm:"unique,not null"`
	RoleID    uint   `gorm:"column:role_id" json:"role_id"`
	Role      Role   `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
