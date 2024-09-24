package pkg

import (
	"fmt"

	"github.com/fierzahaikkal/lsp-case-01-resto-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
