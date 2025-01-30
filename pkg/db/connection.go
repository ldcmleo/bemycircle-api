package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection(dsn string, gormConfig *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
