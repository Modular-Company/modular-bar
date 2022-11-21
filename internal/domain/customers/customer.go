package customers

import "github.com/Modular-Company/modular-bar/internal/domain/commands"

type Customer struct {
	ID          string
	Name        string
	Surname     string
	Document    string
	PhoneNumber string
	Commands    []commands.Command
}
