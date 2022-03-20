package routes

import (
	"dk-project-service/controller"
	"dk-project-service/repository"
	"dk-project-service/service"

	"github.com/gin-gonic/gin"
)

var (
	tsRepo       = repository.NewTransRepo(DB)
	tsService    = service.NewTransService(tsRepo, userRepo)
	tsController = controller.NewtransactionController(tsService)
)

func TransactionRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/transaction", MainMiddleware, tsController.TransactionByUser)
		v1.POST("/transaction/record", MainMiddleware, tsController.NewRecord)
	}
}
