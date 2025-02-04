package storage

import "github.com/jackc/pgx/v5/pgxpool"

type PostStorage struct {
	pg *pgxpool.Pool
}

func RegisterNewPostStorage(pg *pgxpool.Pool) *PostStorage {
	return &PostStorage{
		pg: pg,
	}
}
