package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Dex564/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/Dex564/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/Dex564/golang-todoapp/internal/core/transport/http/server"
	users_transport_http "github.com/Dex564/golang-todoapp/internal/features/users/transoprt/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init applicatoin logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Starting ToDo application!")

	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(nil)
	userRoutes := usersTransportHTTP.Routes()

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(userRoutes...)

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestId(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
