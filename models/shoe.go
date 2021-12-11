package shoe

import (
	"database/sql"
)

// Data
type Shoe struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Style    string  `json:"style"`
	Colour   string  `json:"colour"`
	Material string  `json:"material"`
	Price    float64 `json:"price"`
}

type Inventory struct {
	DB *sql.DB
}

//Interfaces
// GetAll interface
type AllGetter interface {
	GetAll() []Shoe
}

// GetShoe interface
type Getter interface {
	GetShoe(id uint) Shoe
}

// AddShoe interface
type Adder interface {
	AddShoe(shoe Shoe)
}

// UpdateShoe interface
type Updater interface {
	UpdateShoe(shoe Shoe)
}

// DeleteShoe interface
type Deleter interface {
	DeleteShoe(id uint)
}

// Functions
// Creates a new shoe inventory
func New(db *sql.DB) *Inventory {

	db.Exec(`CREATE TABLE IF NOT EXISTS shoe (
				id SERIAL PRIMARY KEY,
				name VARCHAR(255),
		   		style VARCHAR(255),
				colour VARCHAR(255),
				material VARCHAR(255),
				price DOUBLE PRECISION
			);
	`)

	return &Inventory{
		DB: db,
	}
}

// Gets all shoes from the database
func (i *Inventory) GetAll() []Shoe {
	rows, _ := i.DB.Query(`
		SELECT *
		FROM shoe
	`)

	shoes := []Shoe{}
	shoe := Shoe{}

	for rows.Next() {
		rows.Scan(&shoe.ID, &shoe.Name, &shoe.Style, &shoe.Colour, &shoe.Material, &shoe.Price)
		shoes = append(shoes, shoe)
	}

	return shoes
}

// Gets a specific shoe from the database using a ID
func (i *Inventory) GetShoe(id uint) Shoe {
	row, _ := i.DB.Query(`
		SELECT *
		FROM shoe
		WHERE id=$1
	`, id)

	shoe := Shoe{}
	for row.Next() {
		row.Scan(&shoe.ID, &shoe.Name, &shoe.Style, &shoe.Colour, &shoe.Material, &shoe.Price)
	}

	return shoe
}

// Adds a shoe to the database
func (i *Inventory) AddShoe(shoe Shoe) {

	i.DB.Exec(`
		INSERT INTO shoe (name, style, colour, material, price)
		VALUES ($1, $2, $3, $4, $5);	
	`, shoe.Name, shoe.Style, shoe.Colour, shoe.Material, shoe.Price)
}

// Updates a specific shoe from the database using a ID
func (i *Inventory) UpdateShoe(shoe Shoe) {
	i.DB.Exec(`
		UPDATE shoe
		SET name = $1, style = $2, colour = $3, material = $4, price = $5
		WHERE id = $6;
	`, shoe.Name, shoe.Style, shoe.Colour, shoe.Material, shoe.Price, shoe.ID)
}

// Deletes a specific shoe from the database using a ID
func (i *Inventory) DeleteShoe(id uint) {
	i.DB.Query(`
		DELETE FROM shoe
		WHERE id=$1
	`, id)
}
