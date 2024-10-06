package gateway

import (
	"auction-systems/api/protobufs"
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func StartRESTGateway(grpcAddress, restAddress string) error {
	// Создаем контекст
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Создаем новый HTTP mux (мультиплексор)
	mux := runtime.NewServeMux()

	// Опции для соединения с gRPC сервером
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// Подключаем gRPC-сервис к REST Gateway
	err := protobufs.RegisterAuctionServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	// Запуск HTTP сервера для обработки REST запросов
	log.Printf("Starting REST gateway at %s...", restAddress)
	return http.ListenAndServe(restAddress, mux)
}
