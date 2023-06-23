package model

import (
	"github.com/jackc/pgtype"
)

// PgLock contains information on locks held.
type PgLock struct {
	Database string `db:"database"`
	DatName  string `db:"datname"`
	Mode     string `db:"mode"`
	Count    int    `db:"count"`
}

// PgStatActivity contains information on tx state.
type PgStatActivity struct {
	Database      string  `db:"database"`
	DatName       string  `db:"datname"`
	State         string  `db:"state"`
	Count         int     `db:"count"`
	MaxTxDuration float64 `db:"max_tx_duration"`
}

// PgStatUserTable contains information on user tables.
type PgStatUserTable struct {
	Database         string             `db:"database"`
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

// PgStatIOUserTable contains I/O information on user tables.
type PgStatIOUserTable struct {
	Database      string `db:"database"`
	SchemaName    string `db:"schemaname"`
	RelName       string `db:"relname"`
	HeapBlksRead  int    `db:"heap_blks_read"`
	HeapBlksHit   int    `db:"heap_blks_hit"`
	IndexBlksRead int    `db:"idx_blks_read"`
	IndexBlksHit  int    `db:"idx_blks_hit"`
	ToastBlksRead int    `db:"toast_blks_read"`
	ToastBlksHit  int    `db:"toast_blks_hit"`
	TidxBlksRead  int    `db:"tidx_blks_read"`
	TidxBlksHit   int    `db:"tidx_blks_hit"`
}

// PgStatUserIndexes contains information on user indexes.
type PgStatUserIndex struct {
	Database      string `db:"database"`
	SchemaName    string `db:"schemaname"`
	RelName       string `db:"relname"`
	IndexRelName  string `db:"indexrelname"`
	IndexScan     int    `db:"idx_scan"`
	IndexTupRead  int    `db:"idx_tup_read"`
	IndexTupFetch int    `db:"idx_tup_fetch"`
}

// PgStatIOUserIndex contains I/O information on user indexes.
type PgStatIOUserIndex struct {
	Database      string `db:"database"`
	SchemaName    string `db:"schemaname"`
	RelName       string `db:"relname"`
	IndexRelName  string `db:"indexrelname"`
	IndexBlksRead int    `db:"idx_blks_read"`
	IndexBlksHit  int    `db:"idx_blks_hit"`
}

// PgStatStatement contains information on statements.
type PgStatStatement struct {
	Database            string  `db:"database"`
	RolName             string  `db:"rolname"`
	DatName             string  `db:"datname"`
	QueryID             int     `db:"queryid"`
	Query               string  `db:"query"`
	Calls               int     `db:"calls"`
	TotalTimeSeconds    float64 `db:"total_time_seconds"`
	MinTimeSeconds      float64 `db:"min_time_seconds"`
	MaxTimeSeconds      float64 `db:"max_time_seconds"`
	MeanTimeSeconds     float64 `db:"mean_time_seconds"`
	StdDevTimeSeconds   float64 `db:"stddev_time_seconds"`
	Rows                int     `db:"rows"`
	SharedBlksHit       int     `db:"shared_blks_hit"`
	SharedBlksRead      int     `db:"shared_blks_read"`
	SharedBlksDirtied   int     `db:"shared_blks_dirtied"`
	SharedBlksWritten   int     `db:"shared_blks_written"`
	LocalBlksHit        int     `db:"local_blks_hit"`
	LocalBlksRead       int     `db:"local_blks_read"`
	LocalBlksDirtied    int     `db:"local_blks_dirtied"`
	LocalBlksWritten    int     `db:"local_blks_written"`
	TempBlksRead        int     `db:"temp_blks_read"`
	TempBlksWritten     int     `db:"temp_blks_written"`
	BlkReadTimeSeconds  int     `db:"blk_read_time_seconds"`
	BlkWriteTimeSeconds int     `db:"blk_write_time_seconds"`
}
