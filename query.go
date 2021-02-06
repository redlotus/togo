package main

const (
	QueryAddNewBoard = `INSERT INTO board(created_by, title, visibility, settings) VALUES ($1, $2, $3, $4)`
	QueryAddNewTask  = ``
	QueryAddBoardTag = `INSERT INTO board_tag(board_id, tag_id) VALUES ($1, $2)`
	QueryAddTagBoard = `INSERT INTO tag(name, type) VALUES($1, 0)`
	QueryAddTagTask  = `INSERT INTO tag(name, type) VALUES($1, 1)`
)
