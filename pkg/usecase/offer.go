package usecase

import (
	"CrocsClub/pkg/domain"
	"CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
)

type offerUsecase struct {
	repo interfaces.OfferRepository
}

func NewOferUsecase(repo interfaces.OfferRepository) services.OfferUsecase {
	return &offerUsecase{
		repo: repo,
	}
}

func (ou *offerUsecase) AddProductOffer(ProductOffer models.ProductOfferResp) error {
	if err := ou.repo.AddProductOffer(ProductOffer); err != nil {
		return errors.New("error in adding offer")
	}
	return nil
}
func (ou *offerUsecase) AddCategoryOffer(CategoryOffer models.CategorytOfferResp) error {
	if err := ou.repo.AddCategoryOffer(CategoryOffer); err != nil {
		return errors.New("error in adding offer")
	}
	return nil
}
func (ou *offerUsecase) GetProductOffer() ([]domain.ProductOffer, error) {
	offers, err := ou.repo.GetProductOffer()
	if err != nil {
		return []domain.ProductOffer{}, err
	}
	return offers, nil
}

func (ou *offerUsecase) GetCategoryOffer() ([]domain.CategoryOffer, error) {
	offers, err := ou.repo.GetCategoryOffer()
	if err != nil {
		return []domain.CategoryOffer{}, err
	}
	return offers, nil
}

func (ou *offerUsecase) ExpireProductOffer(id int) error {
	if err := ou.repo.ExpireProductOffer(id); err != nil {
		return err
	}

	return nil
}

func (ou *offerUsecase) ExpireCategoryOffer(id int) error {
	if err := ou.repo.ExpireCategoryOffer(id); err != nil {
		return err
	}

	return nil
}
