package model

type RequestCreatePayment struct {
	OrderID string `json:"orderd_id" validate:"required,uuid"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
	PaymentMethod string `json:"payment_method" validate:"required,oneof=credit_card bank_transfer e_wallet"`
}