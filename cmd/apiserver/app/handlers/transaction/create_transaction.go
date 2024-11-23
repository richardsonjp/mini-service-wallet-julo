package transaction

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/services/transaction"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
	"net/http"
)

func (h *TransactionHandler) DepositTransaction(ctx *gin.Context) {
	customerIDctx := ctx.Value("customer_id")
	customerID, ok := customerIDctx.(string)
	if !ok {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	param := transaction.CreateTransactionPayload{}
	if err := ctx.ShouldBind(&param); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}

	param.TransactionType = "deposit"
	param.CustomerID = customerID
	response, err := h.transactionService.CreateTransaction(ctx, param)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Base{
		Status: "success",
		Data:   response,
	})
}

func (h *TransactionHandler) WithdrawalTransaction(ctx *gin.Context) {
	customerIDctx := ctx.Value("customer_id")
	customerID, ok := customerIDctx.(string)
	if !ok {
		errors.ErrorCode(ctx, http.StatusUnauthorized)
		return
	}

	param := transaction.CreateTransactionPayload{}
	if err := ctx.ShouldBind(&param); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}

	param.TransactionType = "withdrawal"
	param.CustomerID = customerID
	response, err := h.transactionService.CreateTransaction(ctx, param)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Base{
		Status: "success",
		Data:   response,
	})
}
