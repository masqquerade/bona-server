package store

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Store struct {
	Config *Config
	Db     *pgx.Conn
}

func NewStore(config *Config) (*Store, error) {
	conn, err := pgx.Connect(context.Background(), config.CreateURL())
	if err != nil {
		return nil, err
	}

	return &Store{
		Config: config,
		Db:     conn,
	}, nil
}
