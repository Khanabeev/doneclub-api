package category

import (
	"context"
	"doneclub-api/internal/domain/category"
	"doneclub-api/pkg/apperrors"
)

type Service interface {
	CreateNewCategory(ctx context.Context, dto *category.RequestCreateCategoryDTO) (*category.ResponseCategoryDTO, *apperrors.AppError)
	UpdateCategory(ctx context.Context, dto *category.RequestUpdateCategoryDTO, categoryId int) (*category.ResponseCategoryDTO, *apperrors.AppError)
	GetCategory(ctx context.Context, CategoryId int) (*category.ResponseCategoryDTO, *apperrors.AppError)
	GetAllCategories(ctx context.Context) (*category.ResponseAllCategoriesDTO, *apperrors.AppError)
	DeleteCategory(ctx context.Context, categoryId int) (*category.ProfileCategoryDeleted, *apperrors.AppError)
}
