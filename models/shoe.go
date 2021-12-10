package shoe

// Data
type Shoe struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Style    string  `json:"style"`
	Color    string  `json:"color"`
	Material string  `json:"material"`
	Price    float32 `json:"price"`
}

type Inventory struct {
	Shoes []Shoe
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
func New() *Inventory {
	return &Inventory{
		Shoes: []Shoe{},
	}
}

// Gets all shoes from the inventory
func (i *Inventory) GetAll() []Shoe {
	return i.Shoes
}

// Gets a specific shoe from the inventory using a ID
func (i *Inventory) GetShoe(id uint) Shoe {
	for _, item := range i.Shoes {
		if id == item.ID {
			return item
		}
	}
	return Shoe{ID: 0}
}

// Adds a shoe to the inventory
func (i *Inventory) AddShoe(shoe Shoe) {
	i.Shoes = append(i.Shoes, shoe)
}

// Updates a specific shoe from the inventory using a ID
func (i *Inventory) UpdateShoe(shoe Shoe) {
	for index, item := range i.Shoes {
		if shoe.ID == item.ID {
			i.Shoes[index] = shoe
			break
		}
	}
}

// Deletes a specific shoe from the inventory using a ID
func (i *Inventory) DeleteShoe(id uint) {
	for index, item := range i.Shoes {
		if id == item.ID {
			i.Shoes = append(i.Shoes[:index], i.Shoes[index+1:]...)
			break
		}
	}
}
