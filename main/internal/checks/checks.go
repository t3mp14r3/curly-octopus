package checks

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/t3mp14r3/curly-octopus/checks/gen"
	"github.com/t3mp14r3/curly-octopus/main/internal/config"
)

func New(checksConfig *config.ChecksConfig) *gen.ChecksClient {
    var conn *grpc.ClientConn

    conn, err := grpc.Dial("octopus-checks:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

    if err != nil {
        log.Fatalf("failed to dial to the checks grpc server: %v", err)
    }

    client := gen.NewChecksClient(conn)

    return &client
}
