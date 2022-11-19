package products

type Product struct {
	ID       string
	Category Category
	Name     string
	Value    float64
	Stock    int64
}

type Category string

const (
	Drink   Category = "drink"
	Food    Category = "food"
	Service Category = "service"
)
