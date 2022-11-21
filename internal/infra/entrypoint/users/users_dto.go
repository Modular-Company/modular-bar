package users

import "github.com/Modular-Company/modular-bar/internal/domain/users"

type CreateUserDTO struct {
	Username string
	Password string
	Name     string
	Roles    []users.Role
}

func (d *CreateUserDTO) ToDomain() *users.User {
	return &users.User{
		Username: d.Username,
		Password: d.Password,
		Name:     d.Name,
		Role:     d.Roles,
	}
}
