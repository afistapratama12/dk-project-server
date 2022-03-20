package service

import (
	"dk-project-service/entity"
	"dk-project-service/repository"
	"fmt"
	"strconv"
)

type (
	WdService interface {
		GetAllWdReq() ([]entity.WdReqDetail, error)
		GetWdReqWeek() ([]entity.WdReqDetail, error)

		GetAllWdReqByUserID(userId string) ([]entity.WdReqModel, error)

		WdReqRoBalance(input entity.WdReqInput) error
		WdReqMoneyBalance(input entity.WdReqInput) error
	}

	wdService struct {
		wdRepo   repository.WdRepo
		userRepo repository.UserRepository
	}
)

func NewWdService(wdRepo repository.WdRepo, userRepo repository.UserRepository) *wdService {
	return &wdService{
		wdRepo:   wdRepo,
		userRepo: userRepo,
	}
}

func (s *wdService) GetAllWdReq() ([]entity.WdReqDetail, error) {
	return s.wdRepo.GetAllWd()
}

func (s *wdService) GetWdReqWeek() ([]entity.WdReqDetail, error) {
	return s.wdRepo.GetAllWdInWeek()
}

func (s *wdService) GetAllWdReqByUserID(userId string) ([]entity.WdReqModel, error) {
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	return s.wdRepo.GetAllWdByUserID(userIdInt)
}

func (s *wdService) WdReqMoneyBalance(input entity.WdReqInput) error {
	// pengecekan money dari front end,
	// kalau saldo money lebih kecil dari yang ditarik,
	// berarti munculkan error di FE

	// check ada wd req atau nggk
	recordWdReqWeek, err := s.wdRepo.GetWdReqInWeekByUserID(input.UserId)
	if err != nil {
		return err
	}

	if recordWdReqWeek.Id == 0 && recordWdReqWeek.UserId == 0 && recordWdReqWeek.BankAccId == 0 {
		var newWdReq = entity.WdReqModel{
			UserId:    input.UserId,
			BankAccId: input.BankAccId,
		}

		if input.Moneybalance != 0 {
			newWdReq.MoneyBalance = input.Moneybalance
		}

		err = s.wdRepo.CreateWdReq(newWdReq)
		if err != nil {
			return err
		}
	} else {
		recordWdReqWeek.MoneyBalance += input.Moneybalance

		err = s.wdRepo.UpdateWdReqByID(recordWdReqWeek)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *wdService) WdReqRoBalance(input entity.WdReqInput) error {
	// check ro balance by front end aja
	// kalau request RO melebihi saldo , dikasih error langusng di FE

	// check ada WD req atau nggk
	recordWdReqWeek, err := s.wdRepo.GetWdReqInWeekByUserID(input.UserId)
	if err != nil {
		return err
	}

	// RUMUS :
	// data: (a) ro req, (b) ro akumulasi history, (c) ro dari upline / downline
	// si upline cair 2
	// si down 1
	// a = 1, b = 0, c = 2
	// a check c - b, jika a < c (masukkan a), jika c <= a (masukkan c), jika b >= c (tolak)
	// 1 check 2 - 0, a < b = masukkan a

	// check kebawah
	// si u4 cair 2
	// a = 1, b = 0, c = 2
	// a check c - b, a = 1, c = 2, a < c kita dapat 3000

	// si u4 cair 2
	// a = 1, b = 0, c = 1
	// a check c - b, a = 1, c = 1, c <= a kita dapat 3000

	err = s.RoBonusNetworkUpline(input.UserId, recordWdReqWeek.RoBalance, input.RoBalance)
	if err != nil {
		return err
	}

	getBonus, err := s.GetTotalBonusNetworkDL([]int{input.UserId}, recordWdReqWeek.RoBalance, input.RoBalance)
	if err != nil {
		return err
	}

	// Kalau nggk ada create baru dengan logic yang ada
	if recordWdReqWeek.Id == 0 && recordWdReqWeek.UserId == 0 && recordWdReqWeek.BankAccId == 0 {
		var newWdReq = entity.WdReqModel{
			UserId:         input.UserId,
			BankAccId:      input.BankAccId,
			RoBalance:      input.RoBalance,
			RoMoneyBalance: (input.RoBalance * entity.BonusUser) + getBonus,
			Approved:       false,
		}

		err = s.wdRepo.CreateWdReq(newWdReq)
		if err != nil {
			return err
		}
	} else {
		recordWdReqWeek.RoBalance += input.RoBalance
		recordWdReqWeek.RoMoneyBalance += (input.RoBalance * entity.BonusUser) + getBonus

		err = s.wdRepo.UpdateWdReqByID(recordWdReqWeek)
		if err != nil {
			return err
		}
	}

	return nil
}

// for update ro bonus jaringan upline
func (s *wdService) RoBonusNetworkUpline(baseParentId int, inputRoAccumUser int, inputRoNowUser int) error {
	parentId := baseParentId

	// looping ke atas, bonus untuk upline
	for {
		// perlu 2 data
		user, err := s.userRepo.GetuserId(parentId)
		if err != nil {
			fmt.Println("RoBonusNetworkUpline: error loop get data user parent bonus jaringan upline")
			return err
		}

		parentWdReqWeek, err := s.wdRepo.GetWdReqInWeekByUserID(user.ParentId)
		if err != nil {
			fmt.Println("RoBonusNetworkUpline: error loop bonus wd ro upline")
			return err
		}

		if parentWdReqWeek.Id != 0 && parentWdReqWeek.UserId != 0 && parentWdReqWeek.BankAccId != 0 {
			roCheck := parentWdReqWeek.RoBalance
			roAccumUser := inputRoAccumUser

			roNowUser := inputRoNowUser

			if roAccumUser < roCheck {
				// ro upline dikurangi akum
				roCheck -= roAccumUser

				if roNowUser < roCheck {
					parentWdReqWeek.RoMoneyBalance += (roNowUser * entity.BonusJaringan)
				} else if roCheck <= roNowUser {
					parentWdReqWeek.RoMoneyBalance += (roCheck * entity.BonusJaringan)
				}

				err = s.wdRepo.UpdateWdReqByID(parentWdReqWeek)
				if err != nil {
					fmt.Println("RoBonusNetworkUpline: error UPDATE sql loop update wd ro upline")
					return err
				}
			}
		}

		if user.ParentId == 1 {
			break
		} else {
			parentId = user.ParentId
		}
	}

	return nil
}

// for get total bonus jaringan kita, dari check ke downline
func (s *wdService) GetTotalBonusNetworkDL(baseListId []int, inputRoAccumUser int, inputRoNowUser int) (int, error) {
	var bonus = 0

	var downlineCheckId []int = baseListId

	// loop untuk child user
	for {
		var allNextDlIds []int

		for _, dlId := range downlineCheckId {
			allDownlineUser, err := s.userRepo.CheckUserId(dlId)
			if err != nil {
				fmt.Println("GetTotalBonusNetworkDL: query all downline (left, center, right) loop")
				return 0, err
			}

			if len(allDownlineUser) < 1 {
				continue
			} else {
				var nextDlIds []int

				for _, dlUser := range allDownlineUser {
					dlWdReqWeek, err := s.wdRepo.GetWdReqInWeekByUserID(dlUser.Id)
					if err != nil {
						fmt.Println("GetTotalBonusNetworkDL: error loop bonus wd ro downline fo us")
						return 0, err
					}

					if dlWdReqWeek.Id != 0 && dlWdReqWeek.UserId != 0 && dlWdReqWeek.BankAccId != 0 {
						roDl := dlWdReqWeek.RoBalance
						roAccumUser := inputRoAccumUser

						roNowUser := inputRoNowUser

						if roAccumUser < roDl {
							roDl -= roAccumUser

							if roNowUser < roDl {
								bonus += (roNowUser * entity.BonusJaringan)
							} else if roDl <= roNowUser {
								bonus += (roDl * entity.BonusJaringan)
							}
						}
					}

					nextDlIds = append(nextDlIds, dlUser.Id)
				}

				allNextDlIds = append(allNextDlIds, nextDlIds...)
			}
		}

		if len(allNextDlIds) < 1 {
			break
		} else {
			downlineCheckId = allNextDlIds
		}
	}

	return bonus, nil
}
