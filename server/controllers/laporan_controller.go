package controllers

import (
	"net/http"
	"tecnhical-test/config"
	"tecnhical-test/middlewares"
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
}

type LaporanPenjualanItem struct {
	ID         uint    `json:"id"`
	NoFaktur   string  `json:"no_faktur"`
	Customer   string  `json:"customer"`
	Total      float64 `json:"total"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
	Username   string  `json:"username"`
	TotalItems int     `json:"total_items"`
}

type LaporanPembelianItem struct {
	ID         uint    `json:"id"`
	NoFaktur   string  `json:"no_faktur"`
	Supplier   string  `json:"supplier"`
	Total      float64 `json:"total"`
	Status     string  `json:"status"`
	CreatedAt  string  `json:"created_at"`
	Username   string  `json:"username"`
	TotalItems int     `json:"total_items"`
}

func LaporanStokAkhir(w http.ResponseWriter, r *http.Request) {
	var results []LaporanStokItem

	query := `
		SELECT 
			ms.id,
			ms.barang_id,
			mb.kode_barang,
			mb.nama_barang,
			mb.satuan,
			ms.stok_akhir,
			mb.harga_beli,
			mb.harga_jual,
			(ms.stok_akhir * mb.harga_beli) as nilai_stok
		FROM m_stocks ms
		JOIN master_barangs mb ON ms.barang_id = mb.id
		WHERE ms.deleted_at IS NULL AND mb.deleted_at IS NULL
		ORDER BY mb.nama_barang
	`

	if err := config.DB.Raw(query).Scan(&results).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var totalNilai float64
	for _, item := range results {
		totalNilai += item.NilaiStok
	}

	response := map[string]interface{}{
		"data":        results,
		"total_items": len(results),
		"total_nilai": totalNilai,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}

func LaporanPenjualan(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	query := config.DB.Table("jual_headers jh").
		Select(`jh.id, jh.no_faktur, jh.customer, jh.total, jh.status, 
		        jh.created_at, u.username, 
		        COUNT(jd.id) as total_items`).
		Joins("LEFT JOIN users u ON jh.user_id = u.id").
		Joins("LEFT JOIN jual_details jd ON jh.id = jd.jual_header_id AND jd.deleted_at IS NULL").
		Where("jh.deleted_at IS NULL").
		Group("jh.id, jh.no_faktur, jh.customer, jh.total, jh.status, jh.created_at, u.username")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(jh.created_at) BETWEEN ? AND ?", startDate, endDate)
	}

	var results []LaporanPenjualanItem
	if err := query.Order("jh.created_at DESC").Scan(&results).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
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

	query := config.DB.Table("beli_headers bh").
		Select(`bh.id, bh.no_faktur, bh.supplier, bh.total, bh.status, 
		        bh.created_at, u.username, 
		        COUNT(bd.id) as total_items`).
		Joins("LEFT JOIN users u ON bh.user_id = u.id").
		Joins("LEFT JOIN beli_details bd ON bh.id = bd.beli_header_id AND bd.deleted_at IS NULL").
		Where("bh.deleted_at IS NULL").
		Group("bh.id, bh.no_faktur, bh.supplier, bh.total, bh.status, bh.created_at, u.username")

	if startDate != "" && endDate != "" {
		query = query.Where("DATE(bh.created_at) BETWEEN ? AND ?", startDate, endDate)
	}

	var results []LaporanPembelianItem
	if err := query.Order("bh.created_at DESC").Scan(&results).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
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
