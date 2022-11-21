package orders

import (
	"github.com/Modular-Company/modular-bar/internal/domain/products"
	"time"
)

type Order struct {
	ID            string
	CommandID     string
	CreatedDate   time.Time
	LastUpdated   time.Time
	CreatedBy     string
	OrderProducts []OrderProduct
}

type OrderProduct struct {
	Product  []products.Product
	Quantity int64
}
