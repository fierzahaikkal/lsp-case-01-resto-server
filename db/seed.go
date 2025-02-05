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

    // Insert some Food (Menu)
    foods := []Menu{
        {ID: uuid.New(), Nama: "Nasi Goreng", Deskripsi:"Nasi goreng adalah hidangan khas Indonesia yang terdiri dari nasi yang digoreng dengan bumbu-bumbu seperti bawang putih, kecap, dan cabai. Ditambahkan bahan-bahan lain seperti telur, ayam, atau udang, dan disajikan dengan acar, kerupuk, dan potongan timun.", Stok: 50, Kategori: "menu utama", Harga:30000, URI_image: "https://images.unsplash.com/photo-1512058564366-18510be2db19?q=80&w=2072&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
        {ID: uuid.New(), Nama: "Es Teh", Deskripsi:"Es teh adalah minuman yang terbuat dari teh yang diseduh dan kemudian didinginkan dengan es. Disajikan dengan tambahan gula atau sirup, memberikan rasa yang menyegarkan, terutama di cuaca panas.", Stok: 100, Kategori: "minuman", Harga:10000, URI_image: "https://images.unsplash.com/photo-1683170275059-302acae79168?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTB8fGljZSUyMHRlYXxlbnwwfHwwfHx8MA%3D%3D"},
        {ID: uuid.New(), Nama: "Mie Rebus", Deskripsi:"Mie rebus adalah hidangan berupa mie yang direbus dan disajikan dengan kuah panas yang gurih. Kuahnya bisa berupa kaldu ayam atau udang, dengan tambahan telur, sayuran, dan bawang goreng.", Stok: 50, Kategori: "menu utama", Harga:25000, URI_image: "https://images.unsplash.com/photo-1593179241557-bce1eb92e47e?q=80&w=1974&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"},
        {ID: uuid.New(), Nama: "Ice Cream", Deskripsi:"Ice cream adalah makanan penutup dingin yang dibuat dari campuran susu, krim, gula, dan perasa seperti cokelat, vanila, atau buah-buahan. Teksturnya lembut dan manis, sangat cocok dinikmati di hari yang panas.", Stok: 100, Kategori: "appetizer", Harga:10000, URI_image: "https://images.unsplash.com/photo-1590080962330-747c6aba8028?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTF8fGljZSUyMGNyZWFtfGVufDB8fDB8fHww"},

    }

    // Insert Orders (Pesanan)
    orders := []Pesanan{
        {ID: uuid.New(), StatusOrder: "menunggu", Jumlah: 2, Harga: 20000, JenisTransaksi: "tunai", UserID: customers[0].ID, MenuID: foods[0].ID},
        {ID: uuid.New(), StatusOrder: "dibuat", Jumlah: 1, Harga: 5000, JenisTransaksi: "non tunai", UserID: customers[1].ID, MenuID: foods[1].ID},
        {ID: uuid.New(), StatusOrder: "selesai", Jumlah: 1, Harga: 5000, JenisTransaksi: "tunai", UserID: customers[1].ID, MenuID: foods[2].ID},
        {ID: uuid.New(), StatusOrder: "selesai", Jumlah: 1, Harga: 5000, JenisTransaksi: "non tunai", UserID: customers[0].ID, MenuID: foods[3].ID},

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
