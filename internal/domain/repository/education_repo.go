// Education repository interfacepackage repository
package repository

import (
	"context"
	model "matrimonial-service/internal/domain/model"
)


type EducationRepository interface {
	GetQualifications(ctx context.Context) ([]*model.Qualification, error)
}