package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"github.com/t3mp14r3/curly-octopus/main/internal/config"
)

type RepoClient struct {
    db      *sqlx.DB
    logger  *zap.Logger
}

func New(postgresConfig *config.PostgresConfig, logger *zap.Logger) *RepoClient {
    conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Name,
	)
    
    postgresConn, err := sqlx.Connect("postgres", conn)

	if err != nil {
        log.Fatalf("failed to initialize postgres connection: %v", err)
	}

    migrate(postgresConn.DB)

    return &RepoClient{
        db:     postgresConn,
        logger: logger,
    }
}

func migrate(db *sql.DB) {
    if err := goose.SetDialect("postgres"); err != nil {
        log.Fatalf("failed to set goose dialect: %v", err)
    }

    if err := goose.Up(db, "migrations"); err != nil {
        log.Fatalf("failed to migrate the database: %v", err)
    }
}

func (r *RepoClient) Close() {
    if err := r.db.Close(); err != nil {
        log.Fatalf("error while closing postgres connection: %v", err)
    }
}
