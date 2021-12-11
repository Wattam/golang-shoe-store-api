package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wattam/shoeStoreDB/gopostgres/dbconfig"
	"github.com/wattam/shoeStoreDB/handler"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func main() {

	db, err := sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	store := shoe.New(db)

	r := gin.Default()

	r.GET("/get", handler.ShoeGetAll(store))
	//r.GET("/:id", handler.ShoeGet(store))
	r.POST("/post", handler.ShoeAdd(store))
	r.PUT("/put", handler.ShoeUpdate(store))
	//r.DELETE("/:id", handler.ShoeDelete(store))

	r.Run()

	defer db.Close()
}
