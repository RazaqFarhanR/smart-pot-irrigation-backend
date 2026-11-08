package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"smart-irrigation-backend/internal/api"
	"smart-irrigation-backend/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	log.Printf("[INFO] Memulai Smart Pot Irrigation Backend (Inisiasi Proyek) pada port :%s", cfg.Port)

	// Inisialisasi router dasar
	router := api.NewRouter()

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Goroutine server listener
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] Gagal menjalankan HTTP server: %v", err)
		}
	}()

	fmt.Printf("\n🚀 Smart Pot Irrigation Backend aktif di http://localhost:%s\n", cfg.Port)
	fmt.Printf("📊 Endpoint Inisiasi:\n")
	fmt.Printf("   - Health Check: http://localhost:%s/healthz\n\n", cfg.Port)

	// Graceful shutdown listener
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[INFO] Menerima sinyal terminasi, mematikan server secara graceful...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("[ERROR] Kesalahan saat shutdown server: %v", err)
	}

	log.Println("[INFO] Server berhasil dihentikan dengan aman.")
}
