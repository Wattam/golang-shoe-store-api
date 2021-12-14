package shoe

import (
	"fmt"

	"gorm.io/gorm"
)

// Data
type Shoe struct {
	gorm.Model
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Style    string  `json:"style"`
	Colour   string  `json:"colour"`
	Material string  `json:"material"`
	Price    float64 `json:"price"`
}

type Inventory struct {
	GormDB *gorm.DB
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
func New(gormDB *gorm.DB) *Inventory {

	gormDB.AutoMigrate(&Shoe{})

	return &Inventory{
		GormDB: gormDB,
	}
}

// Gets all shoes from the database
func (i *Inventory) GetAll() []Shoe {

	shoes := []Shoe{}
	i.GormDB.Find(&shoes)

	return shoes
}

// Gets a specific shoe from the database using a ID
func (i *Inventory) GetShoe(id uint) Shoe {
	var shoe Shoe
	i.GormDB.First(&shoe, id)

	fmt.Println(shoe)
	return shoe
}

// Adds a shoe to the database
func (i *Inventory) AddShoe(shoe Shoe) {

	i.GormDB.Create(&shoe)
}

// Updates a specific shoe from the database using a ID
func (i *Inventory) UpdateShoe(shoe Shoe) {
	i.GormDB.Save(&shoe)
}

// Deletes a specific shoe from the database using a ID
func (i *Inventory) DeleteShoe(id uint) {
	i.GormDB.Delete(&Shoe{}, id)
}
