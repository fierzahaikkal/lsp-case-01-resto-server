package main

import (
	"log"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/config"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/db"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/admin"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/customer"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/makanan"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/pesanan"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	cfg := config.LoadConfig()

	dbConn, err := pkg.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 60 * 1000, // Limit 20 requests per minute
	}))

    // Run migrations
    if err := db.Migrate(dbConn); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Seed data
    if err := db.Seed(dbConn); err != nil {
        log.Fatalf("Failed to seed database: %v", err)
    }

	// Repositories
	adminRepo := admin.NewAdminRepository(dbConn)
	customerRepo := customer.NewCustomerRepository(dbConn)
	makananRepo := makanan.NewMakananRepository(dbConn)
	pesananRepo := pesanan.NewPesananRepository(dbConn)

	// Services
	adminService := admin.NewAdminService(adminRepo)
	customerService := customer.NewCustomerService(customerRepo)
	makananService := makanan.NewMakananService(makananRepo)
	pesananService := pesanan.NewPesananService(pesananRepo)

	// Handlers
	adminHandler := admin.NewAdminHandler(adminService)
	customerHandler := customer.NewCustomerHandler(customerService)
	makananHandler := makanan.NewMakananHandler(makananService)
	pesananHandler := pesanan.NewPesananHandler(pesananService)

	// Admin routes
	app.Post("/admins", adminHandler.CreateAdmin)
	app.Get("/admins/:id", adminHandler.GetAdmin)
	app.Put("/admins/:id", adminHandler.UpdateAdmin)
	app.Delete("/admins/:id", adminHandler.DeleteAdmin)

	// Customer routes
	app.Post("/customers", customerHandler.CreateCustomer)
	app.Get("/customers/:id", customerHandler.GetCustomer)
	app.Put("/customers/:id", customerHandler.UpdateCustomer)
	app.Delete("/customers/:id", customerHandler.DeleteCustomer)

	// Makanan routes
	app.Post("/makanans", makananHandler.CreateMakanan)
	app.Get("/makanans/:id", makananHandler.GetMakanan)
	app.Put("/makanans/:id", makananHandler.UpdateMakanan)
	app.Delete("/makanans/:id", makananHandler.DeleteMakanan)

	// Pesanan routes
	app.Post("/pesanans", pesananHandler.CreatePesanan)
	app.Get("/pesanans/:id", pesananHandler.GetPesanan)
	app.Put("/pesanans/:id", pesananHandler.UpdatePesanan)
	app.Delete("/pesanans/:id", pesananHandler.DeletePesanan)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
