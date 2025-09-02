package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TxFunc func(pgx.Tx) error

type TxManager struct {
	pool *pgxpool.Pool
}

func NewTxManager(pool *pgxpool.Pool) *TxManager {
	return &TxManager{pool: pool}
}

// Run executes a function inside a transaction
func (m *TxManager) Run(ctx context.Context, fn TxFunc) error {
	tx, err := m.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx) // safe rollback (no-op if committed)

	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
