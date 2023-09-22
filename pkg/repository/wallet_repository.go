package repository

import (
	"database/sql"
	"stock-simulation/pkg/model"
)

type WalletRepository struct {
	*SQLRepository
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{
		SQLRepository: NewSQLRepository(db),
	}
}

func (r *WalletRepository) FindByUserID(userID string) (*model.Wallet, error) {
	rows, err := r.FindBy("wallets", "user_id", userID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	if !rows.Next() {
		return nil, sql.ErrNoRows
	}

	var wallet model.Wallet
	if err := rows.Scan(&wallet.UserID, &wallet.Currency, &wallet.Balance); err != nil {
		return nil, err
	}

	return &wallet, nil

}
