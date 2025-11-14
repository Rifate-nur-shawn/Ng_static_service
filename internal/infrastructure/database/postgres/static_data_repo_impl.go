package postgres

import (
	"context"
	"database/sql"
	"matrimonial-service/internal/domain/model"
	"matrimonial-service/internal/domain/repository"
)

type pgStaticDataRepository struct {
	db *sql.DB
}

// NewPgStaticDataRepository creates a new static data repository
func NewPgStaticDataRepository(db *sql.DB) repository.StaticDataRepository {
	return &pgStaticDataRepository{db: db}
}

// genericGetOptions is a helper to fetch simple static data
func (r *pgStaticDataRepository) genericGetOptions(ctx context.Context, tableName string) ([]*model.StaticOption, error) {
	// Note: Be careful with string formatting in SQL. 
	// This is safe ONLY because tableName is controlled internally by us, not by user input.
	query := `SELECT id, name FROM "` + tableName + `" ORDER BY name`
	if tableName == "marital_statuses" {
		query = `SELECT id, status as name FROM "marital_statuses" ORDER BY id`
	}
	if tableName == "blood_groups" {
		query = `SELECT id, group_name as name FROM "blood_groups" ORDER BY id`
	}
	if tableName == "skin_tones" {
		query = `SELECT id, tone as name FROM "skin_tones" ORDER BY id`
	}

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var options []*model.StaticOption
	for rows.Next() {
		var opt model.StaticOption
		if err := rows.Scan(&opt.ID, &opt.Name); err != nil {
			return nil, err
		}
		options = append(options, &opt)
	}
	return options, nil
}

func (r *pgStaticDataRepository) GetReligions(ctx context.Context) ([]*model.StaticOption, error) {
	return r.genericGetOptions(ctx, "religions")
}

func (r *pgStaticDataRepository) GetProfessions(ctx context.Context) ([]*model.StaticOption, error) {
	return r.genericGetOptions(ctx, "professions")
}

func (r *pgStaticDataRepository) GetMaritalStatuses(ctx context.Context) ([]*model.StaticOption, error) {
	return r.genericGetOptions(ctx, "marital_statuses")
}