package handler

import (
	"CrocsClub/pkg/usecase/interfaces"
	"CrocsClub/pkg/utils/response"
	"fmt"
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
	userIDRaw, exists := c.Get("id")
	fmt.Println("userid", userIDRaw)
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "user_id not found", nil, "user_id is required")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID, ok := userIDRaw.(int)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "invalid user_id type", nil, "user_id must be an integer")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	WalletDetails, err := w.WalletUseCase.GetWallet(userID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Wallet Details", WalletDetails, nil)
	c.JSON(http.StatusOK, success)
}

func (w *WalletHandler) WalletHistory(c *gin.Context) {
	userIDRaw, exists := c.Get("id")
	if !exists {
		errs := response.ClientResponse(http.StatusBadRequest, "user_id not found", nil, "user_id is required")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	userID, ok := userIDRaw.(int)
	if !ok {
		errs := response.ClientResponse(http.StatusBadRequest, "invalid user_id type", nil, "user_id must be an integer")
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	WalletDetails, err := w.WalletUseCase.WalletHistory(userID)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}

	success := response.ClientResponse(http.StatusOK, "Wallet Details", WalletDetails, nil)
	c.JSON(http.StatusOK, success)
}
