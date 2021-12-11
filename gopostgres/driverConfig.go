package dbconfig

import "fmt"

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5432"

const Password = "postgres123"

const DbName = "ShoeStore"

const TableName = "shoe"

// String that will be used on main package to initialize the database with the function sq.Open()
var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
