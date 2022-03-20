package service

import (
	"dk-project-service/entity"
	"dk-project-service/repository"
	"fmt"
)

type (
	TransService interface {
		NewRecord(input entity.TransInput) error
		TransactionByUser(id int) ([]entity.Transaction, error)

		NewDownline(inputUplineId int) error
	}

	transService struct {
		transRepo repository.TransRepo
		userRepo  repository.UserRepository
	}
)

func NewTransService(tr repository.TransRepo, ur repository.UserRepository) *transService {
	return &transService{
		transRepo: tr,
		userRepo:  ur,
	}
}

func (s *transService) NewRecord(input entity.TransInput) error {
	userFrom, err := s.userRepo.GetuserId(input.FromId)
	if err != nil {
		fmt.Println("error inserting transaction, NewRecord, line 28")
		return err
	}

	userTo, err := s.userRepo.GetuserId(input.ToId)
	if err != nil {
		fmt.Println("error inserting transaction, NewRecord, line 34")
		return err
	}

	if input.SASBalance != 0 {
		if userFrom.Role == "user" {
			if userFrom.SASBalance == 0 {
				return fmt.Errorf("error transaction, balance user %v, SASBalance : 0", input.FromId)
			} else {
				userFrom.SASBalance -= input.SASBalance
			}
		}

		userTo.SASBalance += input.SASBalance
	}

	if input.ROBalance != 0 {
		if userFrom.Role == "user" {
			if userFrom.ROBalance == 0 {
				return fmt.Errorf("error transaction, balance user %v, ROBalance 0", input.FromId)
			} else {
				userFrom.ROBalance -= input.ROBalance
			}
		}

		userTo.ROBalance += input.ROBalance
	}

	if input.MoneyBalance != 0 {
		if userFrom.Role == "user" {
			if userFrom.MoneyBalance == 0 {
				return fmt.Errorf("error transaction, balance user %v, MoneyBalance 0", input.FromId)
			} else {
				userFrom.MoneyBalance -= input.MoneyBalance
			}
		}

		userTo.MoneyBalance += input.MoneyBalance
	}

	err = s.userRepo.UpdateBalance(userFrom)
	if err != nil {
		fmt.Println("error inserting transaction, NewRecord, line 67")
		return err
	}

	err = s.userRepo.UpdateBalance(userTo)
	if err != nil {
		fmt.Println("error inserting transaction, NewRecord, line 73")
		return err
	}

	err = s.transRepo.InsertTrans(input)
	if err != nil {
		fmt.Println("error inserting transaction, NewRecord, line 79")
		return err
	}

	return nil
}

func (s *transService) TransactionByUser(id int) ([]entity.Transaction, error) {
	return s.transRepo.GetTransactionById(id)
}

func (s *transService) NewDownline(inputUplineId int) error {
	var uplineId = inputUplineId

	for i := 0; i < 5; i++ {
		user, err := s.userRepo.GetuserId(uplineId)
		if err != nil {
			fmt.Println("error inserting transaction, NewDownline, line 112")
			return err
		}

		if user.Id != 0 && user.Role != "admin" {
			var getMoney = 0

			if i == 0 {
				getMoney = 5000
			} else {
				getMoney = 3000
			}

			user.MoneyBalance += getMoney

			err := s.userRepo.UpdateBalance(user)
			if err != nil {
				fmt.Println("error inserting transaction, NewDonwline, line 127")
				return err
			}

			transInput := entity.TransInput{
				FromId:       1,
				ToId:         user.Id,
				MoneyBalance: getMoney,
			}

			err = s.transRepo.InsertTrans(transInput)
			if err != nil {
				fmt.Println("error inserting transaction, NewDonwline, line 139")
				return err
			}

			uplineId = user.ParentId
		} else {
			break
		}
	}

	return nil
}
