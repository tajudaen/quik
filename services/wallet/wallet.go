package wallet

import (
	"quik/domain/wallet"
	"quik/utils"
	"quik/utils/errors"
)

type WalletService struct {
	wallet.WalletDAOInterface
}

type WalletServiceInterface interface {
	GetWalletBalance(int64) (*int64, *utils.RestErr)
	CreditWallet(int64, int64) *utils.RestErr
	DebitWallet(int64, int64) *utils.RestErr
}

func (s *WalletService) GetWalletBalance(wallet_id int64) (*int64, *utils.RestErr) {

	result, err := s.Get(wallet_id)

	if err != nil {
		return nil, err
	}
	return &result.Balance, nil
}

func (s *WalletService) CreditWallet(wallet_id, amount int64) *utils.RestErr {
	err := s.IncreaseWalletBalance(wallet_id, amount)

	if err != nil {
		return err
	}
	return nil
}

func (s *WalletService) DebitWallet(wallet_id, amount int64) *utils.RestErr {
	result, err := s.Get(wallet_id)

	if err != nil {
		return err
	}

	if amount > result.Balance {
 		return errors.NewInsufficientError("insufficient wallet balance")
	}
	err = s.DecreaseWalletBalance(wallet_id, amount)

	if err != nil {
		return err
	}
	return nil
}