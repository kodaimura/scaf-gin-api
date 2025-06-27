package router

import (
	"github.com/gin-gonic/gin"

	"scaf-gin/internal/domain/account"
	"scaf-gin/internal/domain/auth"
	"scaf-gin/internal/infrastructure/db"
	"scaf-gin/internal/middleware"
)

var gorm = db.NewGormDB()

//var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = account.NewRepository()

/* DI (Service) */
var authService = auth.NewService(accountRepository)
var accountService = account.NewService(accountRepository)

/* DI (Controller) */
var authController = auth.NewController(gorm, authService)
var accountController = account.NewController(gorm, accountService)

func SetApi(r *gin.RouterGroup) {
	r.Use(middleware.ApiErrorHandler())
	r.POST("/accounts/signup", authController.ApiSignup)
	r.POST("/accounts/login", authController.ApiLogin)
	r.POST("/accounts/refresh", authController.ApiRefresh)
	r.POST("/accounts/logout", authController.ApiLogout)

	auth := r.Group("", middleware.ApiAuth())
	{
		auth.GET("/accounts/me", accountController.ApiGetMe)
		auth.PUT("/accounts/me", accountController.ApiPutMe)
		auth.PUT("/accounts/me/password", authController.ApiPutMePassword)
		auth.DELETE("/accounts/me", accountController.ApiDeleteMe)
	}
}
