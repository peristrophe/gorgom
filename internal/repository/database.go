package repository

import (
	"gorgom/internal/setting"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDBConn() *gorm.DB {
	dsn := os.ExpandEnv(setting.POSTGRES_DSN_TEMPLATE)
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
