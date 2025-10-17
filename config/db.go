package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	_ = godotenv.Load()

	driver := os.Getenv("DB_DRIVER")
	var err error

	switch driver {
	case "mysql":
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, name)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	case "postgres":
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		name := os.Getenv("DB_NAME")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, pass, name, port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	default:
		name := os.Getenv("DB_NAME")
		if name == "" {
			name = "mybudget.db"
		}
		DB, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}

	log.Println("✅ Database connected successfully!")
}
