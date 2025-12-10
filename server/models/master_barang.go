package models

import "gorm.io/gorm"

type MasterBarang struct {
	gorm.Model
	KodeBarang string  `gorm:"type:varchar(50);unique;not null" json:"kode_barang"`
	NamaBarang string  `gorm:"type:varchar(200);not null" json:"nama_barang"`
	Deskripsi  string  `gorm:"type:text" json:"deskripsi"`
	Satuan     string  `gorm:"type:varchar(50)" json:"satuan"`
	HargaBeli  float64 `gorm:"type:decimal(15,2);default:0" json:"harga_beli"`
	HargaJual  float64 `gorm:"type:decimal(15,2);default:0" json:"harga_jual"`
}
