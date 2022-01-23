package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type ProviderGorm struct {
	connection *gorm.DB
	dialector  gorm.Dialector
	config     *gorm.Config
}

func NewProviderGorm() *ProviderGorm {
	db := new(ProviderGorm)
	err := db.init()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func (g *ProviderGorm) Connection() (*gorm.DB, error) {
	if g.connection == nil {
		return nil, errors.New("database not init")
	}
	return g.connection, nil
}

func (g *ProviderGorm) init() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))
	g.dialector = postgres.Open(dsn)
	g.config = &gorm.Config{}
	g.connection, err = gorm.Open(g.dialector, g.config)
	return err
}
