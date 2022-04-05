package database

import "quik/utils"

type Storage interface {
	GetWallet(int64) (*utils.Wallet, *utils.RestErr)
	IncreaseWalletBalance(int64, int64) *utils.RestErr
}
