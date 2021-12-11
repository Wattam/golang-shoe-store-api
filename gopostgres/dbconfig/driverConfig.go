package dbconfig

import "fmt"

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "postgres123"

const DbName = "ShoeStore"

const TableName = "shoe"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

// Data
type Shoe struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Style    string  `json:"style"`
	Color    string  `json:"color"`
	Material string  `json:"material"`
	Price    float64 `json:"price"`
}
