package repository

import (
	"context"
	model "matrimonial-service/internal/domain/model"
)

type StaticDataRepository interface {
	GetReligions(ctx context.Context) ([]*model.StaticOption, error)
	GetProfessions(ctx context.Context) ([]*model.StaticOption, error)
	GetMaritalStatuses(ctx context.Context) ([]*model.StaticOption, error)
}
