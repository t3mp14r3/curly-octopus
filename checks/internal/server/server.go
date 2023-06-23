package server

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"

	"github.com/t3mp14r3/curly-octopus/checks/gen"
	"github.com/t3mp14r3/curly-octopus/checks/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
    server      *grpc.Server
    logger      *zap.Logger
    listener    net.Listener
}

type Service struct {
    logger          *zap.Logger
    fontPath        string
    templatePath    string
    storagePath     string
    gen.UnimplementedChecksServer
}

func New(config *config.ServerConfig, logger *zap.Logger) *Server {
    listener, err := net.Listen("tcp", config.Addr)

    if err != nil {
        log.Fatalf("failed to create new tcp listener: %v", err)
    }

    serverRegistar := grpc.NewServer()
    service := &Service{
        logger: logger,
        fontPath: config.FontPath,
        templatePath: config.TemplatePath,
        storagePath: config.StoragePath,
    }

    gen.RegisterChecksServer(serverRegistar, service)

    return &Server{
        server:   serverRegistar,
        logger:   logger,
        listener: listener,
    }
}

func (s *Server) Run(ctx context.Context) error {
    errChan := make(chan error, 1)

    wg := &sync.WaitGroup{}
    wg.Add(1)

    go func() {
        defer wg.Done()
        if err := s.server.Serve(s.listener); err != nil {
            s.logger.Error("grpc server eror", zap.Error(err))
            errChan <- err
        }
    }()

    var err error

    select {
        case <-ctx.Done():
            err = errors.New("grpc server stop via context")
        case err = <-errChan:
    }

    wg.Wait()

    return err
}
