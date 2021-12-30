package domain

import (
	"context"
	"pg_slowest/models"
)

//Stater ...
type Stater interface {
	Stat(ctx context.Context,request *models.APIRequest) (*models.DBResponse, error)
}
