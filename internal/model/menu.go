package model

type RequestCreateMenu struct {
	Nama      string    `json:"namae" validate:"required,min=3,max=100"`
	Deskripsi string	`json:"deskripsi" validate:"required"`
	Stok      float32   `json:"stok" validate:"required,gt=0"`
	Harga	  float64	`json:"price" validate:"required,gt=0"`
	Kategori  string    `json:"kategori" validate:"required"`
	URI_image string    `json:"uri_image" validate:"url"`
}

//omitempty => output didnt print if value is zero
type RequestUpdateMenu struct {
	Nama      *string    `json:"namae" validate:"omitempty,min=3,max=100"`
	Deskripsi *string	`json:"deskripsi" validate:"omitempty"`
	Stok      *float32   `json:"stok" validate:"omitempty,gt=0"`
	Harga	  *float64	`json:"price" validate:"omitempty,gt=0"`
	Kategori  *string    `json:"kategori" validate:"omitempty"`
	URI_image *string    `json:"uri_image" validate:"omitempty,url"`
}