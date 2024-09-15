package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ronb.co/project/models"
)

type Database struct{}

func InitDB() {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASS")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgDB := os.Getenv("POSTGRES_DB")
	pgSSL := "disable"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		pgHost,
		pgUser,
		pgPass,
		pgDB,
		pgPort,
		pgSSL,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.Store{},
		&models.Visits{},
		&models.Points{},
		&models.Rewards{},
		&models.RewardItem{},
	)
	// &Customer{}, &Store{}, &Visits{}, &Points{}, &Rewards{})
}
