package usecase

// import (
// 	"CrocsClub/pkg/utils/models"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_AddAddress(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userRepo := mockRepository.NewMockUserRepository(ctrl)

// 	userUseCase := NewUserUseCase(userRepo)

// 	testData := map[string]struct {
// 		input   models.AddAddress
// 		stub    func(*mockRepository.MockUserRepository, models.AddAddress)
// 		wantErr error
// 	}{
// 		"success": {
// 			input: models.AddAddress{
// 				Name:      "ananya",
// 				HouseName: "eduvayil house",
// 				Street:    "peringottukara",
// 				City:      "Thrissur",
// 				State:     "kerala",
// 				Phone:     "1234567890",
// 				Pin:       "680571",
// 			},
// 			stub: func(userRepo *mockRepository.MockUserRepository, data models.AddAddress) {
// 				userRepo.EXPECT().AddAddress(gomock.Any(), data).Return(nil).Times(1)
// 			},
// 			wantErr: nil,
// 		},
// 		"failure": {
// 			input: models.AddAddress{
// 				Name:      "ananya",
// 				HouseName: "eduvayil house",
// 				Street:    "peringottukara",
// 				City:      "Thrissur",
// 				State:     "kerala",
// 				Phone:     "1234567890",
// 				Pin:       "680571",
// 			},
// 			stub: func(userRepo *mockRepository.MockUserRepository, data models.AddAddress) {
// 				userRepo.EXPECT().AddAddress(gomock.Any(), data).Return(errors.New("could not add the address")).Times(1)
// 			},
// 			wantErr: errors.New("could not add the address"),
// 		},
// 	}

// 	for testName, test := range testData {
// 		t.Run(testName, func(t *testing.T) {
// 			test.stub(userRepo, test.input)
// 			err := userUseCase.AddAddress(1, test.input)
// 			assert.Equal(t, test.wantErr, err)
// 		})
// 	}
// }
