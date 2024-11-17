package database

import (
	"context"

	"gorm.io/gorm"
)

// BaseRepository defines the base structure with common transaction handling
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository creates a new instance of BaseRepository
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

// BeginTransactionAndSetContext starts a transaction, updates the context, and returns the new context
func (r *BaseRepository) BeginTransactionAndSetContext(ctx context.Context) (context.Context, error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return ctx, tx.Error
	}

	// Update context with the transaction
	return WithTx(ctx, tx), nil
}

// CommitTransaction commits the transaction in the context
func (r *BaseRepository) CommitTransaction(ctx context.Context) error {
	tx := GetTx(ctx, r.db)
	return tx.Commit().Error
}

// RollbackTransaction rolls back the transaction in the context
func (r *BaseRepository) RollbackTransaction(ctx context.Context) {
	tx := GetTx(ctx, r.db)
	tx.Rollback()
}
