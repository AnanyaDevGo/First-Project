package usecase

import (
	"CrocsClub/pkg/domain"
	helper_interface "CrocsClub/pkg/helper/interface"
	interfaces "CrocsClub/pkg/repository/interfaces"
	services "CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/models"
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepository interfaces.AdminRepository
	helper          helper_interface.Helper
}

func NewAdminUseCase(repo interfaces.AdminRepository, h helper_interface.Helper) services.AdminUseCase {
	return &adminUseCase{
		adminRepository: repo,
		helper:          h,
	}
}

func (ad *adminUseCase) LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {

	adminCompareDetails, err := ad.adminRepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	var adminDetailsResponse models.AdminDetailsResponse

	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	access, refresh, err := ad.helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin:        adminDetailsResponse,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

}
func (ad *adminUseCase) DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := ad.adminRepository.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := ad.adminRepository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := ad.adminRepository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := ad.adminRepository.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := ad.adminRepository.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	fmt.Println(totalRevenue.MonthRevenue)
	fmt.Println(totalRevenue.YearRevenue)
	fmt.Println(totalRevenue.TodayRevenue)

	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashboardProduct: productDetails,
		DashboardOrder:   orderDetails,
		DashboardRevenue: totalRevenue,
		DashboardAmount:  amountDetails,
	}, nil
}

func (ad *adminUseCase) FilteredSalesReport(timePeriod string) (models.SalesReport, error) {

	startTime, endTime := ad.helper.GetTimeFromPeriod(timePeriod)
	saleReport, err := ad.adminRepository.FilteredSalesReport(startTime, endTime)

	if err != nil {
		return models.SalesReport{}, err
	}
	return saleReport, nil
}

func (ad *adminUseCase) BlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		return errors.New("already blocked")
	} else {
		user.Blocked = true
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

// unblock user
func (ad *adminUseCase) UnBlockUser(id string) error {

	user, err := ad.adminRepository.GetUserByID(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = ad.adminRepository.UpdateBlockUserByID(user)
	if err != nil {
		return err
	}

	return nil

}

func (ad *adminUseCase) GetUsers(page int) ([]models.UserDetailsAtAdmin, error) {

	userDetails, err := ad.adminRepository.GetUsers(page)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (i *adminUseCase) NewPaymentMethod(id string) error {

	exists, err := i.adminRepository.CheckIfPaymentMethodAlreadyExists(id)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("payment method already exists")
	}

	err = i.adminRepository.NewPaymentMethod(id)
	if err != nil {
		return err
	}

	return nil

}

func (a *adminUseCase) ListPaymentMethods() ([]domain.PaymentMethod, error) {

	categories, err := a.adminRepository.ListPaymentMethods()
	if err != nil {
		return []domain.PaymentMethod{}, err
	}
	return categories, nil

}

func (a *adminUseCase) DeletePaymentMethod(id int) error {

	err := a.adminRepository.DeletePaymentMethod(id)
	if err != nil {
		return err
	}
	return nil

}
