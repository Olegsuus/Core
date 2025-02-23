package app

import (
	"fmt"
	"github.com/Olegsuus/Core/cmd/config"
	"github.com/Olegsuus/Core/internal/handlers"
	"github.com/Olegsuus/Core/internal/logger"
	"github.com/Olegsuus/Core/internal/service"
	storage "github.com/Olegsuus/Core/internal/storage/postgres"
	"github.com/Olegsuus/Core/pkg/db/postgres"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net"
	"os"
)

type App struct {
	GrpcServer *grpc.Server
	Listener   net.Listener
	DB         *sqlx.DB
	LogFile    *os.File
}

func NewApp(cfg *config.Config, interceptor grpc.UnaryServerInterceptor) (*App, error) {
	dbCfg := postgres.ConfigDB{
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		DBName:   cfg.DB.DBName,
	}

	db, err := postgres.NewConnectDB(dbCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %w", err)
	}

	logFile, err := logger.InitLogger(cfg.Env, cfg.Log.LogFilePath)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to init logger: %w", err)
	}

	l := slog.Default()

	repository := storage.NewRepositoryImpl(db, l)

	services := service.NewServicesImpl(repository)

	handlersGRPC := handlers.NewGRPCHandlers(services, l)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		db.Close()
		logFile.Close()
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	var grpcServer *grpc.Server
	if interceptor != nil {
		grpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	} else {
		grpcServer = grpc.NewServer()
	}

	postpb.RegisterPostServiceServer(grpcServer, handlersGRPC)
	postpb.RegisterUserServiceServer(grpcServer, handlersGRPC)
	postpb.RegisterSubscriptionServiceServer(grpcServer, handlersGRPC)

	return &App{
		GrpcServer: grpcServer,
		Listener:   lis,
		DB:         db,
		LogFile:    logFile,
	}, nil
}

func (a *App) Serve() error {
	log.Printf("gRPC server started on %s", a.Listener.Addr().String())
	return a.GrpcServer.Serve(a.Listener)
}

func (a *App) Stop() {
	a.GrpcServer.Stop()
	if a.Listener != nil {
		a.Listener.Close()
	}
	if a.DB != nil {
		a.DB.Close()
	}
	if a.LogFile != nil {
		a.LogFile.Close()
	}
}
