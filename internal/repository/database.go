package repository

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConn() (*gorm.DB, error) {
	dsn := os.ExpandEnv("host=${GORGOM_DB_HOST} user=${GORGOM_DB_USER} password=${GORGOM_DB_PASSWORD} dbname=${GORGOM_DB_NAME} port=${GORGOM_DB_PORT} sslmode=disable TimeZone=Asia/Tokyo")
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
