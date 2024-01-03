package handler

import (
	"CrocsClub/pkg/domain"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"CrocsClub/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUseCase services.CategoryUseCase
}

func NewCategoryHandler(usecase services.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		CategoryUseCase: usecase,
	}
}

// @Summary Add Category
// @Description Add a new category.
// @Accept json
// @Produce json
// @Tags Admin Category Management
// @security BearerTokenAuth
// @Param category body models.CatRes true "Category details to be added in JSON format"
// @Success 200 {object} response.Response "Successfully added Category"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or Could not add the Category"
// @Router /admin/category [post]
func (Cat *CategoryHandler) AddCategory(c *gin.Context) {

	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	CategoryResponse, err := Cat.CategoryUseCase.AddCategory(category)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Category", CategoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary Update Category
// @Description Update the name of an existing category.
// @Accept json
// @Produce json
// @Tags Admin Category Management
// @security BearerTokenAuth
// @Param body body models.SetNewName true "Category name details in JSON format"
// @Success 200 {object} response.Response "Successfully renamed the category"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or Could not update the Category"
// @Router /admin/category [put]
func (Cat *CategoryHandler) UpdateCategory(c *gin.Context) {

	var p models.SetNewName

	if err := c.BindJSON(&p); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	a, err := Cat.CategoryUseCase.UpdateCategory(p.Current, p.New)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully renamed the category", a, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Delete Category
// @Description Delete an existing category.
// @Accept json
// @Produce json
// @Tags Admin Category Management
// @security BearerTokenAuth
// @Param id query string true "ID of the category to be deleted"
// @Success 200 {object} response.Response "Successfully deleted the Category"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or Could not delete the Category"
// @Router /admin/category [delete]
func (Cat *CategoryHandler) DeleteCategory(c *gin.Context) {

	categoryID := c.Query("id")
	err := Cat.CategoryUseCase.DeleteCategory(categoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the Category", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary Get Categories
// @Description Retrieve all categories.
// @Accept json
// @Produce json
// @Tags Admin Category Management
// @security BearerTokenAuth
// @Success 200 {object} response.Response "Successfully got all categories"
// @Failure 400 {object} response.Response "Fields provided are in the wrong format or Could not retrieve categories"
// @Router /admin/category [get]
func (Cat *CategoryHandler) GetCategory(c *gin.Context) {

	categories, err := Cat.CategoryUseCase.GetCategory()
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all categories", categories, nil)
	c.JSON(http.StatusOK, successRes)

}
