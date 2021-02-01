package main

import (
	"github.com/jackc/pgtype"
)

type (
	// Board
	Board struct {
		ID         int          `db:"id" json:"id"`
		AccountID  pgtype.UUID  `db:"account_id" json:"account_id"`
		Settings   pgtype.JSONB `db:"settings" json:"settings"`
		Status     int          `db:"status" json:"status"`
		Visibility int          `db:"visibility" json:"visibility"`
		Base
	}

	// Base
	Base struct {
		UpdatedAt pgtype.Timestamp `db:"updated_at"`
		CreatedAt pgtype.Timestamp `db:"created_at"`
	}
)
