package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func getConfig() DBConfig {
	return DBConfig{
		Host:     "postgres",
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     "5432",
	}
}

func connectDB(config DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func GetDB() *gorm.DB {
	dbOnce.Do(func() {
		config := getConfig()
		db, err := connectDB(config)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		dbInstance = db
	})

	return dbInstance
}

func CloseDB() error {
	sqlDB, err := dbInstance.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
