package customer

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/services/customer"
	"go-skeleton/pkg/utils/api"
	"go-skeleton/pkg/utils/errors"
	"go-skeleton/pkg/utils/validator"
)

func (h *CustomerHandler) InitCustomer(ctx *gin.Context) {
	param := customer.InitCustomerPayload{}
	if err := ctx.ShouldBind(&param); err != nil {
		errors.ErrorString(ctx, validator.GetValidatorMessage(err))
		return
	}

	response, err := h.customerService.InitCustomer(ctx, param)
	if err != nil {
		errors.E(ctx, err)
		return
	}

	ctx.JSON(201, api.Base{
		Status: "success",
		Data:   response,
	})
}
