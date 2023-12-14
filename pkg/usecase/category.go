package usecase

import (
	"CrocsClub/pkg/domain"
	helper_interface "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"errors"
)

type categoryUseCase struct {
	repository interfaces.CategoryRepository
	helper     helper_interface.Helper
}

func NewCategoryUseCase(repo interfaces.CategoryRepository, h helper_interface.Helper) services.CategoryUseCase {
	return &categoryUseCase{
		repository: repo,
		helper:     h,
	}
}

func (Cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {
	if category.Category == ""{
		return domain.Category{}, errors.New("category should not be empty")
	}

	productResponse, err := Cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil

}

func (Cat *categoryUseCase) UpdateCategory(current string, new string) (domain.Category, error) {
	if current == "" || new == "" {
		return domain.Category{},errors.New("values should not be empty")
	}

	result, err := Cat.repository.CheckCategory(current)
	if err != nil {
		return domain.Category{}, err
	}

	if !result {
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}

	newcat, err := Cat.repository.UpdateCategory(current, new)
	if err != nil {
		return domain.Category{}, err
	}

	return newcat, err
}

func (Cat *categoryUseCase) DeleteCategory(categoryID string) error {

	err := Cat.repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil

}

func (Cat *categoryUseCase) GetCategory() ([]domain.Category, error) {

	categories, err := Cat.repository.GetCategory()
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}
func (c *categoryUseCase) CategoryExists(categoryID int) bool {
	exists, err := c.repository.CategoryExists(categoryID)
	if err != nil {
		return false
	}
	return exists
}
