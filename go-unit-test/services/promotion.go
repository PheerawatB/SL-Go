package services

import (
	"gotest/repositories"
)

type PromotionService interface {
	CalculateDiscount(purchaseAmount int) (int, error)
}

type promotionService struct {
	promoRepo repositories.PromotionRepository
}

func NewPromotionService(promoRepo repositories.PromotionRepository) PromotionService {
	return promotionService{promoRepo}
}

func (base promotionService) CalculateDiscount(purchaseAmount int) (int, error) {
	if purchaseAmount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := base.promoRepo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if purchaseAmount >= promotion.PurchaseMin {
		return purchaseAmount - (promotion.DicountPercent * purchaseAmount / 100), nil
	}
	return purchaseAmount, nil
}
