package app

import (
	"log/slog"
	"net/http"

	"restapi/internal/config"
	"restapi/internal/http-server/handlers"
	"restapi/internal/http-server/middleware"
	"restapi/internal/http-server/middleware/logger"
	"restapi/internal/storage"
	"github.com/gin-gonic/gin"
)

func App (storage storage.Storage, log *slog.Logger, cfg *config.Config) error {
	server := setupServer(*cfg, log, storage)
	log.Info("Serving on address", slog.String("address", cfg.Address))
	return server.ListenAndServe()
}

func setupRouter (db storage.Storage, log *slog.Logger, cfg config.ServiceAddresses) *gin.Engine {
	router := gin.Default()

	middleware.LoadRouterWithMiddleware(router, 
		middleware.CorsWithConfig(cfg), 
		logger.URLFormat(),
		logger.New(log),
		middleware.RequestIDMiddleware(),
	)
	
	appHandlers := handlers.NewHandlers(db, log)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	privateRoute := router.Group("/user")
	privateRoute.Use(middleware.AllowInternalRequests(log))
	{
		privateRoute.POST("/", appHandlers.User.SaveUser)
	}

	publicProtectedRoute := router.Group("")
	publicProtectedRoute.Use(middleware.JWNAuthMiddleware())
	{
		userRouter := publicProtectedRoute.Group("/user")
		{
			userRouter.GET("", appHandlers.User.GetUser)
			userRouter.PUT("/password", appHandlers.User.UpdateUserPassword)
			userRouter.DELETE("", appHandlers.User.DeleteUser)
		}

		taskRouter := publicProtectedRoute.Group("/tasks")
		{
			taskRouter.POST("", appHandlers.Task.SaveTask)
			taskRouter.GET("", appHandlers.Task.GetTasksByUserID)
			taskRouter.GET("/:taskId", appHandlers.Task.GetTaskByTaskID)
			taskRouter.PUT("/:taskId", appHandlers.Task.UpdateTask)
			taskRouter.DELETE("/:taskId", appHandlers.Task.DeleteTask)
		}
	}

	return router
}

func setupServer(cfg config.Config, log *slog.Logger, db storage.Storage) *http.Server {
	router := setupRouter(db, log, cfg.ServiceAddresses)
	log.Info("Router was set up")

	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IddleTimeout,
	}

	return server
}
