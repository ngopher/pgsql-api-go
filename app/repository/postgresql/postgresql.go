package postgresql

import (
	"context"
	"gorm.io/gorm"
	"pg_slowest/app/repository"
	"pg_slowest/models"
)

type pgsql struct {
	db *gorm.DB
}

//NewPGSQL returns an instance of DBStats.
func NewPGSQL(db *gorm.DB) repository.DBStats {
	return &pgsql{db: db}
}

func (p *pgsql) Stats(ctx context.Context, filter string, pg *models.PaginationRequest) (*models.DBResponse, error) {
	q := p.db.Raw("SELECT * FROM public.pg_stat_statements")
	// Apply Filter if there is any.
	if filter != "" {
		q.Where("query = (?)", filter)
	}

	// Apply Pagination and Sort
	q.Limit(pg.Limit).
		Offset(pg.Offset).
		Order("total_time")

	var result models.DBResponse
	err := q.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}
