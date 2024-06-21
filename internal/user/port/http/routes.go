package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdangfit/gocommon/validation"

	"main/internal/user/repository"
	"main/internal/user/service"
	"main/pkg/dbs"
	"main/pkg/middleware"
)

func Routes(r *gin.RouterGroup, sqlDB dbs.IDatabase, validator validation.Validation) {
	userRepo := repository.NewUserRepository(sqlDB)
	userSvc := service.NewUserService(validator, userRepo)
	userHandler := NewUserHandler(userSvc)

	authMiddleware := middleware.JWTAuth()
	refreshAuthMiddleware := middleware.JWTRefresh()
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", userHandler.Login)
		authRoute.POST("/register", userHandler.Register)
		authRoute.GET("/me", authMiddleware, userHandler.GetMe)
		authRoute.POST("/refresh-token", refreshAuthMiddleware, userHandler.RefreshToken)
		authRoute.PUT("/update-user", authMiddleware, userHandler.UpdateUser)
		authRoute.PUT("/verfiy-code-email", authMiddleware, userHandler.VerfiyCodeEmail)
		authRoute.PUT("/verfiy-code-phone-number", authMiddleware, userHandler.VerfiyCodePhoneNumber)
		authRoute.PUT("/resend-verfiy-code-phone-number", authMiddleware, userHandler.VerfiyCodePhoneNumberResend)
		authRoute.PUT("/resend-verfiy-code-email", authMiddleware, userHandler.VerfiyCodeEmailResend)
		authRoute.GET("/users", authMiddleware, userHandler.ListUsers)
		authRoute.DELETE("/", authMiddleware, userHandler.DeleteUser)

	}
}
