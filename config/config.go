package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDbConnection() *gorm.DB {

	return db
}

func InitialiseDbConnection() {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	println("%s\n", os.Getenv("DB"))
	//DBMS := os.Getenv("DB")
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOLNEW")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	//dsn := "root:Vahak@123@tcp(localhost:3306)/test_go?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := USER + ":" + PASS + "@tcp(" + PROTOCOL + ":3306)" + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := USER + ":" + PASS + "@tcp(" + PROTOCOL + ")" + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
