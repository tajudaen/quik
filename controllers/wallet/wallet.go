package wallet

import (
	"net/http"
	"strconv"

	"quik/services/wallet"
	"quik/utils"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	wallet.WalletServiceInterface
}

func (w *WalletController) GetWalletBalance(c *gin.Context) {
	wallet_id, IdErr := getWalletId(c.Param("wallet_id"))
	if IdErr != nil {
		utils.APIError(c, IdErr)
		return
	}

	balance, err := w.WalletServiceInterface.GetWalletBalance(wallet_id)
	if err != nil {
		utils.APIError(c, err)
		return
	}

	utils.APISuccess(c, 200, gin.H{"balance": balance})
}

func getWalletId(wallet_id string) (int64, *utils.RestErr) {
	walletId, err := strconv.ParseInt(wallet_id, 10, 64)
	if err != nil {
		return 0, utils.NewErrorResponse("invalid wallet id", http.StatusBadRequest, "bad_request")
	}
	return walletId, nil
}
