# Shoe Store API
![Shoe Store Logo](img/../assets/ShoeStore.png)

## REST API for a shoe store using Golang, GORM and Gin Web Framework

This API uses a local PostgreSQL database that's set through the `/database/databaseConfig.go` file.


### To run the project
- Once you are in the `/main` folder run the command:
```
go run main.go
```
- Or run directly the command:
```
go run main/main.go
```
---
### To get all the dependencies
- Gin Web Framework:
```
go get -u github.com/gin-gonic/gin
```
- GORM (ORM library):
```
go get -u gorm.io/gorm
```
- Go Postgres driver:
```
go get github.com/lib/pq
```
---
### To syncronize all the dependencies
```
export GO111MODULE=on
```
```
go mod tidy
```
---
### Packages
- `package main`: router creation and calls for GET, POST, PUT and DELETE requests.
- `package model`: data model for a shoe.
- `package database`: database configuration and initialization.
- `package controller`: shoe controller with functions to perform HTTP requests.
---
### Models
- #### `Shoe`

| Attribute |   Type  |
|:---------:|:-------:|
|     id    |   uint  |
|    name   |  string |
|   color   |  string |
|   price   | float64 |



---
### API Requests
#### `GetAllShoes Request`

- `GET` `localhost:8080/api/v1/shoes/get`: gets all shoes from the database.
- Response example:
```json
[
	{
		"id": "1",
		"name": "Name 1",
		"color": "Color 1",
		"price": "1"
	},
	{
		"id": "2",
		"name": "Name 2",
		"color": "Color 2",
		"price": "2"
	}
]
```

---
#### `GetShoe Request`

- `GET` `localhost:8080/api/v1/shoes/{id}`: gets a specific shoe from the database using the ID.
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"color": "Color 1",
	"price": "1"
}
```

---
#### `PostShoe Request`

- `POST` `localhost:8080/api/v1/shoes/post`: creates a shoe in the database passing the parameters through the body.
- Body example:
```json
{
	"name":"Shoe 1",
	"color":"Color 1",
	"price":"1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"color": "Color 1",
	"price": "1"
}
```

---
#### `PutShoe Request`

- `PUT` `localhost:8080/api/v1/shoes/put`: edits a shoe in the database with the ID equal to the one passed on the body.
- Body example:
```json
{
	"id":"1",
	"name":"Shoe 1 updated",
	"color":"Color 1 updated",
	"price":"1.1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Shoe 1 updated",
	"color": "Color 1 updated",
	"price": "1.1"
}
```

---
#### `DeleteShoe Request`

- `DELETE` `localhost:8080/api/v1/shoes/{id}`: deletes a shoe from the database using the ID.
---