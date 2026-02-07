package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgres(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
