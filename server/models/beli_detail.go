package models

import "gorm.io/gorm"

type BeliDetail struct {
	gorm.Model
	BeliHeaderID uint         `gorm:"not null" json:"beli_header_id"`
	BeliHeader   BeliHeader   `gorm:"foreignKey:BeliHeaderID"`
	BarangID     uint         `gorm:"not null" json:"barang_id"`
	Barang       MasterBarang `gorm:"foreignKey:BarangID"`
	NamaBarang   string       `gorm:"type:varchar(255)" json:"nama_barang"`
	Qty          int          `gorm:"type:integer;not null" json:"qty"`
	Harga        float64      `gorm:"type:decimal(15,2);not null" json:"harga"`
	Subtotal     float64      `gorm:"type:decimal(15,2);not null" json:"subtotal"`
}
