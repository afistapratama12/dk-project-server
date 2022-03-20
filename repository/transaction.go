package repository

import (
	"dk-project-service/entity"
	"strconv"

	"gorm.io/gorm"
)

type (
	TransRepo interface {
		InsertTrans(trans entity.TransInput) error
		GetTransactionById(id int) ([]entity.Transaction, error)
	}

	transRepo struct {
		db *gorm.DB
	}
)

func NewTransRepo(db *gorm.DB) *transRepo {
	return &transRepo{db: db}
}

func (r *transRepo) InsertTrans(trans entity.TransInput) error {
	var query string

	if trans.SASBalance != 0 {
		query = `INSERT INTO transactions (from_id, to_id, sas_balance) VALUES (?, ?, ?)`

		if err := r.db.Exec(query, trans.FromId, trans.ToId, trans.SASBalance).Error; err != nil {
			return err
		}
	}

	if trans.ROBalance != 0 {
		query = `INSERT INTO transactions (from_id, to_id, ro_balance) VALUES (?, ?, ?)`

		if err := r.db.Exec(query, trans.FromId, trans.ToId, trans.ROBalance).Error; err != nil {
			return err
		}
	}

	if trans.MoneyBalance != 0 {
		query = `INSERT INTO transactions (from_id, to_id, money_balance) VALUES (?, ?, ?)`

		if err := r.db.Exec(query, trans.FromId, trans.ToId, trans.MoneyBalance).Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *transRepo) GetTransactionById(id int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	idStr := strconv.Itoa(id)

	if err := r.db.Raw("SELECT id, from_id, to_id, sas_balance, ro_balance, money_balance FROM transactions WHERE from_id = ? OR to_id = ?", idStr, idStr).Scan(&transactions).Error; err != nil {
		return transactions, err
	}

	return transactions, nil
}
