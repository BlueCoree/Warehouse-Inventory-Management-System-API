package models

import "gorm.io/gorm"

type HistoryStok struct {
	gorm.Model
	BarangID       uint         `gorm:"not null" json:"barang_id"`
	Barang         MasterBarang `gorm:"foreignKey:BarangID"`
	NamaBarang     string       `gorm:"type:varchar(255)" json:"nama_barang"`
	UserID         uint         `gorm:"not null" json:"user_id"`
	User           User         `gorm:"foreignKey:UserID"`
	JenisTransaksi string       `gorm:"type:varchar(50);not null" json:"jenis_transaksi"`
	Jumlah         int          `gorm:"type:integer;not null" json:"jumlah"`
	StokSebelum    int          `gorm:"type:integer;not null" json:"stok_sebelum"`
	StokSesudah    int          `gorm:"type:integer;not null" json:"stok_sesudah"`
	Keterangan     string       `gorm:"type:TEXT" json:"keterangan"`
}
