package db

import (
	"cinema/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(host, user, password, dbname string, port int) *GormDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return &GormDB{DB: db}
}

func (g *GormDB) Migrate() {
	if err := g.DB.AutoMigrate(
		&model.Cinemas{},
		&model.Schedule{},
		&model.Film{},
		&model.User{},
		&model.Seat{},
		&model.Booking{},
		&model.Room{},
		&model.Token{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
