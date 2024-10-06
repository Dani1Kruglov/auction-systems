package main

import (
	"context"
	"fmt"
	"net/http"

	"auction-systems/api/protobufs"
	"auction-systems/config"
	"auction-systems/internal/application"
	"auction-systems/internal/domain/repositories"
	g "auction-systems/internal/interfaces/grpc"
	"auction-systems/pkg/database"
	"auction-systems/pkg/logger"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
)

var envFile = ".env"

func main() {
	zapLogger, err := logger.NewLogger()
	if err != nil {
		panic("error initialization logger")
	}

	ctx, contextCancel := context.WithCancel(context.Background())
	defer contextCancel()

	appConfig, err := config.NewConfig(envFile)
	if err != nil {
		zapLogger.Fatal("config initialization error", zap.Error(err))
	}

	pgConn, err := database.NewPostgresClient(appConfig.PostgresDatabaseConfig, zapLogger)
	if err != nil {
		zapLogger.Error("database connection postgres error", zap.Error(err))
		return
	}

	repo, err := repositories.NewRepositories(pgConn)
	if err != nil {
		zapLogger.Fatal("repo initialization error", zap.Error(err))
	}

	services, err := application.NewServices(zapLogger, repo)
	if err != nil {
		zapLogger.Fatal("application initialization error", zap.Error(err))
	}

	mux := runtime.NewServeMux()
	err = protobufs.RegisterAuctionServiceHandlerServer(ctx, mux, g.NewAuctionHandler(zapLogger, services))
	if err != nil {
		zapLogger.Fatal("failed to register gateway", zap.Error(err))
	}

	zapLogger.Info("HTTP Gateway listening", zap.String("host", appConfig.HttpHost), zap.String("port", appConfig.HttpPort))
	if err = http.ListenAndServe(fmt.Sprintf("%s:%s", appConfig.HttpHost, appConfig.HttpPort), mux); err != nil {
		zapLogger.Fatal("error start HTTP server", zap.Error(err))
	}
}
