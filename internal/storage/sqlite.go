package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqlite(storagePath string) (*sql.DB, error) {
	const op = "storage.sqlite.New"

	dataSource := fmt.Sprintf("file:%s?_foreign_keys=on", storagePath)
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
