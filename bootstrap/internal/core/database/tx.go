package database

import (
	"context"

	"gorm.io/gorm"
)

// txKey is an unexported key type for storing transaction in context
type txKey struct{}

// WithTx adds the transaction to the context
func WithTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// GetTx retrieves the transaction from the context
// If no transaction is found, it returns the provided defaultDB instance
func GetTx(ctx context.Context, defaultDB *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok && tx != nil {
		return tx
	}
	return defaultDB
}
