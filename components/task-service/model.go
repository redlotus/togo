package main

import (
	"github.com/jackc/pgtype"
)

type (
	// Board
	Board struct {
		ID                int          `db:"id" json:"id"`
		CreatedBy         pgtype.UUID  `db:"created_by" json:"created_by"`
		Title             string       `db:"title" json:"title"`
		Status            int          `db:"status" json:"-"`
		DisplayStatus     string       `json:"status"`
		Visibility        int          `db:"visibility" json:"-"`
		DisplayVisibility string       `json:"visibility"`
		Settings          pgtype.JSONB `db:"settings" json:"settings"`
		Tags              []string     `json:"tags"`
		Base
	}

	Task struct {
		ID                   int         `db:"id" json:"id"`
		CreatedBy            pgtype.UUID `db:"created_by" json:"created_by"`
		Title                string      `db:"title" json:"title"`
		Description          string      `db:"description" json:"description"`
		Status               int         `db:"status" json:"-"`
		DisplayStatus        string      `json:"status"`
		PriorityLevel        int         `db:"priority_level" json:"-"`
		DisplayPriorityLevel string      `json:"priority_level"`
		IsDone               bool        `db:"is_done" json:"is_done"`
		OrderNumber          int         `db:"order_number" json:"order_number"`
		Tags                 []string    `json:"tags"`
		Base
	}

	// Base
	Base struct {
		UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updated_at"`
		CreatedAt pgtype.Timestamp `db:"created_at" json:"created_at"`
	}
)
