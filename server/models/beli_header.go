package models

import (
	"gorm.io/gorm"
)

type BeliHeader struct {
	gorm.Model
	NoFaktur string       `gorm:"type:varchar(100);unique;not null" json:"no_faktur"`
	Supplier string       `gorm:"type:varchar(200);not null" json:"supplier"`
	Total    float64      `gorm:"type:decimal(15,2);default:0" json:"total"`
	Status   string       `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	UserID   uint         `gorm:"not null" json:"user_id"`
	User     User         `gorm:"foreignKey:UserID"`
	Details  []BeliDetail `gorm:"foreignKey:BeliHeaderID" json:"details"`
}
