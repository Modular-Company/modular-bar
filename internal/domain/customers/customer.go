package customers

type Customer struct {
	ID          string
	Name        string
	Surname     string
	Document    string
	PhoneNumber string
	Commands    []commands.Command
}
