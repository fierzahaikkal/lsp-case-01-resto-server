package model

type RequestCreateOrder struct {
	CustomerID  string          `json:"customer_id" validate:"required,uuid"`
    Items       []RequestOrderItem `json:"items" validate:"required,dive"`
    DeliveryAddress string      `json:"delivery_address" validate:"required"`
    Notes       string          `json:"notes"`
}

type RequestOrderItem struct {
	MenuID   string `json:"menu_id" validate:"required,uuid"`
    Quantity int    `json:"quantity" validate:"required,gt=0"`
}