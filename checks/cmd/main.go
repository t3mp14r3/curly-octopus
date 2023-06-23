package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/t3mp14r3/curly-octopus/checks/internal/config"
	"github.com/t3mp14r3/curly-octopus/checks/internal/server"
	"github.com/t3mp14r3/curly-octopus/checks/internal/logger"
	"go.uber.org/zap"
)

func main() {
    config := config.New()
    logger := logger.New()

    ctx, cancel := context.WithCancel(context.Background())

    server := server.New(&config.ServerConfig, logger)

    wg := &sync.WaitGroup{}

    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := server.Run(ctx); err != nil {
            logger.Error("grpc server run error", zap.Error(err))
            cancel()
        }
    }()

    logger.Info("grpc server started", zap.String("port", config.ServerConfig.Port))

    exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

    select {
        case <-ctx.Done():
            logger.Error("grpc server stop via context")
        case <-exit:
            logger.Info("grpc server stop")
    }

    cancel()
    wg.Wait()

    logger.Info("grpc server stopped")
}
