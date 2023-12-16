package interfaces

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/utils/models"
)

type OfferUsecase interface {
	AddProductOffer(ProductOffer models.ProductOfferResp) error
	GetProductOffer() ([]domain.ProductOffer, error)
	ExpireProductOffer(id int) error
	AddCategoryOffer(CategoryOffer models.CategorytOfferResp) error
	GetCategoryOffer() ([]domain.CategoryOffer, error)
	ExpireCategoryOffer(id int) error
}
