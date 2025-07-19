package main

import (
	_ "effective_mobile/docs"
	"effective_mobile/internal/db"
	"effective_mobile/internal/handler"
	"effective_mobile/internal/repository"
	"effective_mobile/internal/service"
	"effective_mobile/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger.Init()
	defer logger.Sync()

	dbPool, err := db.Init()
	if err != nil {
		logger.Error("failed init db pool", zap.Error(err))
		return
	}
	defer db.Close(dbPool)

	repo := repository.NewRepo(dbPool)

	service := service.NewService(repo)

	handler := handler.NewHandler(service, repo)

	router := gin.Default()

	router.POST("/api/subscriptions", handler.CreateSubscriptionHandler)
	router.GET("/api/subscriptions/:id", handler.GetSubscriptionHandler)
	router.PUT("/api/subscriptions/:id", handler.UpdateSubscriptionHandler)
	router.DELETE("/api/subscriptions/:id", handler.DeleteSubscriptionHandler)
	router.GET("/api/subscriptions", handler.ListSubscriptionHandler)
	router.GET("/api/subscriptions/totalcost", handler.GetTotalCostHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info("Server starting...")
	if err := router.Run(":8080"); err != nil {
		logger.Error("Failed starting server: ", zap.Error(err))
		os.Exit(1)
	}

}
