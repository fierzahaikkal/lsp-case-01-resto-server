package main

import (
	"flag"
	"log"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/config"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/db"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/admin"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/customer"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/menu"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/pesanan"
	"github.com/fierzahaikkal/lsp-case-01-resto-server/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	cfg := config.LoadConfig()

	doSeed := flag.Bool("seed", false, "Seed the database with initial data")
	flag.Parse()

	dbConn, err := pkg.InitDB(cfg)
	if err != nil {
		panic(err)
	}

	if *doSeed {
		if err := db.Seed(dbConn); err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
		log.Println("Database seeded successfully")
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

	// Repositories
	adminRepo := admin.NewAdminRepository(dbConn)
	customerRepo := customer.NewCustomerRepository(dbConn)
	menuRepo := menu.NewMenuRepository(dbConn)
	pesananRepo := pesanan.NewPesananRepository(dbConn)

	// Services
	adminService := admin.NewAdminService(adminRepo)
	customerService := customer.NewCustomerService(customerRepo)
	menuService := menu.NewMenuService(menuRepo)
	pesananService := pesanan.NewPesananService(pesananRepo)

	// Handlers
	adminHandler := admin.NewAdminHandler(adminService)
	customerHandler := customer.NewCustomerHandler(customerService)
	menuHandler := menu.NewMenuHandler(menuService)
	pesananHandler := pesanan.NewPesananHandler(pesananService)

	// Admin routes
	app.Post("/admin", adminHandler.CreateAdmin)
	app.Get("/admin/:id", adminHandler.GetAdmin)
	app.Put("/admin/:id", adminHandler.UpdateAdmin)
	app.Delete("/admin/:id", adminHandler.DeleteAdmin)

	// Customer routes
	app.Post("/customer", customerHandler.CreateCustomer)
	app.Get("/customer/:id", customerHandler.GetCustomer)
	app.Put("/customer/:id", customerHandler.UpdateCustomer)
	app.Delete("/customer/:id", customerHandler.DeleteCustomer)

	// Menu routes
	app.Post("/menu", menuHandler.CreateMenu)
	app.Get("/menu/:id", menuHandler.GetMenu)
	app.Put("/menu/:id", menuHandler.UpdateMenu)
	app.Delete("/menu/:id", menuHandler.DeleteMenu)

	// Pesanan routes
	app.Get("/pesanan", pesananHandler.GetPesanan)
	app.Post("/pesanan", pesananHandler.CreatePesanan)
	app.Get("/pesanan/:id", pesananHandler.GetPesanan)
	app.Get("/pesanan/cetak", pesananHandler.CetakPesanan)
	app.Get("/pesanan/cetak/:id", pesananHandler.CetakPesananByID)
	app.Put("/pesanan/:id", pesananHandler.UpdatePesanan)
	app.Delete("/pesanan/:id", pesananHandler.DeletePesanan)

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
