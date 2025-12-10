package main

import (
	"log"
	"net/http"
	"os"

	"tecnhical-test/config"
	"tecnhical-test/controllers"
	"tecnhical-test/middlewares"
	"tecnhical-test/models"
	"tecnhical-test/seeds"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	config.ConnectDatabase()

	config.DB.AutoMigrate(
		&models.User{},
		&models.MasterBarang{},
		&models.MStok{},
		&models.HistoryStok{},
		&models.BeliHeader{},
		&models.BeliDetail{},
		&models.JualHeader{},
		&models.JualDetail{},
	)

	seeds.RunSeeder(config.DB)

	router := mux.NewRouter()

	router.HandleFunc("/api/auth/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/auth/register", controllers.Register).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	// Master Barang routes
	api.HandleFunc("/barang", controllers.GetBarang).Methods("GET")
	api.HandleFunc("/barang/stok", controllers.GetBarangStok).Methods("GET")
	api.HandleFunc("/barang/{id}", controllers.GetDetailBarang).Methods("GET")
	api.HandleFunc("/barang", controllers.CreateBarang).Methods("POST")
	api.HandleFunc("/barang/{id}", controllers.UpdateBarang).Methods("PUT")
	api.HandleFunc("/barang/{id}", controllers.DeleteBarang).Methods("DELETE")

	// Stok routes
	api.HandleFunc("/stok", controllers.GetStok).Methods("GET")
	api.HandleFunc("/stok/{barang_id}", controllers.GetStokByBarang).Methods("GET")

	// History Stok routes
	api.HandleFunc("/history-stok", controllers.GetHistoryStok).Methods("GET")
	api.HandleFunc("/history-stok/{barang_id}", controllers.GetHistoryByBarang).Methods("GET")

	// Pembelian routes
	api.HandleFunc("/pembelian", controllers.CreatePembelian).Methods("POST")
	api.HandleFunc("/pembelian", controllers.GetAllPembelian).Methods("GET")
	api.HandleFunc("/pembelian/{id}", controllers.GetDetailPembelian).Methods("GET")

	// Penjualan routes
	api.HandleFunc("/penjualan", controllers.CreatePenjualan).Methods("POST")
	api.HandleFunc("/penjualan", controllers.GetAllPenjualan).Methods("GET")
	api.HandleFunc("/penjualan/{id}", controllers.GetDetailPenjualan).Methods("GET")

	// Laporan routes
	api.HandleFunc("/laporan/stok", controllers.LaporanStokAkhir).Methods("GET")
	api.HandleFunc("/laporan/penjualan", controllers.LaporanPenjualan).Methods("GET")
	api.HandleFunc("/laporan/pembelian", controllers.LaporanPembelian).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
