package setting

const (
	POSTGRES_DSN_TEMPLATE = "host=${GORGOM_DB_HOST} " +
		"user=${GORGOM_DB_USER} " +
		"password=${GORGOM_DB_PASSWORD} " +
		"dbname=${GORGOM_DB_NAME} " +
		"port=${GORGOM_DB_PORT} " +
		"sslmode=disable " +
		"TimeZone=Asia/Tokyo"
)
