package app

import (
	walletController "quik/controllers/wallet"
	"quik/database/mysql"
	walletDomain "quik/domain/wallet"
	"quik/middlewares"
	"quik/providers/redis"
	walletService "quik/services/wallet"
)

var (
	WalletDAO walletDomain.WalletDAOInterface = &walletDomain.Wallet_DAO{mysql.MySQL}

	WalletServices walletService.WalletServiceInterface = &walletService.WalletService{WalletDAO}
	RedisCache     redis.RedisInterface                 = &redis.Redis{}
)

func registerRoutes() {

	wc := &walletController.WalletController{WalletServices, RedisCache}
	v1 := router.Group("/api/v1/")
	{
		wallets := v1.Group("/wallets/")
		wallets.GET("/:wallet_id/balance", middlewares.CacheBalance(), wc.GetWalletBalance)
		wallets.POST("/:wallet_id/credit", wc.CreditWallet)
		wallets.POST("/:wallet_id/debit", wc.DebitWallet)
	}
}
