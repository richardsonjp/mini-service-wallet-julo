package routes

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/cmd/apiserver/app/store"
)

func V1Route(group *gin.RouterGroup) {
	group.POST("/init", store.CustomerHandler.InitCustomer)

	group.Use(store.MiddlewareAccess.AuthenticateToken())
	group.POST("/wallet", store.WalletHandler.EnableWallet)
	group.PATCH("/wallet", store.WalletHandler.DisableWallet)
	group.GET("/wallet", store.WalletHandler.GetBalance)
	group.POST("/wallet/deposits", store.TransactionHandler.DepositTransaction)
	group.POST("/wallet/withdrawals", store.TransactionHandler.WithdrawalTransaction)
	group.GET("/wallet/transactions", store.TransactionHandler.ListTransaction)
}
