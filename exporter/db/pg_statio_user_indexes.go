package db

import (
	"context"

	"github.com/odonate/postgres-exporter/exporter/db/model"
)

const sqlSelectPgStatioUserIndexes = `
SELECT
    schemaname,
    relname,
    indexrelname,
    idx_blks_read,
    idx_blks_hit
FROM pg_statio_user_indexes`

// SelectPgStatioUserIndexes selects stats on user indexes.
func (db *Client) SelectPgStatioUserIndexes(ctx context.Context) ([]*model.PgStatioUserIndex, error) {
	pgStatioUserIndexes := []*model.PgStatioUserIndex{}
	if err := db.Select(ctx, &pgStatioUserIndexes, sqlSelectPgStatioUserIndexes); err != nil {
		return nil, err
	}
	return pgStatioUserIndexes, nil
}
