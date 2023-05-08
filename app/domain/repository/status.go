package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	// Fetch status which has specified by id\
	FindById(ctx context.Context, id object.StatusID) (*object.Status, error)

	// Fetch statuses which meets the conditions
	FindMany(ctx context.Context, sinceId object.StatusID, maxId object.StatusID, limit int64, onlyMedia bool) (*[]object.Status, error)

	// Create New Status
	Create(ctx context.Context, status *object.Status) (error)

	// Delete Status
	Delete(ctx context.Context, id object.StatusID) error
}
