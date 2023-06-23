package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/t3mp14r3/curly-octopus/main/internal/auth"
	"github.com/t3mp14r3/curly-octopus/main/internal/config"
	"github.com/t3mp14r3/curly-octopus/main/internal/logger"
	"github.com/t3mp14r3/curly-octopus/main/internal/repository"
	"github.com/t3mp14r3/curly-octopus/main/internal/server"
	"go.uber.org/zap"
)

func main() {
    config := config.New()
    logger := logger.New()

    ctx, cancel := context.WithCancel(context.Background())

    repo := repository.New(&config.PostgresConfig, logger)
    defer repo.Close()
    
    auth := auth.New(&config.AuthConfig, logger)

    server := server.New(&config.ServerConfig, repo, auth, logger)

    wg := &sync.WaitGroup{}

    wg.Add(1)
    go func() {
        defer wg.Done()
        if err := server.Run(ctx); err != nil {
            logger.Error("server run error", zap.Error(err))
            cancel()
        }
    }()

    logger.Info("server started", zap.String("addr", config.ServerConfig.Addr))

    exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

    select {
        case <-ctx.Done():
            logger.Error("main server stop via context")
        case <-exit:
            logger.Info("main server stop")
    }

    cancel()
    wg.Wait()

    logger.Info("main server stopped")
}
