package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	dbconfig "github.com/wattam/shoeStoreDB/gopostgres"
	"github.com/wattam/shoeStoreDB/handler"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func main() {

	// Initializing the connection with the database
	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	// Verifying if the database connection was a success
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	store := shoe.New(db)

	// Creates a instance of the Gin Web Framework
	r := gin.Default()

	// REST API functions
	r.GET("/get", handler.ShoeGetAll(store))
	r.GET("/:id", handler.ShoeGet(store))
	r.POST("/post", handler.ShoeAdd(store))
	r.PUT("/put", handler.ShoeUpdate(store))
	r.DELETE("/:id", handler.ShoeDelete(store))

	r.Run()

	defer db.Close()
}
