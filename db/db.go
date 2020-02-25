package db

import (
	"fmt"
	"log"
	"os"

	"../model"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// Postgres Driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

// Init データべースの初期化
func Init() {
	// .envの読み込み。失敗したら落ちる
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		log.Fatal(loadEnv.Error())
	}

	// DB接続
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"), os.Getenv("SSLMODE"))
	db, err = gorm.Open(os.Getenv("DB"), connection)

	// DB接続に失敗したら落ちる
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Freq{})
	db.LogMode(true)

}

// Get データべデータのインスタンスを渡す
func Get() *gorm.DB {
	if db == nil {
		Init()
	}
	return db
}

// Close データベースから切断する
func Close() {
	db.Close()
}
