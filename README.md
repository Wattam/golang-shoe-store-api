# Shoe Store API
![Shoe Store Logo](img/../assets/ShoeStore.png)

## Description

Restful API for a shoe store. It contains a implementation of Shoes with the following operations:

- Get all shoes.
- Get a shoe.
- Create a shoe.
- Update a shoe.
- Delete a shoe.

## Technologies used
- Go.
- Gin Web Framework.
- GORM.
- PostgreSQL.

## To run the project
- Once you are in the `/main` folder run the command:
```
go run main.go
```
- Or run directly the command:
```
go run main/main.go
```

## To get dependencies
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

## To syncronize dependencies
```
export GO111MODULE=on
```
```
go mod tidy
```


## Models
- ### `Shoe`

| Attribute |   Type  |
|:---------:|:-------:|
|     id    |   int   |
|    name   |  string |
|   color   |  string |
|   price   | float64 |



## Shoe Requests
### `Index Shoes`

- `GET` `localhost:8080/api/v1/shoes`: indexes shoes.
- Response example:

**`200 OK`**
```json
[
	{
		"id": "1",
		"name": "Nike SB",
		"color": "Black",
		"price": "150"
	},
	{
		"id": "2",
		"name": "Adidas Pictoris",
		"color": "White",
		"price": "200"
	}
]
```

---
### `Show Shoe`

- `GET` `localhost:8080/api/v1/shoes/{id}`: shows a shoe by the ID.
- Response example:

**`200 OK`**
```json
{
	"id": "1",
	"name": "Nike SB",
	"color": "Black",
	"price": "150"
}
```

---
### `Store Shoe`

- `POST` `localhost:8080/api/v1/shoes`: stores a shoe getting it's attributes throught the JSON body.
- Body example:
```json
{
	"name": "Nike SB",
	"color": "Black",
	"price": "150.00"
}
```
- Response example:

**`201 Created`**
```json
{
	"id": "1",
	"name": "Nike SB",
	"color": "Black",
	"price": "150"
}
```

---
### `Update Shoe`

- `PUT` `localhost:8080/api/v1/shoes/{id}`: updates a shoe by the ID getting it's attributes throught the JSON body.
- Body example:
```json
{
	"name": "Nike SB",
	"color": "White",
	"price": "180.00"
}
```
- Response example:

**`200 OK`**
```json
{
	"id": "1",
	"name": "Nike SB",
	"color": "White",
	"price": "180"
}
```

---
### `DeleteShoe Request`

- `DELETE` `localhost:8080/api/v1/shoes/{id}`: deletes a shoe by the ID.
- Response example: **`204 No Content`**
---