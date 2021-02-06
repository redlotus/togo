package main

import (

	// "github.com/georgysavva/scany/pgxscan"
	// "github.com/jackc/pgtype"

	"github.com/jackc/pgx/v4/pgxpool"
)

const ()

type (
	StoreTodo struct {
		DB *pgxpool.Pool
	}
)

func (s *StoreTodo) AddNewBoard(b Board) (Board, error) {
	rb := Board{}
	tx, err := s.DB.Begin(dbContext)
	if err != nil {
		return rb, err
	}

	err = tx.Commit(dbContext)
	if err != nil {
		return rb, err
	}
	return rb, nil
}
