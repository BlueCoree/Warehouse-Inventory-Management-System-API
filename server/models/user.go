package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Email    string `gorm:"type:varchar(150);unique;not null" json:"email"`
	Fullname string `gorm:"type:varchar(200)" json:"full_name"`
	Role     string `gorm:"type:varchar(50);default:'staff'" json:"role"`
}
