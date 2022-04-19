package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv)
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, name)

	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic(errDb)
	}
	return db
}

func Close(db *gorm.DB) {
	dbSql, errDb := db.DB()

	if errDb != nil {
		panic(errDb)
	}
	errClose := dbSql.Close()

	if errClose != nil {
		panic(errClose)
	}
}
