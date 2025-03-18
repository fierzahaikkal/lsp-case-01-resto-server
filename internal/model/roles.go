package model

type RequestAddRoles struct{
	Name string `json:"name" validate:"required,min=3,max=10"`
	Level int `json:"level" validate:"required,min=1,max=3"`
}

type RequestUpdateRoles struct{
	Name string `json:"name" validate:"required,min=3,max=10"`
	Level int `json:"level" validate:"required,min=1,max=3"`
}

