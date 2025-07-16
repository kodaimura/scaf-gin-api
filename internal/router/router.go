package router

import (
	"github.com/gin-gonic/gin"

	"scaf-gin/internal/adapter/db"

	"scaf-gin/internal/module/account"
	"scaf-gin/internal/module/account_profile"

	account_uc "scaf-gin/internal/usecase/account"
	account_profile_uc "scaf-gin/internal/usecase/account_profile"
	auth_uc "scaf-gin/internal/usecase/auth"

	account_h "scaf-gin/internal/handler/account"
	account_profile_h "scaf-gin/internal/handler/account_profile"
	auth_h "scaf-gin/internal/handler/auth"
)

var gorm = db.NewGormDB()

//var sqlx = db.NewSqlxDB()

/* DI (Repository) */
var accountRepository = account.NewRepository()
var accountProfileRepository = account_profile.NewRepository()

/* DI (Service) */
var accountService = account.NewService(accountRepository)
var accountProfileService = account_profile.NewService(accountProfileRepository)

/* DI (Usecase) */
var authUsecase = auth_uc.NewUsecase(gorm, accountService, accountProfileService)
var accountUsecase = account_uc.NewUsecase(gorm, accountService)
var accountProfileUsecase = account_profile_uc.NewUsecase(gorm, accountProfileService)

/* DI (Handler) */
var accountHandler = account_h.NewHandler(accountUsecase)
var accountProfileHandler = account_profile_h.NewHandler(accountProfileUsecase)
var authHandler = auth_h.NewHandler(authUsecase)

func SetApi(r *gin.RouterGroup) {
	r.Use(ApiErrorHandler())
	r.POST("/accounts/signup", authHandler.ApiSignup)
	r.POST("/accounts/login", authHandler.ApiLogin)
	r.POST("/accounts/refresh", authHandler.ApiRefresh)
	r.POST("/accounts/logout", authHandler.ApiLogout)

	auth := r.Group("", ApiAuthMiddleware())
	{
		auth.PUT("/accounts/me/password", authHandler.ApiPutMePassword)
		auth.GET("/accounts/me", accountHandler.ApiGetMe)
		auth.PUT("/accounts/me", accountHandler.ApiPutMe)
		auth.DELETE("/accounts/me", accountHandler.ApiDeleteMe)

		auth.GET("/accounts/me/profile", accountProfileHandler.ApiGetMe)
		auth.PUT("/accounts/me/profile", accountProfileHandler.ApiPutMe)
	}
}
