package infra

import (
	"context"
	"net/http"

	"github.com/Modular-Company/modular-bar/internal/infra/config"
	"github.com/Modular-Company/modular-bar/internal/infra/entrypoint"
	"github.com/Modular-Company/modular-bar/internal/infra/entrypoint/users"
	"github.com/Modular-Company/modular-bar/pkg/mysql"
)

func Run() error {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := mysql.Load(ctx, cfg.Database)
	if err != nil {
		return err
	}
	defer db.Close()

	usersHandler := users.NewUsersHandler()

	router, err := entrypoint.Load(usersHandler)
	if err != nil {
		return err
	}

	return http.ListenAndServe(cfg.Port, router)
}
