package interfaces

import "CrocsClub/pkg/domain"

type CategoryRepository interface {
	AddCategory(category domain.Category) (domain.Category, error)
	CheckCategory(current string) (bool, error)
	UpdateCategory(current, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategory() ([]domain.Category, error)
	CategoryExists(categoryID int) (bool, error)
}
