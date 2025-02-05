package model

type RequestAddRoles struct{
	Name string `json:"name" validate:"required,min=3,max=10"`
}