package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCustomerIntegration(t *testing.T) {
	app := fiber.New()
	customerID := uuid.New()
	roleID := uuid.New()
	t.Run("Create Customer", func(t *testing.T) {
		customer := model.RequestSignUpCustomer{
			Nama:     "Test Customer",
			Alamat:   "Test Address",
			Telepon:  "08123456789",
			Email:    "test@test.com",
			Username: "testuser",
			Sandi:    "password123",
			RoleID:     roleID.String(),
		}

		jsonCustomer, _ := json.Marshal(customer)
		req := httptest.NewRequest(http.MethodPost, "/api/customers", bytes.NewBuffer(jsonCustomer))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Get Customer By ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/customers/"+customerID.String(), nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestMenuIntegration(t *testing.T) {
	app := fiber.New()
	menuID := uuid.New()

	t.Run("Create Menu", func(t *testing.T) {
		menu := model.RequestCreateMenu{
			Nama:      "Test Menu",
			Deskripsi: "Test Description",
			Stok:      10,
			Harga:     50000,
			Kategori:  "menu utama",
			URI_image: "test.jpg",
		}

		jsonMenu, _ := json.Marshal(menu)
		req := httptest.NewRequest(http.MethodPost, "/api/menus", bytes.NewBuffer(jsonMenu))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Get Menu By ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/menus/"+menuID.String(), nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestOrderIntegration(t *testing.T) {
	app := fiber.New()
	orderID := uuid.New()
	customerID := uuid.New()
	menuID := uuid.New()

	t.Run("Create Order", func(t *testing.T) {
		order := model.RequestCreateOrder{
			CustomerID:      customerID.String(),
			Items: []model.RequestOrderItem{
				{
					MenuID:   menuID.String(),
					Quantity: 2,
				},
			},
			DeliveryAddress: "Test Address",
			Notes:          "Test Notes",
		}

		jsonOrder, _ := json.Marshal(order)
		req := httptest.NewRequest(http.MethodPost, "/api/orders", bytes.NewBuffer(jsonOrder))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Get Order By ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/orders/"+orderID.String(), nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}

func TestRolesIntegration(t *testing.T) {
	app := fiber.New()
	roleID := uuid.New()

	t.Run("Create Role", func(t *testing.T) {
		role := model.RequestAddRoles{
			Name: "Test Role",
		}

		jsonRole, _ := json.Marshal(role)
		req := httptest.NewRequest(http.MethodPost, "/api/roles", bytes.NewBuffer(jsonRole))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("Get Role By ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/roles/"+roleID.String(), nil)
		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
