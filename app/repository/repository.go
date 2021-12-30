package repository

import (
	"context"
	"pg_slowest/models"
)

//DBStats ...
type DBStats interface {
	Stats(ctx context.Context, filter string, pg *models.PaginationRequest) (*models.DBResponse, error)
}
