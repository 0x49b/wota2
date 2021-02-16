package models

import "gorm.io/gorm"

type Result struct {
	gorm.Model
	QuestionID uint     `gorm:"column:question_id" json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Answer     bool     `json:"answer" form:"answer" query:"answer"`
	InviteeID  uint     `gorm:"column:invitee_id" json:"invitee_id"`
	Invitee    Invitee  `gorm:"foreignKey:InviteeID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
