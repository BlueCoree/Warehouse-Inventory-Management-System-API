package models

import "gorm.io/gorm"

type MStok struct {
	gorm.Model
	BarangID  uint         `gorm:"not null" json:"barang_id"`
	Barang    MasterBarang `gorm:"foreignKey:BarangID"`
	StokAkhir int          `gorm:"type:integer;default:0" json:"stok_akhir"`
}
