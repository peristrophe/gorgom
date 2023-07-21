package setting

import "os"

var (
	TOKEN_SECRET_KEY = os.Getenv("TOKEN_SECRET_KEY")
	APP_HOST         = os.Getenv("APP_HOST")
)

const (
	POSTGRES_DSN_TEMPLATE = "host=${GORGOM_DB_HOST} " +
		"user=${GORGOM_DB_USER} " +
		"password=${GORGOM_DB_PASSWORD} " +
		"dbname=${GORGOM_DB_NAME} " +
		"port=${GORGOM_DB_PORT} " +
		"sslmode=disable " +
		"TimeZone=Asia/Tokyo"
	TOKEN_EXPIRE = 24 // hour
)
