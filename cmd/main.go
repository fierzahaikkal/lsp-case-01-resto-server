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
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Initialize default config
app.Use(cors.New())

// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "https://gofiber.io, https://gofiber.net",
    AllowHeaders: "Origin, Content-Type, Accept",
}))

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
	app.Post("/api/v1/admin", adminHandler.CreateAdmin)
	app.Get("/api/v1/admin/:id", adminHandler.GetAdmin)
	app.Put("/api/v1/admin/:id", adminHandler.UpdateAdmin)
	app.Delete("/api/v1/admin/:id", adminHandler.DeleteAdmin)

	// Customer routes
	app.Post("/api/v1/customer", customerHandler.CreateCustomer)
	app.Get("/api/v1/customer/:id", customerHandler.GetCustomer)
	app.Put("/api/v1/customer/:id", customerHandler.UpdateCustomer)
	app.Delete("/api/v1/customer/:id", customerHandler.DeleteCustomer)

	// Menu routes
	app.Get("/api/v1/menu", menuHandler.GetMenu)
	app.Post("/api/v1/menu", menuHandler.CreateMenu)
	app.Get("/api/v1/menu/:id", menuHandler.GetMenuByID)
	app.Put("/api/v1/menu/:id", menuHandler.UpdateMenu)
	app.Delete("/api/v1/menu/:id", menuHandler.DeleteMenu)

	// Pesanan routes
	app.Get("/api/v1/pesanan", pesananHandler.GetPesanan)
	app.Post("/api/v1/pesanan", pesananHandler.CreatePesanan)
	app.Get("/api/v1/pesanan/:id", pesananHandler.GetPesananByID)
	app.Get("/api/v1/pesanan/cetak", pesananHandler.CetakPesanan)
	app.Get("/api/v1/pesanan/cetak/:id", pesananHandler.CetakPesananByID)
	app.Put("/api/v1/pesanan/:id", pesananHandler.UpdatePesanan)
	app.Delete("/api/v1/pesanan/:id", pesananHandler.DeletePesanan)

	// Start the server
	log.Fatal(app.Listen(":8000"))
}
