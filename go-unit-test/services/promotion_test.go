package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	type testCase struct {
		name            string
		PurchaseMin     int
		purchaseAmount  int
		discountPercent int
		expected        int
	}
	cases := []testCase{
		{"PurchaseMin is 100, purchaseAmount is 100, discountPercent is 20", 100, 100, 20, 80},
		{"PurchaseMin is 100, purchaseAmount is 100, discountPercent is 10", 100, 100, 10, 90},
		{"PurchaseMin is 100, purchaseAmount is 100, discountPercent is 5", 100, 100, 5, 95},
		{"PurchaseMin is 100, purchaseAmount is 100, discountPercent is 0", 100, 100, 0, 100},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			repo := repositories.NewPromotionRepositoryMock()
			repo.On("GetPromotion").Return(repositories.Promotion{
				ID:             tc.PurchaseMin,
				DicountPercent: tc.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(repo)
			discount, _ := promoService.CalculateDiscount(tc.purchaseAmount)
			assert.Equal(t, tc.expected, discount)
		})
	}
	t.Run("zero purchase amount", func(t *testing.T) {
		//Arrang
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			ID:             1,
			PurchaseMin:    100,
			DicountPercent: 20,
		}, nil)

		promoService := services.NewPromotionService(promoRepo)
		_, err := promoService.CalculateDiscount(0)
		//Assert
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		promoRepo.AssertNotCalled(t, "GetPromotion")
	})

	t.Run("repository error", func(t *testing.T) {
		//Arrang
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New("error"))

		promoService := services.NewPromotionService(promoRepo)

		//Act
		_, err := promoService.CalculateDiscount(100)

		//Assert
		assert.ErrorIs(t, err, services.ErrRepository)
	})

}
