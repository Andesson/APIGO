package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializePostgree()
	if err != nil {
		return fmt.Errorf("error initializing postgree: %v", err)
	}
	return nil
}

func GetPostgree() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
