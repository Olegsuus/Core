package main

import (
	"fmt"
	"github.com/Olegsuus/Core/cmd/config"
	"github.com/Olegsuus/Core/internal/app"
	"github.com/Olegsuus/Core/internal/metrics"
	"log"
)

func main() {
	cfg := config.MustConfig()

	metricsAddr := fmt.Sprintf(":%d", cfg.Metrics.Port)
	go metrics.StartHTTPServer(metricsAddr)

	appInstance, err := app.NewApp(cfg, metrics.UnaryServerInterceptor())
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	if err := appInstance.Serve(); err != nil {
		log.Printf("failed to connect gRPC server: %v", err)
	}
	defer appInstance.Stop()
}
