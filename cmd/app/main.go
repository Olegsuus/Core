package main

import (
	"fmt"
	"github.com/Olegsuus/Core/internal/config"
	handlers "github.com/Olegsuus/Core/internal/delivery/post"
	service "github.com/Olegsuus/Core/internal/service/post"
	storage "github.com/Olegsuus/Core/internal/storage/post"
	"github.com/Olegsuus/Core/internal/storage/postgres"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
)

func main() {
	cfg := config.MustConfig()

	pool, err := postgres.NewConnectDB(*cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %w", err)
	}
	defer pool.Close()

	l := slog.Default()

	postStorage := storage.RegisterNewPostStorage(pool)
	postService := service.RegisterPostService(postStorage, l)
	postGRPCHandler := handlers.RegisterPostGRPCHandler(postService)

	addr := fmt.Sprintf(":%d", cfg.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", addr, err)
	}

	grpcServer := grpc.NewServer()

	postpb.RegisterPostServiceServer(grpcServer, postGRPCHandler)

	log.Printf("gRPC server started on %s", addr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
