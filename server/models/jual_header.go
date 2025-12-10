package models

import "gorm.io/gorm"

type JualHeader struct {
	gorm.Model
	NoFaktur string       `gorm:"type:varchar(100);unique;not null" json:"no_faktur"`
	Customer string       `gorm:"type:varchar(200);not null" json:"customer"`
	Total    float64      `gorm:"type:decimal(15,2);default:0" json:"total"`
	Status   string       `gorm:"type:varchar(50);not null;default:'selesai'" json:"status"`
	UserID   uint         `gorm:"not null" json:"user_id"`
	User     User         `gorm:"foreignKey:UserID"`
	Details  []JualDetail `gorm:"foreignKey:JualHeaderID" json:"details"`
}
