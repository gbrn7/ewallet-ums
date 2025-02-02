package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func bootstrapEngin() {

}

func ServeHttp() {
	healthcheckSvc := &services.Healthcheck{}
	HealthcheckAPI := api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		RegisterRepo: registerRepo,
	}

	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	r := gin.Default()

	r.GET("/health", HealthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerAPI.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		logrus.Fatal("err")
	}
}
