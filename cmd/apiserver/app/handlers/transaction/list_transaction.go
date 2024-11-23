package transaction

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"net/http"
)

func (h *TransactionHandler) ListTransaction(ctx *gin.Context) {
	customerIDctx := ctx.Value("customer_id")
	customerID, ok := customerIDctx.(string)
	if !ok {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	response, err := h.transactionService.ListTransactions(ctx, customerID)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Base{
		Status: "success",
		Data:   response,
	})
}
