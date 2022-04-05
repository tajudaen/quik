package mysql

import (
	"database/sql"
	"fmt"

	"quik/config"
	"quik/database"
	"quik/logger"
	"quik/utils"
	"quik/utils/errors"

	"github.com/go-sql-driver/mysql"
)

var (
	client *sql.DB
	MySQL  database.Storage = &mySQL{}
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		config.C.MySQLUserName, config.C.MySQLPassword, config.C.MySQLHost, config.C.MySQLSchema,
	)
	var err error
	client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.Log)
	logger.Info("databases successfully configured", nil)
}

type mySQL struct {
}

func (m *mySQL) GetWallet(wallet_id int64) (*utils.Wallet, *utils.RestErr) {
	stmt, err := client.Prepare("SELECT id, user_id, balance FROM quik.wallets WHERE id=?")

	if err != nil {
		logger.Error("error when trying to prepare get wallet query", err, nil)
		return nil, errors.NewInternalServerError("Server Error")
	}

	result := stmt.QueryRow(wallet_id)
	w := utils.Wallet{}
	if err := result.Scan(&w.ID, &w.UserId, &w.Balance); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("wallet not found")
		}
		logger.Error("error when trying to map wallet record", err, nil)
		return nil, errors.NewInternalServerError("Server Error")
	}
	defer stmt.Close()

	return &w, nil
}

func (m *mySQL) IncreaseWalletBalance(wallet_id int64, amount int64) *utils.RestErr {
	stmt, err := client.Prepare("UPDATE  quik.wallets SET balance = balance + ? WHERE id=?")

	if err != nil {
		logger.Error("error when trying to prepare get wallet query", err, nil)
		return errors.NewInternalServerError("Server Error")
	}

	result, err := stmt.Exec(amount, wallet_id)

	if check, _ := result.RowsAffected(); check < 1 {
		return errors.NewNotFoundError("wallet not found")
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError("wallet not found")
		}
		logger.Error("error when trying to credit wallet", err, nil)
		return errors.NewInternalServerError("Server Error")
	}
	defer stmt.Close()

	return nil
}

func (m *mySQL) DecreaseWalletBalance(wallet_id int64, amount int64) *utils.RestErr {
	stmt, err := client.Prepare("UPDATE  quik.wallets SET balance = balance - ? WHERE id=?")

	if err != nil {
		logger.Error("error when trying to prepare get wallet query", err, nil)
		return errors.NewInternalServerError("Server Error")
	}

	result, err := stmt.Exec(amount, wallet_id)

	if check, _ := result.RowsAffected(); check < 1 {
		return errors.NewNotFoundError("wallet not found")
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.NewNotFoundError("wallet not found")
		}
		logger.Error("error when trying to credit wallet", err, nil)
		return errors.NewInternalServerError("Server Error")
	}
	defer stmt.Close()

	return nil
}