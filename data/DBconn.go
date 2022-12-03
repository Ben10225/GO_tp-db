package data

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// const DB_USERNAME = "root"
// const DB_PASSWORD = "root456789"
// const DB_NAME = "go_taipei"
// const DB_HOST = "127.0.0.1"
// const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := envGet("DB_USERNAME") + ":" + envGet("DB_PASSWORD") + "@tcp" + "(" + envGet("DB_HOST") + ":" + envGet("DB_PORT") + ")/" + envGet("DB_NAME") + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)

	}

	return db
}

func envGet(s string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(s)
}
