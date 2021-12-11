## REST API for a shoe store using Go and Gin Web Framework

This API uses a local PostgreSQL database that's set through the `/gopostgres/driverConfig.go` file.


### To run the project
Once you are in the `/main` folder run the command:
```
$ go run main.go
```
---
### To get all the dependencies
- Gin Web Framework:
```
$ go get -u github.com/gin-gonic/gin
```
- Go Postgres driver:
```
$ go get github.com/lib/pq
```
---
### Packages
- `package main`: database initialization the call for GET, POST, PUT and DELETE functions.
- `package shoe`: data model for a shoe and the functions to get, add, update or delete it from the database.
- `package dbconfig`: database parameters that will be used on it's initialization.
- `package handler`: functions that will use Gin Web Framework to handle JSON activities.
---
### API Authentication
#### `GetAll Request`

- `GET` `localhost:8080/get`: gets all shoes from the database.
- Response example:
```json
[
	{
		"id": 1,
		"name": "Shoe 1",
		"style": "Style 1",
		"colour": "Colour 1",
		"material": "Material 1",
		"price": 1
	},
	{
		"id": 2,
		"name": "Shoe 2",
		"style": "Style 2",
		"colour": "Colour 2",
		"material": "Material 2",
		"price": 2
	}
]
```
---
#### `GetShoe Request`

- `GET` `localhost:8080/:id`: gets a specific shoe from the database using the ID.
- Response example:
```json
{
	"id": 1,
	"name": "Shoe 1",
	"style": "Style 1",
	"colour": "Colour 1",
	"material": "Material 1",
	"price": 1
}
```
---
#### `AddShoe Request`

- `POST` `localhost:8080/post`: adds a shoe to the database passing the parameters through the body.
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
---
#### `UpdateShoe Request`

- `PUT` `localhost:8080/put`: updates a shoe from the database with the ID equal to the one passed on the body.
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
---
#### `DeleteShoe Request`

- `DELETE` `localhost:8080/:id`: deletes a specific shoe from the database using the ID.
---
