## REST API for a shoe store using Golang, GORM and Gin Web Framework

This API uses a local PostgreSQL database that's set through the `/database/databaseConfig.go` file.


### To run the project
- Once you are in the `/main` folder run the command:
```
$ go run main.go
```
- Or run directly the command:
```
$ go run main/main.go
```
---
### To get all the dependencies
- Gin Web Framework:
```
$ go get -u github.com/gin-gonic/gin
```
- GORM (ORM library):
```
$ go get -u gorm.io/gorm
```
- Go Postgres driver:
```
$ go get github.com/lib/pq
```
---
### To syncronize all the dependencies
```
$ export GO111MODULE=on
```
```
$ go mod tidy
```
---
### Packages
- `package main`: router creation and calls for GET, POST, PUT and DELETE requests.
- `package models`: data model for a shoe.
- `package database`: database configuration and initialization.
- `package shoeHandlers`: functions to perform HTTP requests.
---
### API Requests
#### `GetAllShoes Request`

- `GET` `localhost:8080/shoes/get`: gets all shoes from the database.
- Response example:
```json
[
	{
		"id": "1",
		"name": "Name 1",
		"style": "Style 1",
		"colour": "Colour 1",
		"material": "Material 1",
		"price": "1"
	},
	{
		"id": "2",
		"name": "Name 2",
		"style": "Style 2",
		"colour": "Colour 2",
		"material": "Material 2",
		"price": "2"
	}
]
```

---
#### `GetShoe Request`

- `GET` `localhost:8080/shoes/:id`: gets a specific shoe from the database using the ID.
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"style": "Style 1",
	"colour": "Colour 1",
	"material": "Material 1",
	"price": "1"
}
```

---
#### `CreateShoe Request`

- `POST` `localhost:8080/shoes/post`: creates a shoe in the database passing the parameters through the body.
- Body example:
```json
{
	"name":"Shoe 1",
	"style":"Style 1",
	"colour":"Colour 1",
	"material":"Material 1",
	"price":"1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Name 1",
	"style": "Style 1",
	"colour": "Colour 1",
	"material": "Material 1",
	"price": "1"
}
```

---
#### `EditShoe Request`

- `PUT` `localhost:8080/shoes/put`: edits a shoe in the database with the ID equal to the one passed on the body.
- Body example:
```json
{
	"id":"1",
	"name":"Shoe 1 updated",
	"style":"Style 1 updated",
	"colour":"Color 1 updated",
	"material":"Material 1 updated",
	"price":"1.1"
}
```
- Response example:
```json
{
	"id": "1",
	"name": "Shoe 1 updated",
	"style": "Style 1 updated",
	"colour": "Color 1 updated",
	"material": "Material 1 updated",
	"price": "1.1"
}
```

---
#### `DeleteShoe Request`

- `DELETE` `localhost:8080/shoes/:id`: deletes a shoe from the database using the ID.
---
