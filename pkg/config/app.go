package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  Warning: .env file not found, using system environment variables")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("❌ Failed to connect to database: %v", err))
	}
	db = d
	fmt.Println("✅ Connected to MySQL database successfully!")
}

func GetDB() *gorm.DB {
	return db
}
