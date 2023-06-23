package repository

import (
	"context"
	"database/sql"

	"github.com/t3mp14r3/curly-octopus/main/internal/domain"
	"go.uber.org/zap"
)

func (r *RepoClient) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
    query := `INSERT INTO users(login, email, password, name) VALUES($1, $2, $3, $4) RETURNING id, login, email, password, name;`

    var result domain.User
    err := r.db.GetContext(ctx, &result, query, user.Login, user.Email, user.Password, user.Name)

    if err != nil {
        r.logger.Error("failed to create new user record", zap.Error(err))
    }

    return &result, err
}

func (r *RepoClient) EmailUsed(ctx context.Context, email string) bool {
    query := `SELECT COUNT(id) FROM users WHERE email = $1;`

    var count int
    err := r.db.GetContext(ctx, &count, query, email)

    if err != nil && err != sql.ErrNoRows {
        r.logger.Error("failed to check if email is used", zap.Error(err))
    }

    if count > 0 {
        return true
    }

    return false
}

func (r *RepoClient) LoginUsed(ctx context.Context, login string) bool {
    query := `SELECT COUNT(id) FROM users WHERE login = $1;`

    var count int
    err := r.db.GetContext(ctx, &count, query, login)

    if err != nil && err != sql.ErrNoRows {
        r.logger.Error("failed to check if login is used", zap.Error(err))
    }

    if count > 0 {
        return true
    }

    return false
}

func (r *RepoClient) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
    query := `SELECT id, login, email, password, name FROM users WHERE login = $1;`

    var user domain.User
    err := r.db.GetContext(ctx, &user, query, login)

    if err != nil {
        r.logger.Error("failed to get user record by login", zap.Error(err))
        return nil, err
    }

    return &user, nil
}

func (r *RepoClient) GetUser(ctx context.Context, userID string) (*domain.User, error) {
    query := `SELECT id, login, email, password, name FROM users WHERE id = $1;`

    var user domain.User
    err := r.db.GetContext(ctx, &user, query, userID)

    if err != nil {
        r.logger.Error("failed to get user record", zap.Error(err))
        return nil, err
    }

    return &user, nil
}
