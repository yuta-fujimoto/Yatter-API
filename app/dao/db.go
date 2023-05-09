package dao

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

// Interface of configureation
type DBConfig interface {
	FormatDSN() string
}

// Prepare sqlx.DB
func initDb(config DBConfig) (*sqlx.DB, error) {
	driverName := "mysql"
	db, err := sqlx.Open(driverName, config.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open failed: %w", err)
	}

	return db, nil
}

func initDbMock() (*sqlx.DB,sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	dbx := sqlx.NewDb(db, "in memory")
	return dbx, mock, nil
}
