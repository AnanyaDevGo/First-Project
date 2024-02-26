package usecase

import (
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/domain"
	mockhelper "CrocsClub/pkg/helper/mock"
	mockRepository "CrocsClub/pkg/repository/mock"

	"CrocsClub/pkg/utils/models"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockRepository.NewMockUserRepository(ctrl)
	cfg := config.Config{}
	wallet := mockRepository.NewMockWalletRepository(ctrl)
	otp := mockRepository.NewMockOtpRepository(ctrl)
	inv := mockRepository.NewMockInventoryRepository(ctrl)
	order := mockRepository.NewMockOrderRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)

	userUseCase := NewUserUseCase(userRepo, cfg, wallet, otp, inv, order, helper)
	testingData := map[string]struct {
		input   int
		stub    func(*mockRepository.MockUserRepository, *mockhelper.MockHelper, int)
		want    []domain.Address
		wantErr error
	}{
		"success": {
			input: 1,
			stub: func(ur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, i int) {
				ur.EXPECT().GetAddress(i).Times(1).Return([]domain.Address{}, nil)
			},
			want:    []domain.Address{},
			wantErr: nil,
		},
		"failed": {
			input: 1,
			stub: func(ur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, i int) {
				ur.EXPECT().GetAddress(i).Times(1).Return([]domain.Address{}, errors.New("error"))
			},
			want:    []domain.Address{},
			wantErr: errors.New("error in getting addresses"),
		},
	}
	for _, test := range testingData {
		test.stub(userRepo, helper, test.input)
		result, err := userUseCase.GetAddress(test.input)
		assert.Equal(t, test.want, result)
		assert.Equal(t, test.wantErr, err)
	}

}
func Test_GetUserDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mockRepository.NewMockUserRepository(ctrl)
	cfg := config.Config{}
	wallet := mockRepository.NewMockWalletRepository(ctrl)
	otp := mockRepository.NewMockOtpRepository(ctrl)
	inv := mockRepository.NewMockInventoryRepository(ctrl)
	order := mockRepository.NewMockOrderRepository(ctrl)
	helper := mockhelper.NewMockHelper(ctrl)

	userUseCase := NewUserUseCase(userRepo, cfg, wallet, otp, inv, order, helper)
	testingData := map[string]struct {
		input   int
		stub    func(*mockRepository.MockUserRepository, *mockhelper.MockHelper, int)
		want    models.UserDetailsResponse
		wantERr error
	}{
		"success": {
			input: 1,
			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, i int) {
				mur.EXPECT().GetUserDetails(i).Times(1).Return(models.UserDetailsResponse{}, nil)
			},
			want:    models.UserDetailsResponse{},
			wantERr: nil,
		},
		"failed": {
			input: 1,
			stub: func(mur *mockRepository.MockUserRepository, mh *mockhelper.MockHelper, i int) {
				mur.EXPECT().GetUserDetails(i).Times(1).Return(models.UserDetailsResponse{}, errors.New("error"))
			},
			want:    models.UserDetailsResponse{},
			wantERr: errors.New("error in getting details"),
		},
	}
	for _, test := range testingData {
		test.stub(userRepo, helper, test.input)
		result, err := userUseCase.GetUserDetails(test.input)
		assert.Equal(t, test.want, result)
		assert.Equal(t, test.wantERr, err)
	}

}
