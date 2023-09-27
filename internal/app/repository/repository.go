package repository

import "database/sql"

type Repository struct {
	Conn *sql.DB
	CFT  *sql.DB
}

func New(conn *sql.DB) *Repository {
	return &Repository{Conn: conn}
}
