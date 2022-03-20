package controller

import (
	"dk-project-service/entity"
	"dk-project-service/service"
	"dk-project-service/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

type BankAccountController interface {
}

type bankAccountController struct {
	bas service.BankAccountService
}

func NewBankAccountController(bas service.BankAccountService) *bankAccountController {
	return &bankAccountController{bas: bas}
}

func (bc *bankAccountController) GetAllBankRecord(c *gin.Context) {
	idLogin, ok := c.Get("user_id")

	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user not login")))
		return
	}

	res, err := bc.bas.GetByUser(idLogin.(int))

	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, res)
}

func (bc *bankAccountController) UpdateBankRecord(c *gin.Context) {
	idLogin, ok := c.Get("user_id")

	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user not login")))
		return
	}

	var input entity.BankAccountInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := bc.bas.UpdateByID(idLogin.(int), input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update bank account data",
	})
}

func (bc *bankAccountController) InsertNewBankRecord(c *gin.Context) {
	idLogin, ok := c.Get("user_id")

	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user not login")))
		return
	}

	var input entity.BankAccountInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	input.UserId = idLogin.(int)

	err := bc.bas.Insert(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "bank account success record",
	})

}
