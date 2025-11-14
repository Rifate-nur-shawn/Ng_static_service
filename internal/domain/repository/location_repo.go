package repository

import (
	"context"
	model "matrimonial-service/internal/domain/model"
)

// Location repository interface
type LocationRepository interface {
	GetDivisions(ctx context.Context) ([]*model.Division, error)
	GetDistrictsByDivision(ctx context.Context, divisionID int64) ([]*model.District, error)
	GetUpazilasByDistrict(ctx context.Context, districtID int64) ([]*model.Upazila, error)
}