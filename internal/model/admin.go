package model

type RequestSignInAdmin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=30"`
}

type RequestSignUpAdmin struct {
	Username         string `json:"username" validate:"required,min=3,max=30"`
	Email            string `json:"email" validate:"required,email"`
	Nama             string `json:"nama" validate:"required,min=3,max=30"`
	Password         string `json:"password" validate:"required,min=8,max=30"`
	ValidatePassword string `json:"validate_password" validate:"required"`
}

type RequestUpdateAdmin struct {
	Username *string `json:"username" validate:"required,min=3,max=30"`
	Email    *string `json:"email" validate:"required,email"`
	Nama     *string `json:"nama" validate:"required,min=3,max=30"`
	Password *string `json:"password" validate:"required,min=8,max=30"`
}
