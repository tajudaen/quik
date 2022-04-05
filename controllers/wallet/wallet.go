package wallet

import (
	"encoding/json"
	"net/http"
	"strconv"

	"quik/logger"
	"quik/services/wallet"
	"quik/types"
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

func (w *WalletController) CreditWallet(c *gin.Context) {
	wallet_id, IdErr := getWalletId(c.Param("wallet_id"))
	if IdErr != nil {
		utils.APIError(c, IdErr)
		return
	}

	var data types.CreditRequest
	input := json.NewDecoder(c.Request.Body)

	if err := input.Decode(&data); err != nil {
		logger.Error(
			"error while unmarshaling credit data",
			err, map[string]interface{}{
				"method": c.Request.Method,
				"url":    c.Request.URL.String(),
				"data":   c.Request.Body,
			})

		utils.InvalidRequestBodyAPIError(c)
		return
	}

	if ok, errValidate := utils.StructValidateHelper(data); ok {
		utils.RequestValidationAPIError(c, errValidate[0])
		return
	}

	err := w.WalletServiceInterface.CreditWallet(wallet_id, data.Amount)
	if err != nil {
		utils.APIError(c, err)
		return
	}

	utils.APISuccess(c, 200, gin.H{"message": "success"})
}

func (w *WalletController) DebitWallet(c *gin.Context) {
	wallet_id, IdErr := getWalletId(c.Param("wallet_id"))
	if IdErr != nil {
		utils.APIError(c, IdErr)
		return
	}

	var data types.CreditRequest
	input := json.NewDecoder(c.Request.Body)

	if err := input.Decode(&data); err != nil {
		logger.Error(
			"error while unmarshaling debit data",
			err, map[string]interface{}{
				"method": c.Request.Method,
				"url":    c.Request.URL.String(),
				"data":   c.Request.Body,
			})

		utils.InvalidRequestBodyAPIError(c)
		return
	}

	if ok, errValidate := utils.StructValidateHelper(data); ok {
		utils.RequestValidationAPIError(c, errValidate[0])
		return
	}

	err := w.WalletServiceInterface.DebitWallet(wallet_id, data.Amount)
	if err != nil {
		utils.APIError(c, err)
		return
	}

	utils.APISuccess(c, 200, gin.H{"message": "success"})
}

func getWalletId(wallet_id string) (int64, *utils.RestErr) {
	walletId, err := strconv.ParseInt(wallet_id, 10, 64)
	if err != nil {
		return 0, utils.NewErrorResponse("invalid wallet id", http.StatusBadRequest, "bad_request")
	}
	return walletId, nil
}
