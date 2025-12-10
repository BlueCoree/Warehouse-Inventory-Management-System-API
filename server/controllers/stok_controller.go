package controllers

import (
	"net/http"
	"tecnhical-test/config"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type StokResponse struct {
	ID        uint                `json:"id"`
	BarangID  uint                `json:"barang_id"`
	StokAkhir int                 `json:"stok_akhir"`
	UpdatedAt string              `json:"updated_at"`
	Barang    models.MasterBarang `json:"barang"`
}

type HistoryStokResponse struct {
	ID             uint   `json:"id"`
	BarangID       uint   `json:"barang_id"`
	NamaBarang     string `json:"nama_barang"`
	UserID         uint   `json:"user_id"`
	JenisTransaksi string `json:"jenis_transaksi"`
	Jumlah         int    `json:"jumlah"`
	StokSebelum    int    `json:"stok_sebelum"`
	StokSesudah    int    `json:"stok_sesudah"`
	Keterangan     string `json:"keterangan"`
	CreatedAt      string `json:"created_at"`
	Barang         struct {
		KodeBarang string `json:"kode_barang"`
		NamaBarang string `json:"nama_barang"`
	} `json:"barang"`
	User struct {
		Username string `json:"username"`
		Fullname string `json:"full_name"`
	} `json:"user"`
}

func GetStok(w http.ResponseWriter, r *http.Request) {
	var stoks []models.MStok

	if err := config.DB.Preload("Barang").Find(&stoks).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []StokResponse
	for _, s := range stoks {
		results = append(results, StokResponse{
			ID:        s.ID,
			BarangID:  s.BarangID,
			StokAkhir: s.StokAkhir,
			UpdatedAt: s.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			Barang:    s.Barang,
		})
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", results)
}

func GetStokByBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangID := vars["barang_id"]

	var stok models.MStok
	if err := config.DB.Preload("Barang").Where("barang_id = ?", barangID).First(&stok).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusNotFound, "Stock not found")
		return
	}

	result := StokResponse{
		ID:        stok.ID,
		BarangID:  stok.BarangID,
		StokAkhir: stok.StokAkhir,
		UpdatedAt: stok.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Barang:    stok.Barang,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", result)
}

func GetHistoryStok(w http.ResponseWriter, r *http.Request) {
	var histories []models.HistoryStok

	if err := config.DB.Preload("Barang", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("User").Order("created_at DESC").Find(&histories).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []HistoryStokResponse
	for _, h := range histories {
		result := HistoryStokResponse{
			ID:             h.ID,
			BarangID:       h.BarangID,
			NamaBarang:     h.NamaBarang,
			UserID:         h.UserID,
			JenisTransaksi: h.JenisTransaksi,
			Jumlah:         h.Jumlah,
			StokSebelum:    h.StokSebelum,
			StokSesudah:    h.StokSesudah,
			Keterangan:     h.Keterangan,
			CreatedAt:      h.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		if h.Barang.ID != 0 {
			result.Barang.KodeBarang = h.Barang.KodeBarang
			result.Barang.NamaBarang = h.Barang.NamaBarang

			if result.NamaBarang == "" {
				result.NamaBarang = h.Barang.NamaBarang
			}
		} else {
			result.Barang.NamaBarang = h.NamaBarang
		}
		result.User.Username = h.User.Username
		result.User.Fullname = h.User.Fullname

		results = append(results, result)
	}

	response := map[string]interface{}{
		"data":  results,
		"total": len(results),
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}

func GetHistoryByBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	barangID := vars["barang_id"]

	var histories []models.HistoryStok

	if err := config.DB.Preload("Barang", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("User").
		Where("barang_id = ?", barangID).
		Order("created_at DESC").
		Find(&histories).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []HistoryStokResponse
	for _, h := range histories {
		result := HistoryStokResponse{
			ID:             h.ID,
			BarangID:       h.BarangID,
			NamaBarang:     h.NamaBarang,
			UserID:         h.UserID,
			JenisTransaksi: h.JenisTransaksi,
			Jumlah:         h.Jumlah,
			StokSebelum:    h.StokSebelum,
			StokSesudah:    h.StokSesudah,
			Keterangan:     h.Keterangan,
			CreatedAt:      h.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if h.Barang.ID != 0 {
			result.Barang.KodeBarang = h.Barang.KodeBarang
			result.Barang.NamaBarang = h.Barang.NamaBarang

			if result.NamaBarang == "" {
				result.NamaBarang = h.Barang.NamaBarang
			}
		} else {
			result.Barang.NamaBarang = h.NamaBarang
		}
		result.User.Username = h.User.Username
		result.User.Fullname = h.User.Fullname

		results = append(results, result)
	}

	response := map[string]interface{}{
		"data":  results,
		"total": len(results),
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}
