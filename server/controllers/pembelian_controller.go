package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tecnhical-test/config"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"

	"github.com/gorilla/mux"
)

type CreatePembelianRequest struct {
	NoFaktur string                     `json:"no_faktur"`
	Supplier string                     `json:"supplier"`
	Status   string                     `json:"status"`
	Details  []CreatePembelianDetailReq `json:"details"`
}

type CreatePembelianDetailReq struct {
	BarangID uint    `json:"barang_id"`
	Qty      int     `json:"qty"`
	Harga    float64 `json:"harga"`
}

type PembelianResponse struct {
	Header  models.BeliHeader         `json:"header"`
	Details []PembelianDetailResponse `json:"details"`
}

type PembelianDetailResponse struct {
	ID       uint    `json:"id"`
	BarangID uint    `json:"barang_id"`
	Qty      int     `json:"qty"`
	Harga    float64 `json:"harga"`
	Subtotal float64 `json:"subtotal"`
	Barang   struct {
		KodeBarang string `json:"kode_barang"`
		NamaBarang string `json:"nama_barang"`
		Satuan     string `json:"satuan"`
	} `json:"barang"`
}

func CreatePembelian(w http.ResponseWriter, r *http.Request) {
	var req CreatePembelianRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		middlewares.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.NoFaktur == "" || req.Supplier == "" || len(req.Details) == 0 {
		middlewares.ErrorResponse(w, http.StatusUnprocessableEntity, "No faktur, supplier, and details are required")
		return
	}

	userClaims := middlewares.GetUserFromContext(r)
	if userClaims == nil {
		middlewares.ErrorResponse(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	tx := config.DB.Begin()

	// Calculate total
	var total float64
	for _, detail := range req.Details {
		subtotal := float64(detail.Qty) * detail.Harga
		total += subtotal
	}

	if req.Status == "" {
		req.Status = "selesai"
	}

	// Create header
	header := models.BeliHeader{
		NoFaktur: req.NoFaktur,
		Supplier: req.Supplier,
		Total:    total,
		UserID:   userClaims.UserID,
		Status:   req.Status,
	}

	if err := tx.Create(&header).Error; err != nil {
		tx.Rollback()
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create pembelian")
		return
	}

	// Create details and update stock
	for _, detail := range req.Details {
		subtotal := float64(detail.Qty) * detail.Harga

		// Validate barang exists
		var barang models.MasterBarang
		if err := tx.First(&barang, detail.BarangID).Error; err != nil {
			tx.Rollback()
			middlewares.ErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Barang with ID %d not found", detail.BarangID))
			return
		}

		// Create detail
		beliDetail := models.BeliDetail{
			BeliHeaderID: header.ID,
			BarangID:     detail.BarangID,
			Qty:          detail.Qty,
			Harga:        detail.Harga,
			Subtotal:     subtotal,
		}

		if err := tx.Create(&beliDetail).Error; err != nil {
			tx.Rollback()
			middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create detail")
			return
		}

		// Update stock
		var stok models.MStok
		if err := tx.Where("barang_id = ?", detail.BarangID).First(&stok).Error; err != nil {
			// Create new stock if not exists
			stok = models.MStok{
				BarangID:  detail.BarangID,
				StokAkhir: 0,
			}
			if err := tx.Create(&stok).Error; err != nil {
				tx.Rollback()
				middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create stock")
				return
			}
		}

		stokSebelum := stok.StokAkhir
		stok.StokAkhir += detail.Qty

		if err := tx.Save(&stok).Error; err != nil {
			tx.Rollback()
			middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to update stock")
			return
		}

		// Create history
		history := models.HistoryStok{
			BarangID:       detail.BarangID,
			UserID:         userClaims.UserID,
			JenisTransaksi: "masuk",
			Jumlah:         detail.Qty,
			StokSebelum:    stokSebelum,
			StokSesudah:    stok.StokAkhir,
			Keterangan:     fmt.Sprintf("Pembelian %s", req.NoFaktur),
		}

		if err := tx.Create(&history).Error; err != nil {
			tx.Rollback()
			middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to create history")
			return
		}
	}

	tx.Commit()

	// Reload header with details
	config.DB.Preload("Details.Barang").Preload("User").First(&header, header.ID)

	middlewares.SuccessResponse(w, "Pembelian created successfully", header)
}

func GetAllPembelian(w http.ResponseWriter, r *http.Request) {
	var headers []models.BeliHeader

	if err := config.DB.Preload("Details.Barang").Preload("User").Order("created_at DESC").Find(&headers).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", headers)
}

func GetDetailPembelian(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var header models.BeliHeader
	if err := config.DB.Preload("Details.Barang").Preload("User").First(&header, id).Error; err != nil {
		middlewares.ErrorResponse(w, http.StatusNotFound, "Pembelian not found")
		return
	}

	var details []PembelianDetailResponse
	for _, d := range header.Details {
		detail := PembelianDetailResponse{
			ID:       d.ID,
			BarangID: d.BarangID,
			Qty:      d.Qty,
			Harga:    d.Harga,
			Subtotal: d.Subtotal,
		}
		detail.Barang.KodeBarang = d.Barang.KodeBarang
		detail.Barang.NamaBarang = d.Barang.NamaBarang
		detail.Barang.Satuan = d.Barang.Satuan

		details = append(details, detail)
	}

	response := PembelianResponse{
		Header:  header,
		Details: details,
	}

	middlewares.SuccessResponse(w, "Data retrieved successfully", response)
}
