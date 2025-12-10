package controllers

import (
	"net/http"
	"tecnhical-test/config"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"

	"gorm.io/gorm"
)

type LaporanStokItem struct {
	ID         uint    `json:"id"`
	BarangID   uint    `json:"barang_id"`
	KodeBarang string  `json:"kode_barang"`
	NamaBarang string  `json:"nama_barang"`
	Satuan     string  `json:"satuan"`
	StokAkhir  int     `json:"stok_akhir"`
	HargaBeli  float64 `json:"harga_beli"`
	HargaJual  float64 `json:"harga_jual"`
	NilaiStok  float64 `json:"nilai_stok"`
	Deleted    bool    `json:"deleted"`
}

type LaporanPenjualanItem struct {
	ID         uint                         `json:"id"`
	NoFaktur   string                       `json:"no_faktur"`
	Customer   string                       `json:"customer"`
	Total      float64                      `json:"total"`
	Status     string                       `json:"status"`
	CreatedAt  string                       `json:"created_at"`
	Username   string                       `json:"username"`
	TotalItems int                          `json:"total_items"`
	Details    []LaporanPenjualanDetailItem `json:"details"`
}

type LaporanPenjualanDetailItem struct {
	NamaBarang string  `json:"nama_barang"`
	Qty        int     `json:"qty"`
	Harga      float64 `json:"harga"`
	Subtotal   float64 `json:"subtotal"`
}

type LaporanPembelianItem struct {
	ID         uint                         `json:"id"`
	NoFaktur   string                       `json:"no_faktur"`
	Supplier   string                       `json:"supplier"`
	Total      float64                      `json:"total"`
	Status     string                       `json:"status"`
	CreatedAt  string                       `json:"created_at"`
	Username   string                       `json:"username"`
	TotalItems int                          `json:"total_items"`
	Details    []LaporanPembelianDetailItem `json:"details"`
}

type LaporanPembelianDetailItem struct {
	NamaBarang string  `json:"nama_barang"`
	Qty        int     `json:"qty"`
	Harga      float64 `json:"harga"`
	Subtotal   float64 `json:"subtotal"`
}

func LaporanStokAkhir(w http.ResponseWriter, r *http.Request) {
	var stoks []models.MStok
	if err := config.DB.Preload("Barang", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Find(&stoks).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []LaporanStokItem
	var totalNilai float64
	var activeCount int

	for _, s := range stoks {
		nilaiStok := float64(s.StokAkhir) * s.Barang.HargaBeli
		isDeleted := s.Barang.DeletedAt.Valid
		results = append(results, LaporanStokItem{
			ID:         s.ID,
			BarangID:   s.BarangID,
			KodeBarang: s.Barang.KodeBarang,
			NamaBarang: s.Barang.NamaBarang,
			Satuan:     s.Barang.Satuan,
			StokAkhir:  s.StokAkhir,
			HargaBeli:  s.Barang.HargaBeli,
			HargaJual:  s.Barang.HargaJual,
			NilaiStok:  nilaiStok,
			Deleted:    isDeleted,
		})
		if !isDeleted {
			totalNilai += nilaiStok
			activeCount++
		}
	}

	response := map[string]interface{}{
		"data":        results,
		"total_items": activeCount,
		"total_nilai": totalNilai,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}

func LaporanPenjualan(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var headers []models.JualHeader
	query := config.DB.Preload("User").Preload("Details.Barang", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("Details")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate)
	}

	if err := query.Order("created_at DESC").Find(&headers).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []LaporanPenjualanItem
	for _, h := range headers {
		var details []LaporanPenjualanDetailItem
		for _, d := range h.Details {
			namaBarang := d.NamaBarang
			if namaBarang == "" && d.Barang.ID != 0 {
				namaBarang = d.Barang.NamaBarang
			}
			details = append(details, LaporanPenjualanDetailItem{
				NamaBarang: namaBarang,
				Qty:        d.Qty,
				Harga:      d.Harga,
				Subtotal:   d.Subtotal,
			})
		}
		results = append(results, LaporanPenjualanItem{
			ID:         h.ID,
			NoFaktur:   h.NoFaktur,
			Customer:   h.Customer,
			Total:      h.Total,
			Status:     h.Status,
			CreatedAt:  h.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Username:   h.User.Username,
			TotalItems: len(h.Details),
			Details:    details,
		})
	}

	var grandTotal float64
	var totalTransaksi int
	for _, item := range results {
		grandTotal += item.Total
		totalTransaksi++
	}

	response := map[string]interface{}{
		"data":            results,
		"total_transaksi": totalTransaksi,
		"grand_total":     grandTotal,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}

func LaporanPembelian(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	var headers []models.BeliHeader
	query := config.DB.Preload("User").Preload("Details.Barang", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("Details")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(created_at) BETWEEN ? AND ?", startDate, endDate)
	}

	if err := query.Order("created_at DESC").Find(&headers).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []LaporanPembelianItem
	for _, h := range headers {
		var details []LaporanPembelianDetailItem
		for _, d := range h.Details {
			namaBarang := d.NamaBarang
			if namaBarang == "" && d.Barang.ID != 0 {
				namaBarang = d.Barang.NamaBarang
			}
			details = append(details, LaporanPembelianDetailItem{
				NamaBarang: namaBarang,
				Qty:        d.Qty,
				Harga:      d.Harga,
				Subtotal:   d.Subtotal,
			})
		}
		results = append(results, LaporanPembelianItem{
			ID:         h.ID,
			NoFaktur:   h.NoFaktur,
			Supplier:   h.Supplier,
			Total:      h.Total,
			Status:     h.Status,
			CreatedAt:  h.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Username:   h.User.Username,
			TotalItems: len(h.Details),
			Details:    details,
		})
	}

	var grandTotal float64
	var totalTransaksi int
	for _, item := range results {
		grandTotal += item.Total
		totalTransaksi++
	}

	response := map[string]interface{}{
		"data":            results,
		"total_transaksi": totalTransaksi,
		"grand_total":     grandTotal,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}
