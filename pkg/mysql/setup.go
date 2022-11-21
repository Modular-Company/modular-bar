package mysql

import (
	"context"
	"fmt"

	"cmd/app/main.go/internal/infra/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Load(ctx context.Context, cfg config.Database) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4_unicode_ci&parseTime=true&multiStatements=true&loc=America/Sao_Paulo",
		cfg.Username,
		cfg.Password,
		cfg.Server,
		cfg.Schema)

	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIddleConns)

	return db, nil
}
