package database

import (
	"fmt"
	"os"
	"seriesmanager-services/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() *gorm.DB {

	errEnv := godotenv.Load()

	if errEnv != nil {
		panic(errEnv.Error())
	}
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, pass, name)

	db, errDb := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errDb != nil {
		panic(errDb.Error())
	}

	if errMigrate := db.AutoMigrate(&models.User{}, &models.Series{}, &models.Season{}); errMigrate != nil {
		panic(errMigrate)
	}
	return db
}

func Close(db *gorm.DB) {
	dbSql, errDb := db.DB()

	if errDb != nil {
		panic(errDb.Error())
	}
	if errClose := dbSql.Close(); errClose != nil {
		panic(errClose.Error())
	}
}

// TODO: delete with userId each time
