package database

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/wattam/shoe-store-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "postgres123"

const DbName = "ShoeStore"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

var db *gorm.DB

func CreateDatabase() {

	dbGorm, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DATABASE CONNECTED!")
	}

	db = dbGorm

	RunMigrations()
}

func GetDatabase() *gorm.DB {
	return db
}

func RunMigrations() {
	db.AutoMigrate(models.Shoe{})
}

func CloseConnection() {
	config, _ := db.DB()

	config.Close()
}
