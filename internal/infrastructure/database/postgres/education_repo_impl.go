package postgres

import (
	"context"
	"database/sql"
	"matrimonial-service/internal/domain/model"
	"matrimonial-service/internal/domain/repository"
)

type pgEducationRepository struct {
	db *sql.DB
}

// NewPgEducationRepository creates a new education repository
func NewPgEducationRepository(db *sql.DB) repository.EducationRepository {
	return &pgEducationRepository{db: db}
}

// GetQualifications fetches all qualifications, joining with levels and streams
func (r *pgEducationRepository) GetQualifications(ctx context.Context) ([]*model.Qualification, error) {
	query := `
        SELECT 
            q.id, 
            q.level_id, 
            l.stream_id,
            q.qualification_name_en, 
            COALESCE(q.common_name, ''), 
            q.equivalent_to
        FROM qualifications q
        JOIN education_levels l ON q.level_id = l.id
        ORDER BY l.stream_id, q.level_id, q.qualification_name_en
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var qualifications []*model.Qualification
	for rows.Next() {
		var q model.Qualification
		var equivalentTo sql.NullInt64 // Handle NULLable foreign key
		
		if err := rows.Scan(
			&q.ID,
			&q.LevelID,
			&q.StreamID,
			&q.Name,
			&q.CommonName,
			&equivalentTo,
		); err != nil {
			return nil, err
		}
		
		if equivalentTo.Valid {
			q.EquivalentTo = &equivalentTo.Int64
		}
		
		qualifications = append(qualifications, &q)
	}
	
	return qualifications, nil
}