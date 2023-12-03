package handler

import (
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	WalletUseCase interfaces.WalletUseCase
}

func NewWalletHandler(usecase interfaces.WalletUseCase) *WalletHandler {
	return &WalletHandler{
		WalletUseCase: usecase,
	}
}
func (w *WalletHandler) GetWallet(c *gin.Context) {
	userID, _ := c.Get("user_id")
	WalletDetails, err := w.WalletUseCase.GetWallet(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Wallet Details", WalletDetails, nil)
	c.JSON(http.StatusOK, success)
}

func (w *WalletHandler) WalletHistory(c *gin.Context) {
	userID, _ := c.Get("user_id")
	WalletDetails, err := w.WalletUseCase.WalletHistory(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Wallet Details", WalletDetails, nil)
	c.JSON(http.StatusOK, success)
}
