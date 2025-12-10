package seeds

import (
	"fmt"
	"log"
	"tecnhical-test/helpers"
	"tecnhical-test/models"

	"gorm.io/gorm"
)

func RunSeeder(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)

	if count > 0 {
		log.Println("Database sudah berisi data, skip seeding...")
		return
	}

	log.Println("Memulai seeding data...")

	if err := seedUsers(db); err != nil {
		log.Fatal("Seeding Users Gagal:", err)
	}
	if err := seedMasterBarang(db); err != nil {
		log.Fatal("Seeding Master Barang Gagal:", err)
	}
	if err := seedMStok(db); err != nil {
		log.Fatal("Seeding MStok Gagal:", err)
	}
	if err := seedPembelian(db); err != nil {
		log.Fatal("Seeding Pembelian Gagal:", err)
	}
	if err := seedPenjualan(db); err != nil {
		log.Fatal("Seeding Penjualan Gagal:", err)
	}
	if err := seedHistoryStok(db); err != nil {
		log.Fatal("Seeding History Stok Gagal:", err)
	}

	log.Println("Seeding selesai!")
}

func seedUsers(db *gorm.DB) error {
	adminPass, _ := helpers.HashPassword("admin123")
	staff1Pass, _ := helpers.HashPassword("staff123")
	staff2Pass, _ := helpers.HashPassword("staff234")

	users := []models.User{
		{Username: "admin", Password: adminPass, Email: "admin@warehouse.com", Fullname: "Administrator System", Role: "admin"},
		{Username: "staff1", Password: staff1Pass, Email: "staff1@warehouse.com", Fullname: "Staff Gudang A", Role: "staff"},
		{Username: "staff2", Password: staff2Pass, Email: "staff2@warehouse.com", Fullname: "Staff Gudang B", Role: "staff"},
	}
	if err := db.Create(&users).Error; err != nil {
		return err
	}
	fmt.Println("Seed Users done")
	return nil
}

func seedMasterBarang(db *gorm.DB) error {
	barang := []models.MasterBarang{
		{KodeBarang: "BRG001", NamaBarang: "Laptop Dell XPS 13", Deskripsi: "Laptop Business Grade", Satuan: "unit", HargaBeli: 15000000, HargaJual: 17500000},
		{KodeBarang: "BRG002", NamaBarang: "Mouse Wireless Logitech", Deskripsi: "Mouse Wireless 2.4GHz", Satuan: "pcs", HargaBeli: 250000, HargaJual: 350000},
		{KodeBarang: "BRG003", NamaBarang: "Keyboard Mechanical", Deskripsi: "Keyboard Mechanical RGB", Satuan: "pcs", HargaBeli: 800000, HargaJual: 1200000},
		{KodeBarang: "BRG004", NamaBarang: "Monitor 24 inch", Deskripsi: "Monitor LED 24 inch Full HD", Satuan: "unit", HargaBeli: 2000000, HargaJual: 2800000},
		{KodeBarang: "BRG005", NamaBarang: "Webcam HD 1080p", Deskripsi: "Webcam High Definition", Satuan: "pcs", HargaBeli: 450000, HargaJual: 650000},
	}
	if err := db.Create(&barang).Error; err != nil {
		return err
	}
	fmt.Println("Seed MasterBarang done")
	return nil
}

func seedMStok(db *gorm.DB) error {
	stok := []models.MStok{
		{BarangID: 1, StokAkhir: 10},
		{BarangID: 2, StokAkhir: 50},
		{BarangID: 3, StokAkhir: 30},
		{BarangID: 4, StokAkhir: 15},
		{BarangID: 5, StokAkhir: 25},
	}
	if err := db.Create(&stok).Error; err != nil {
		return err
	}
	fmt.Println("Seed MStock done")
	return nil
}

func seedPembelian(db *gorm.DB) error {
	header1 := models.BeliHeader{
		NoFaktur: "BLI001", Supplier: "PT Supplier Elektronik", Total: 32500000, UserID: 2, Status: "selesai",
		Details: []models.BeliDetail{
			{BarangID: 1, Qty: 2, Harga: 15000000, Subtotal: 30000000},
			{BarangID: 2, Qty: 10, Harga: 250000, Subtotal: 2500000},
		},
	}
	header2 := models.BeliHeader{
		NoFaktur: "BLI002", Supplier: "CV Komputer Jaya", Total: 12500000, UserID: 3, Status: "selesai",
		Details: []models.BeliDetail{
			{BarangID: 3, Qty: 5, Harga: 800000, Subtotal: 4000000},
			{BarangID: 4, Qty: 3, Harga: 2000000, Subtotal: 6000000},
			{BarangID: 5, Qty: 4, Harga: 450000, Subtotal: 1800000},
		},
	}

	if err := db.Create(&header1).Error; err != nil {
		return err
	}
	if err := db.Create(&header2).Error; err != nil {
		return err
	}
	fmt.Println("Seed Pembelian done")
	return nil
}

func seedPenjualan(db *gorm.DB) error {
	header1 := models.JualHeader{
		NoFaktur: "JUAL001", Customer: "PT Customer Indonesia", Total: 18700000, UserID: 2, Status: "selesai",
		Details: []models.JualDetail{
			{BarangID: 1, Qty: 1, Harga: 17500000, Subtotal: 17500000},
			{BarangID: 2, Qty: 2, Harga: 350000, Subtotal: 700000},
			{BarangID: 3, Qty: 1, Harga: 1200000, Subtotal: 1200000},
		},
	}
	header2 := models.JualHeader{
		NoFaktur: "JUAL002", Customer: "CV Tech Solution", Total: 4150000, UserID: 3, Status: "selesai",
		Details: []models.JualDetail{
			{BarangID: 2, Qty: 5, Harga: 350000, Subtotal: 1750000},
			{BarangID: 4, Qty: 1, Harga: 2800000, Subtotal: 2800000},
		},
	}

	if err := db.Create(&header1).Error; err != nil {
		return err
	}
	if err := db.Create(&header2).Error; err != nil {
		return err
	}
	fmt.Println("Seed Penjualan done")
	return nil
}

func seedHistoryStok(db *gorm.DB) error {
	history := []models.HistoryStok{
		{BarangID: 1, UserID: 2, JenisTransaksi: "masuk", Jumlah: 2, StokSebelum: 0, StokSesudah: 2, Keterangan: "Pembelian BLI001"},
		{BarangID: 2, UserID: 2, JenisTransaksi: "masuk", Jumlah: 10, StokSebelum: 0, StokSesudah: 10, Keterangan: "Pembelian BLI001"},
		{BarangID: 3, UserID: 3, JenisTransaksi: "masuk", Jumlah: 5, StokSebelum: 0, StokSesudah: 5, Keterangan: "Pembelian BLI002"},
	}
	if err := db.Create(&history).Error; err != nil {
		return err
	}
	fmt.Println("Seed HistoryStok done")
	return nil
}
