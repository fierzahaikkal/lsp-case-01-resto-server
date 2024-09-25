package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Seed initial data into the tables
func Seed(db *gorm.DB) error {
    // Insert Roles (Peranan)
    roles := []Peranan{
        {ID: uuid.New(), Peran: "admin"},
        {ID: uuid.New(), Peran: "customer"},
    }

    // Insert some Admins
    admins := []Admin{
        {ID: uuid.New(), Username: "admin1", Password: "pass1234", RoleID: roles[0].ID},
        {ID: uuid.New(), Username: "admin2", Password: "pass1234", RoleID: roles[0].ID},
    }

    // Insert some Customers
    customers := []Kustomer{
        {ID: uuid.New(), Nama: "Customer One", Alamat: "123 Address", Telepon: "081234567890", Username: "cust1", Sandi: "password1", RoleID: roles[1].ID},
        {ID: uuid.New(), Nama: "Customer Two", Alamat: "456 Address", Telepon: "081234567891", Username: "cust2", Sandi: "password2", RoleID: roles[1].ID},
    }

    // Insert some Food (Makanan)
    foods := []Makanan{
        {ID: uuid.New(), Nama: "Nasi Goreng", Stok: 50, Kategori: "makanan utama"},
        {ID: uuid.New(), Nama: "Es Teh", Stok: 100, Kategori: "minuman"},
    }

    // Insert Orders (Pesanan)
    orders := []Pesanan{
        {ID: uuid.New(), StatusOrder: "menunggu", Jumlah: 2, Harga: 20000, JenisTransaksi: "tunai", UserID: customers[0].ID, MakananID: foods[0].ID},
        {ID: uuid.New(), StatusOrder: "dibuat", Jumlah: 1, Harga: 5000, JenisTransaksi: "non tunai", UserID: customers[1].ID, MakananID: foods[1].ID},
    }

    // Insert into database
    if err := db.Create(&roles).Error; err != nil {
        return err
    }
    if err := db.Create(&admins).Error; err != nil {
        return err
    }
    if err := db.Create(&customers).Error; err != nil {
        return err
    }
    if err := db.Create(&foods).Error; err != nil {
        return err
    }
    if err := db.Create(&orders).Error; err != nil {
        return err
    }

    return nil
}
