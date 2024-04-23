package lib

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	defaultDbHost = "127.0.0.1"
	defaultDbPort = "5432"
	defaultDbName = "postgres"
	defaultDbUser = "postgres"
	defaultDbPass = "postgres"
)

func GetSQLDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("godotenv.Load error:", err)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = defaultDbHost
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = defaultDbPort
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = defaultDbName
	}
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = defaultDbUser
	}
	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		dbPass = defaultDbPass
	}

	uri := fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=require TimeZone=Asia/Tokyo", dbUser, dbPass, dbName, dbHost, dbPort)

	dbRootCert := os.Getenv("DB_ROOT_CERT")
	if dbRootCert != "" {
		uri += fmt.Sprintf(" sslrootcert=%s", dbRootCert)
	}

	dbCert := os.Getenv("DB_CERT")
	dbKey := os.Getenv("DB_KEY")
	if dbRootCert != "" && dbCert != "" && dbKey != "" {
		uri += fmt.Sprintf(" sslcert=%s sslkey=%s", dbCert, dbKey)
	}

	db, err := sql.Open("pgx", uri)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %s", err)
	}

	return db, err
}

func GetGORMDB(db *sql.DB) (*gorm.DB, error) {
	logMode := logger.Silent
	if os.Getenv("MODE") == "DEBUG" {
		logMode = logger.Info
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, fmt.Errorf("gorm.Open error: %s", err)
	}

	return gormDB, nil
}
