package db

import (
	"context"

	"github.com/odonate/postgres-exporter/exporter/db/model"
)

const sqlSelectPgStatioUserTables = `
SELECT
    schemaname,
    relname,
    heap_blks_read,
    heap_blks_hit,
    idx_blks_read,
    idx_blks_hit,
    toast_blks_read,
    toast_blks_hit,
    tidx_blks_read,
    tidx_blks_hit
FROM pg_statio_user_tables`

// SelectPgStatioUserTables selects stats on user tables.
func (db *Client) SelectPgStatioUserTables(ctx context.Context) ([]*model.PgStatioUserTable, error) {
	pgStatioUserTables := []*model.PgStatioUserTable{}
	if err := db.Select(ctx, &pgStatioUserTables, sqlSelectPgStatioUserTables); err != nil {
		return nil, err
	}
	return pgStatioUserTables, nil
}
