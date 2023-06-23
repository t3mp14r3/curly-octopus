package repository

import (
	"context"
	"database/sql"

	"github.com/t3mp14r3/curly-octopus/main/internal/domain"
	"go.uber.org/zap"
)

func (r *RepoClient) CreateProduct(ctx context.Context, product domain.Product) (*domain.Product, error) {
    query := `INSERT INTO products(name, "desc", cost, barcode, user_id) VALUES($1, $2, $3, $4, $5) RETURNING id, name, "desc", barcode, cost;`

    var result domain.Product
    err := r.db.GetContext(ctx, &result, query, product.Name, product.Desc, product.Cost, product.Barcode, product.UserID)

    if err != nil {
        r.logger.Error("failed to create new product record", zap.Error(err))
    }

    return &result, err
}

func (r *RepoClient) GetUserProducts(ctx context.Context, userID string) ([]*domain.Product, error) {
    query := `SELECT id, name, "desc", barcode, cost FROM products WHERE user_id = $1;`

    var products []*domain.Product
    err := r.db.SelectContext(ctx, &products, query, userID)

    if err == sql.ErrNoRows {
        return nil, nil
    }

    if err != nil {
        r.logger.Error("failed to select user products", zap.Error(err))
        return nil, err
    }

    return products, nil
}

func (r *RepoClient) GetProduct(ctx context.Context, productID string) (*domain.Product, error) {
    query := `SELECT id, name, "desc", cost, barcode, user_id FROM products WHERE id = $1;`

    var product domain.Product
    err := r.db.GetContext(ctx, &product, query, productID)

    if err == sql.ErrNoRows {
        return nil, nil
    }

    if err != nil {
        r.logger.Error("failed to get product record", zap.Error(err))
        return nil, err
    }

    return &product, nil
}

func (r *RepoClient) DeleteProduct(ctx context.Context, productID string) error {
    query := `DELETE FROM products WHERE id = $1;`

    _, err := r.db.ExecContext(ctx, query, productID)

    if err != nil {
        r.logger.Error("failed to delete product record", zap.Error(err))
        return err
    }

    return nil
}
