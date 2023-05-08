package dao

import (
	"context"
	"database/sql"
	"errors"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{ db: db }
}

func (r *status) Create(ctx context.Context, status *object.Status) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO status (account_id, content) VALUES (?, ?)",
		status.AccountID, status.Content)

	return err
}

func (r *status) FindById(ctx context.Context, id object.StatusID) (*object.Status, error) {
	entity := new(object.Status)

	err := r.db.QueryRowxContext(ctx, "SELECT * FROM status WHERE id=?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		
		return nil, err
	}
	return entity, nil
} 

func (r *status) FindMany(ctx context.Context, sinceId object.StatusID, maxId object.StatusID, limit int64, onlyMedia bool) (*[]object.Status, error) {
	entity := new([]object.Status)
	rows, err := r.db.QueryContext(ctx, "SELECT id, account_id, content, create_at FROM status WHERE id > ? AND id < ? LIMIT ?", sinceId, maxId, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var status object.Status
		err := rows.Scan(&status.ID, &status.AccountID, &status.Content, &status.CreateAt)
		if err != nil {
			return nil, err
		}

		*entity = append(*entity, status)
	}

	return entity, nil
}


func (r *status) Delete(ctx context.Context, id object.StatusID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM status WHERE id=?", id)
	return err
}