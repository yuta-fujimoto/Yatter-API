package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch status which has specified by id
	FindById(ctx context.Context, username object.StatusID) (*object.Status, error)

	// Create New Status
	Create(ctx context.Context, status *object.Status) (error)

	// Delete Status
	Delete(ctx context.Context, id object.StatusID) error
}
