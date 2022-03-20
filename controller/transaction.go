package controller

import (
	"dk-project-service/entity"
	"dk-project-service/service"
	"dk-project-service/utils"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	transactionController struct {
		transService service.TransService
	}
)

func NewtransactionController(ts service.TransService) *transactionController {
	return &transactionController{transService: ts}
}

func (tc *transactionController) NewRecord(c *gin.Context) {
	// check from id harus sama dengan yang login
	idLogin, ok := c.Get("user_id")

	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user not login")))
		return
	}

	var input entity.TransInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	if idLogin.(int) != input.FromId {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("unauthorize user, cannot create transaction")))
		return
	}

	// service

	err := tc.transService.NewRecord(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	var res = fmt.Sprintf("success craete transaction from : %d, to %d, date : %v", input.FromId, input.ToId, time.Now())

	c.JSON(201, gin.H{
		"message": res,
	})
}

func (tc *transactionController) TransactionByUser(c *gin.Context) {
	idLogin, ok := c.Get("user_id")

	if !ok {
		c.JSON(401, utils.ErrorMessages(utils.ErrorUnauthorizeUser, errors.New("error user not login")))
		return
	}

	res, err := tc.transService.TransactionByUser(idLogin.(int))
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, res)
}
