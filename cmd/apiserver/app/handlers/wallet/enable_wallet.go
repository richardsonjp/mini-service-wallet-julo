package wallet

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/services/wallet"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"net/http"
)

func (h *WalletHandler) EnableWallet(ctx *gin.Context) {
	customerIDctx := ctx.Value("customer_id")
	customerID, ok := customerIDctx.(string)
	if !ok {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	response, err := h.walletService.ToggleWallet(ctx, wallet.ToggleWallet{
		CustomerID: customerID,
		Enabled:    true,
	})
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Base{
		Status: "success",
		Data:   response,
	})
}
