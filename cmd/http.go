package cmd

import (
	"ewallet-ums/cmd/middleware"
	"ewallet-ums/external"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ServeHttp() {
	dependencyInject := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependencyInject.HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependencyInject.RegisterApi.Register)
	userV1.POST("/login", dependencyInject.LoginApi.Login)
	userV1.DELETE("/logout", middleware.MiddlewareValidateAuth, dependencyInject.LogoutApi.Logout)
	userV1.PUT("/refresh-token", middleware.MiddlewareValidateAuth, dependencyInject.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		logrus.Fatal("err")
	}
}

type Dependency struct {
	UserRepository  interfaces.IUserRepository
	HealthcheckAPI  interfaces.IHealthcheckHandler
	RegisterApi     interfaces.IRegisterHandler
	LoginApi        interfaces.ILoginHandler
	LogoutApi       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidationAPI *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	ext := &external.External{}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		UserRepo: userRepo,
		External: ext,
	}

	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}

	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}

	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	refreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}

	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	tokenValidationSvc := &services.TokenValidationService{
		UserRepo: userRepo,
	}

	tokenValidationAPI := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationSvc,
	}

	return Dependency{
		UserRepository:     userRepo,
		HealthcheckAPI:     healthcheckAPI,
		RegisterApi:        registerAPI,
		LoginApi:           loginAPI,
		LogoutApi:          logoutAPI,
		RefreshTokenAPI:    refreshTokenAPI,
		TokenValidationAPI: tokenValidationAPI,
	}
}
