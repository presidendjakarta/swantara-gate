package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/presidendjakarta/swantara-gate/internal/config"
	"github.com/presidendjakarta/swantara-gate/internal/database"
	"github.com/presidendjakarta/swantara-gate/internal/handler"
	"github.com/presidendjakarta/swantara-gate/internal/middleware"
	"github.com/presidendjakarta/swantara-gate/internal/repository"
	"github.com/presidendjakarta/swantara-gate/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Memuat environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ File .env tidak ditemukan, menggunakan environment default")
	}

	// Memuat konfigurasi
	cfg := config.LoadConfig()

	log.Println("🚀 Memulai Swantara Gate API Gateway...")

	// Inisialisasi database
	if err := database.InitDatabase(cfg.DatabasePath, cfg.DatabaseSQLPath); err != nil {
		log.Fatalf("❌ Gagal menginisialisasi database: %v", err)
	}

	// Memastikan database ditutup saat aplikasi berhenti
	defer database.CloseDatabase()

	// Inisialisasi repositories
	userRepo := repository.NewUserRepository(database.GetDB())
	consumerRepo := repository.NewAPIConsumerRepository(database.GetDB())
	hostRepo := repository.NewHostRepository(database.GetDB())
	vhostRepo := repository.NewVirtualHostRepository(database.GetDB())

	// Inisialisasi services
	userService := service.NewUserService(userRepo)
	consumerService := service.NewAPIConsumerService(consumerRepo)
	hostService := service.NewHostService(hostRepo)
	vhostService := service.NewVirtualHostService(vhostRepo)

	// Inisialisasi handlers
	userHandler := handler.NewUserHandler(userService)
	consumerHandler := handler.NewAPIConsumerHandler(consumerService)
	hostHandler := handler.NewHostHandler(hostService)
	vhostHandler := handler.NewVirtualHostHandler(vhostService)

	// Setup Admin Router
	adminMux := setupAdminRouter(userHandler, consumerHandler, hostHandler, vhostHandler)

	// Menjalankan Admin HTTP Server
	go func() {
		addr := ":" + toString(cfg.AdminHTTPPort)
		log.Printf("🌐 Admin HTTP Server berjalan di %s", addr)
		if err := http.ListenAndServe(addr, adminMux); err != nil {
			log.Fatalf("❌ Admin HTTP Server error: %v", err)
		}
	}()

	// TODO: Jalankan Admin HTTPS Server (perlu sertifikat SSL)
	// go func() {
	// 	addr := ":" + toString(cfg.AdminHTTPSPort)
	// 	log.Printf("🔒 Admin HTTPS Server berjalan di %s", addr)
	// 	if err := http.ListenAndServeTLS(addr, cfg.AdminSSLCertPath, cfg.AdminSSLKeyPath, adminMux); err != nil {
	// 		log.Fatalf("❌ Admin HTTPS Server error: %v", err)
	// 	}
	// }()

	// TODO: Jalankan Proxy HTTP & HTTPS Servers
	// Ini akan diimplementasikan setelah admin API selesai

	log.Println("✅ Swantara Gate API Gateway siap digunakan!")
	log.Println("📍 Admin Panel: http://localhost:" + toString(cfg.AdminHTTPPort))
	
	// Blocking agar program tidak selesai
	select {}
}

// setupAdminRouter mengatur routes untuk Admin API
func setupAdminRouter(
	userHandler *handler.UserHandler,
	consumerHandler *handler.APIConsumerHandler,
	hostHandler *handler.HostHandler,
	vhostHandler *handler.VirtualHostHandler,
) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Swantara Gate Admin API is running"}`))
	})

	// Routes untuk Users
	mux.HandleFunc("POST /api/admin/users", userHandler.CreateUser)
	mux.HandleFunc("GET /api/admin/users", userHandler.GetAllUsers)
	mux.HandleFunc("GET /api/admin/users/{id}", userHandler.GetUserByID)
	mux.HandleFunc("PUT /api/admin/users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /api/admin/users/{id}", userHandler.DeleteUser)

	// Routes untuk API Consumers
	mux.HandleFunc("POST /api/admin/consumers", consumerHandler.CreateConsumer)
	mux.HandleFunc("GET /api/admin/consumers", consumerHandler.GetAllConsumers)
	mux.HandleFunc("GET /api/admin/consumers/{id}", consumerHandler.GetConsumerByID)
	mux.HandleFunc("PUT /api/admin/consumers/{id}", consumerHandler.UpdateConsumer)
	mux.HandleFunc("DELETE /api/admin/consumers/{id}", consumerHandler.DeleteConsumer)

	// Routes untuk Hosts
	mux.HandleFunc("POST /api/admin/hosts", hostHandler.CreateHost)
	mux.HandleFunc("GET /api/admin/hosts", hostHandler.GetAllHosts)
	mux.HandleFunc("GET /api/admin/hosts/{id}", hostHandler.GetHostByID)
	mux.HandleFunc("PUT /api/admin/hosts/{id}", hostHandler.UpdateHost)
	mux.HandleFunc("DELETE /api/admin/hosts/{id}", hostHandler.DeleteHost)

	// Routes untuk Virtual Hosts
	mux.HandleFunc("POST /api/admin/virtual-hosts", vhostHandler.CreateVirtualHost)
	mux.HandleFunc("GET /api/admin/virtual-hosts", vhostHandler.GetAllVirtualHosts)
	mux.HandleFunc("GET /api/admin/virtual-hosts/{id}", vhostHandler.GetVirtualHostByID)
	mux.HandleFunc("PUT /api/admin/virtual-hosts/{id}", vhostHandler.UpdateVirtualHost)
	mux.HandleFunc("DELETE /api/admin/virtual-hosts/{id}", vhostHandler.DeleteVirtualHost)

	// Menggunakan middleware
	return middleware.LoggingMiddleware(middleware.CORSMiddleware(mux))
}

// toString mengubah int ke string
func toString(i int) string {
	return strconv.Itoa(i)
}
