package main

import (

	// "github.com/georgysavva/scany/pgxscan"
	// "github.com/jackc/pgtype"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	GenderCodeMale         = "m"
	GenderCodeFemale       = "f"
	GenderCodeOther        = "o"
	SubtitleFontSizeSmall  = "s"
	SubtitleFontSizeMedium = "m"
	SubtitleFontSizeLarge  = "l"
)

type (
	StoreTodo struct {
		DB *pgxpool.Pool
	}
)

func (s *StoreTodo) AddNewViewer(b Board) (Board, error) {
	return b, nil
}
