package implementation

import (
	"context"
	"pg_slowest/app/domain"
	"pg_slowest/app/repository"
	"pg_slowest/models"
)

type impl struct {
	repo repository.DBStats
}

//NewDomain initiates the domain interface implementor.
func NewDomain(repo repository.DBStats) domain.Stater {
	return &impl{repo: repo}
}

//Stat collects data from database and returns.
func (i *impl) Stat(ctx context.Context, request *models.APIRequest) (*models.DBResponse, error) {
	return i.repo.Stats(ctx, request.Filter, request.Pagination)
}
