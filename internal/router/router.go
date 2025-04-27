package router

import (
	"github.com/gin-gonic/gin"

	"scaf-gin/internal/controller"
	"scaf-gin/internal/infrastructure/db"
	"scaf-gin/internal/middleware"
	repository "scaf-gin/internal/repository/impl"
	"scaf-gin/internal/service"
)

var gorm = db.NewGormDB()
//var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = repository.NewGormAccountRepository(gorm)

/* DI (Query) */
//var xxxQuery = query.NewXxxQuery(sqlx)

/* DI (Service) */
var accountService = service.NewAccountService(accountRepository)

/* DI (Controller) */
var accountController = controller.NewAccountController(accountService)

func SetApi(r *gin.RouterGroup) {
	r.Use(middleware.ApiErrorHandler())
	r.POST("/accounts/signup", accountController.ApiSignup)
	r.POST("/accounts/login", accountController.ApiLogin)
	r.POST("/accounts/refresh", accountController.ApiRefresh)
	r.GET("/accounts/logout", accountController.ApiLogout)

	auth := r.Group("", middleware.ApiAuth())
	{
		auth.GET("/accounts/me", accountController.ApiGetOne)
		auth.PUT("/accounts/me", accountController.ApiPutOne)
		auth.PUT("/accounts/me/password", accountController.ApiPutPassword)
		auth.DELETE("/accounts/me", accountController.ApiDeleteOne)
	}
}
