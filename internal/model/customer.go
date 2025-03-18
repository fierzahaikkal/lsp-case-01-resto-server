package model

type RequestSignInCustomer struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type RequestSignUpCustomer struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Nama     string `json:"nama" validate:"required,min=3,max=50"`
	Alamat   string `json:"alamat" validate:"required,max=100"`
	Telepon  string `json:"telepon" validate:"required,max=13"`
	Sandi    string `json:"sandi" validate:"required,min=8,max=30, e164"`
	RoleID   string `json:"role_id" validate:"required"`
}