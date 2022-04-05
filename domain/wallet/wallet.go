package wallet

import (
	"quik/database"
	"quik/utils"
)

type Wallet struct {
	ID      int64  `json:"id"`
	UserId  string `json:"user_id"`
	Balance int64  `json:"balance"`
}

type Wallet_DAO struct {
	database.Storage
}

type WalletDAOInterface interface {
	Get(int64) (*Wallet, *utils.RestErr)
}

func (w *Wallet_DAO) Get(wallet_id int64) (*Wallet, *utils.RestErr) {

	wallet, err := w.GetWallet(wallet_id)

	if err != nil {
		return nil, err
	}
	return &Wallet{ID: wallet.ID, UserId: wallet.UserId, Balance: wallet.Balance}, nil

}
