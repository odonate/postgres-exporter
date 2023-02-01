package model

import (
	"github.com/jackc/pgtype"
)

// PgLock contains information on locks held.
type PgLock struct {
	DatName string `db:"datname"`
	Mode    string `db:"mode"`
	Count   int    `db:"count"`
}

// PgStatActivity contains information on tx state.
type PgStatActivity struct {
	DatName       string  `db:"datname"`
	State         string  `db:"state"`
	Count         int     `db:"count"`
	MaxTxDuration float64 `db:"max_tx_duration"`
}

// PgStatUserTable contiains information on user tables.
type PgStatUserTable struct {
	DatName          string             `db:"datname"`
	SchemaName       string             `db:"schemaname"`
	RelName          string             `db:"relname"`
	SeqScan          int                `db:"seq_scan"`
	SeqTupRead       int                `db:"seq_tup_read"`
	IndexScan        int                `db:"idx_scan"`
	IndexTupFetch    int                `db:"idx_tup_fetch"`
	NTupInsert       int                `db:"n_tup_ins"`
	NTupUpdate       int                `db:"n_tup_upd"`
	NTupDelete       int                `db:"n_tup_del"`
	NTupHotUpdate    int                `db:"n_tup_hot_upd"`
	NLiveTup         int                `db:"n_live_tup"`
	NDeadTup         int                `db:"n_dead_tup"`
	NModSinceAnalyze int                `db:"n_mod_since_analyze"`
	LastVacuum       pgtype.Timestamptz `db:"last_vacuum"`
	LastAutoVacuum   pgtype.Timestamptz `db:"last_autovacuum"`
	LastAnalyze      pgtype.Timestamptz `db:"last_analyze"`
	LastAutoAnalyze  pgtype.Timestamptz `db:"last_autoanalyze"`
	VacuumCount      int                `db:"vacuum_count"`
	AutoVacuumCount  int                `db:"autovacuum_count"`
	AnalyzeCount     int                `db:"analyze_count"`
	AutoAnalyzeCount int                `db:"autoanalyze_count"`
}
