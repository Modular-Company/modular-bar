package users

type User struct {
	ID       string
	Username string
	Name     string
	Password string
	Role     []Role
}

type Role string

const (
	Admin     Role = "admin"
	Bartender Role = "bartender"
	Cashier   Role = "cashier"
	Attendant Role = "attendant"
)
