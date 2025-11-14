package postgres

import (
	"context"
	"database/sql"
	"matrimonial-service/internal/domain/model"
	"matrimonial-service/internal/domain/repository"
)

type pgLocationRepository struct {
	db *sql.DB
}

// NewPgLocationRepository creates a new location repository
func NewPgLocationRepository(db *sql.DB) repository.LocationRepository {
	return &pgLocationRepository{db: db}
}

// GetDivisions fetches all divisions
func (r *pgLocationRepository) GetDivisions(ctx context.Context) ([]*model.Division, error) {
	query := `SELECT id, name_en, COALESCE(name_bn, '') FROM divisions ORDER BY name_en`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var divisions []*model.Division
	for rows.Next() {
		var d model.Division
		if err := rows.Scan(&d.ID, &d.NameEn, &d.NameBn); err != nil {
			return nil, err
		}
		divisions = append(divisions, &d)
	}

	return divisions, nil
}

// GetDistrictsByDivision fetches all districts for a given division
func (r *pgLocationRepository) GetDistrictsByDivision(ctx context.Context, divisionID int64) ([]*model.District, error) {
	query := `
		SELECT id, division_id, name_en, COALESCE(name_bn, '')
		FROM districts
		WHERE division_id = $1
		ORDER BY name_en
	`

	rows, err := r.db.QueryContext(ctx, query, divisionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var districts []*model.District
	for rows.Next() {
		var d model.District
		if err := rows.Scan(&d.ID, &d.DivisionID, &d.NameEn, &d.NameBn); err != nil {
			return nil, err
		}
		districts = append(districts, &d)
	}

	return districts, nil
}

// GetUpazilasByDistrict fetches all upazilas for a given district
func (r *pgLocationRepository) GetUpazilasByDistrict(ctx context.Context, districtID int64) ([]*model.Upazila, error) {
	query := `
		SELECT id, district_id, name_en, COALESCE(name_bn, '')
		FROM upazilas
		WHERE district_id = $1
		ORDER BY name_en
	`

	rows, err := r.db.QueryContext(ctx, query, districtID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var upazilas []*model.Upazila
	for rows.Next() {
		var u model.Upazila
		if err := rows.Scan(&u.ID, &u.DistrictID, &u.NameEn, &u.NameBn); err != nil {
			return nil, err
		}
		upazilas = append(upazilas, &u)
	}

	return upazilas, nil
}

// GetUpazilaById fetches a single upazila by ID
func (r *pgLocationRepository) GetUpazilaById(ctx context.Context, upazilaID int64) (*model.Upazila, error) {
	query := `
		SELECT id, district_id, name_en, COALESCE(name_bn, '')
		FROM upazilas
		WHERE id = $1
	`

	var u model.Upazila
	err := r.db.QueryRowContext(ctx, query, upazilaID).Scan(&u.ID, &u.DistrictID, &u.NameEn, &u.NameBn)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// SearchUpazilasByDistrict searches upazilas by name in a district
func (r *pgLocationRepository) SearchUpazilasByDistrict(ctx context.Context, districtID int64, search string) ([]*model.Upazila, error) {
	query := `
		SELECT id, district_id, name_en, COALESCE(name_bn, '')
		FROM upazilas
		WHERE district_id = $1 AND LOWER(name_en) LIKE LOWER($2)
		ORDER BY name_en
	`

	rows, err := r.db.QueryContext(ctx, query, districtID, "%"+search+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var upazilas []*model.Upazila
	for rows.Next() {
		var u model.Upazila
		if err := rows.Scan(&u.ID, &u.DistrictID, &u.NameEn, &u.NameBn); err != nil {
			return nil, err
		}
		upazilas = append(upazilas, &u)
	}

	return upazilas, nil
}
