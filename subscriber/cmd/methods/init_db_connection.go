package methods

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDbConnection() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=qwerty dbname=postgres port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
