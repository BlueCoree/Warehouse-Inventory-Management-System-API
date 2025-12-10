package models

import "gorm.io/gorm"

type JualDetail struct {
	gorm.Model
	JualHeaderID uint         `gorm:"not null" json:"jual_header_id"`
	JualHeader   JualHeader   `gorm:"foreignKey:JualHeaderID"`
	BarangID     uint         `gorm:"not null" json:"barang_id"`
	Barang       MasterBarang `gorm:"foreignKey:BarangID"`
	Qty          int          `gorm:"type:integer;not null" json:"qty"`
	Harga        float64      `gorm:"type:decimal(15,2);not null" json:"harga"`
	Subtotal     float64      `gorm:"type:decimal(15,2);not null" json:"subtotal"`
}
