package interfaces

import (
	"CrocsClub/pkg/domain"
)

type CategoryUseCase interface {
	AddCategory(category domain.Category) (domain.Category, error)
	UpdateCategory(current string, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategory() ([]domain.Category, error)
}
