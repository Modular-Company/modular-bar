package commands

import (
	"github.com/Modular-Company/modular-bar/internal/domain/orders"
	"time"
)

type Command struct {
	ID          string
	CustomerID  string
	Orders      []orders.Order
	CreatedDate time.Time
	LastUpdated time.Time
	ClosedDate  *time.Time
	CreatedBy   string
	ClosedBy    *string
	IsOpen      bool
}
