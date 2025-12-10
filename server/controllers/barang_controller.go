package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tecnhical-test/config"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"

	"github.com/gorilla/mux"
)

type CreateBarangRequest struct {
	KodeBarang string  `json:"kode_barang"`
	NamaBarang string  `json:"nama_barang"`
	Deskripsi  string  `json:"deskripsi"`
	Satuan     string  `json:"satuan"`
	HargaBeli  float64 `json:"harga_beli"`
	HargaJual  float64 `json:"harga_jual"`
	StokAkhir  int     `json:"stok_akhir"`
}

type BarangWithStok struct {
	models.MasterBarang
	StokAkhir int `json:"stok_akhir"`
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	query := config.DB.Model(&models.MasterBarang{})

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where("LOWER(kode_barang) LIKE ? OR LOWER(nama_barang) LIKE ?", searchPattern, searchPattern)
	}

	var total int64
	query.Count(&total)

	var barang []models.MasterBarang
	if err := query.Offset(offset).Limit(limit).Find(&barang).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	meta := middlewares.PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	middlewares.JSONResponseWithMeta(w, http.StatusOK, true, "Data retrieved successfully", barang, meta)
}

func GetBarangStok(w http.ResponseWriter, r *http.Request) {
	var barang []models.MasterBarang
	if err := config.DB.Find(&barang).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	var results []BarangWithStok
	for _, b := range barang {
		var stok models.MStok
		config.DB.Where("barang_id = ?", b.ID).First(&stok)

		results = append(results, BarangWithStok{
			MasterBarang: b,
			StokAkhir:    stok.StokAkhir,
		})
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", results)
}

func GetDetailBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var barang models.MasterBarang
	if err := config.DB.First(&barang, id).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusNotFound, "Barang not found")
		return
	}

	var stok models.MStok
	if err := config.DB.Where("barang_id = ?", id).First(&stok).Error; err == nil {
		result := BarangWithStok{
			MasterBarang: barang,
			StokAkhir:    stok.StokAkhir,
		}
		middlewares.SuccessResponse(w, "Data retrieved successfully", result)
	} else {
		result := BarangWithStok{
			MasterBarang: barang,
			StokAkhir:    0,
		}
		middlewares.SuccessResponse(w, "Data retrieved successfully", result)
	}
}

func CreateBarang(w http.ResponseWriter, r *http.Request) {
	var req CreateBarangRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middlewares.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.KodeBarang == "" || req.NamaBarang == "" {
		middlewares.ErrorResponse(w, http.StatusUnprocessableEntity, "Kode barang and nama barang are required")
		return
	}

	barang := models.MasterBarang{
		KodeBarang: req.KodeBarang,
		NamaBarang: req.NamaBarang,
		Deskripsi:  req.Deskripsi,
		Satuan:     req.Satuan,
		HargaBeli:  req.HargaBeli,
		HargaJual:  req.HargaJual,
	}

	tx := config.DB.Begin()

	if err := tx.Create(&barang).Error; err != nil {
		tx.Rollback()
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create barang")
		return
	}

	// Init stok
	stok := models.MStok{
		BarangID:  barang.ID,
		StokAkhir: req.StokAkhir,
	}

	if err := tx.Create(&stok).Error; err != nil {
		tx.Rollback()
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to initialize stock")
		return
	}

	tx.Commit()

	middlewares.SuccessResponse(w, "Barang created successfully", barang)
}

func UpdateBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var barang models.MasterBarang
	if err := config.DB.First(&barang, id).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusNotFound, "Barang not found")
		return
	}

	var req CreateBarangRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middlewares.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.KodeBarang != "" {
		barang.KodeBarang = req.KodeBarang
	}
	if req.NamaBarang != "" {
		barang.NamaBarang = req.NamaBarang
	}
	barang.Deskripsi = req.Deskripsi
	barang.Satuan = req.Satuan
	barang.HargaBeli = req.HargaBeli
	barang.HargaJual = req.HargaJual

	tx := config.DB.Begin()

	if err := tx.Save(&barang).Error; err != nil {
		tx.Rollback()
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to update barang")
		return
	}

	// Update stok
	var stok models.MStok
	if err := tx.Where("barang_id = ?", barang.ID).First(&stok).Error; err == nil {
		stok.StokAkhir = req.StokAkhir
		tx.Save(&stok)
	} else {
		newStok := models.MStok{
			BarangID:  barang.ID,
			StokAkhir: req.StokAkhir,
		}
		tx.Create(&newStok)
	}

	tx.Commit()

	middlewares.SuccessResponse(w, "Barang updated successfully", barang)
}

func DeleteBarang(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var barang models.MasterBarang
	if err := config.DB.First(&barang, id).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusNotFound, "Barang not found")
		return
	}

	if err := config.DB.Delete(&barang).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete barang")
		return
	}

	middlewares.SuccessResponse(w, "Barang deleted successfully", nil)
}
