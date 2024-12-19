package config

import (
	"github.com/Andesson/APIGO/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgree() (*gorm.DB, error) {
	logger := GetLogger("postgree")
	dsn := "host=postgres user=postgres password=123 dbname=estudos sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgree opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("postgree automigration error: %v", err)
		return nil, err
	}
	return db, nil
}
