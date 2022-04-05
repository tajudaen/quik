package database

import "quik/utils"


type Storage interface {
	GetWallet(int64) (*utils.Wallet, *utils.RestErr)
}
