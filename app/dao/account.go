package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}

	return entity, nil
}

func (r *account) Create(ctx context.Context, user *object.Account) (error) {
	// it may not be good to rely on DB unique key...(hide logic,)
	_, err := r.db.ExecContext(ctx, "INSERT INTO account (username, password_hash, display_name, avatar, header, note) VALUES (?, ?, ?, ?, ?, ?)",
		user.Username, user.PasswordHash, user.DisplayName, user.Avatar, user.Header, user.Note)

	return err
}
