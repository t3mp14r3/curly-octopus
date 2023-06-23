package server

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/t3mp14r3/curly-octopus/main/internal/config"
	"github.com/t3mp14r3/curly-octopus/main/internal/repository"
	"go.uber.org/zap"
)

type Server struct {
    r       *gin.Engine
    addr    string
    logger  *zap.Logger
    repo    *repository.RepoClient
}

func New(serverCongig *config.ServerConfig, repo *repository.RepoClient, logger *zap.Logger) *Server {
    gin.SetMode(gin.ReleaseMode)
    r := gin.New()

    ginLogger, err := zap.NewProduction()

    if err != nil {
        log.Fatalf("failed to initialize zap logger: %v", err)
    }

    r.Use(ginzap.Ginzap(ginLogger, time.RFC3339, true))

    server := &Server{
        r: r,
        addr: serverCongig.Addr,
        logger: logger,
        repo: repo,
    }

    r.GET("/ping", server.Ping)

    return server
}

func (s *Server) Run(ctx context.Context) error {
    errChan := make(chan error, 1)

    wg := &sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer wg.Done()
        if err := s.r.Run(s.addr); err != nil {
            s.logger.Error("server eror", zap.Error(err))
            errChan <- err
        }
    }()

    var err error

    select {
        case <-ctx.Done():
            err = errors.New("server stop via context")
        case err = <-errChan:
    }

    wg.Wait()

    return err
}
