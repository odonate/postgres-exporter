package collectors

import (
	"context"
	"fmt"
	"sync"

	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/prometheus/client_golang/prometheus"
)

// PgStatUserTableCollector collects from pg_stat_user_tables.
type PgStatUserTableCollector struct {
	db    *db.Client
	mutex sync.RWMutex

	seqScan          *prometheus.Desc
	seqTupRead       *prometheus.Desc
	idxScan          *prometheus.Desc
	idxTupFetch      *prometheus.Desc
	nTupIns          *prometheus.Desc
	nTupUpd          *prometheus.Desc
	nTupDel          *prometheus.Desc
	nTupHotUpd       *prometheus.Desc
	nLiveTup         *prometheus.Desc
	nDeadTup         *prometheus.Desc
	nModSinceAnalyze *prometheus.Desc
	lastVacuum       *prometheus.Desc
	lastAutoVacuum   *prometheus.Desc
	lastAnalyze      *prometheus.Desc
	lastAutoAnalyze  *prometheus.Desc
	vacuumCount      *prometheus.Desc
	autoVacuumCount  *prometheus.Desc
	analyzeCount     *prometheus.Desc
	autoAnalyzeCount *prometheus.Desc
}

// NewPgStatUserTableCollector instantiates and returns a new PgStatUserTableCollector.
func NewPgStatUserTableCollector(db *db.Client) *PgStatUserTableCollector {
	variableLabels := []string{"datname", "schemaname", "relname"}
	return &PgStatUserTableCollector{
		db: db,
		seqScan: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "sequential_scan"),
			"Number of sequential scans initiated on this table",
			variableLabels,
			nil,
		),
		seqTupRead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "sequential_scan_tup_read"),
			"Number of live rows fetched by sequential scans",
			variableLabels,
			nil,
		),
		idxScan: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "index_scan"),
			"Number of index scans initiated on this table",
			variableLabels,
			nil,
		),
		idxTupFetch: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "index_tup_fetch"),
			"Number of live rows fetched by index scans",
			variableLabels,
			nil,
		),
		nTupIns: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_tup_ins"),
			"Number of rows inserted",
			variableLabels,
			nil,
		),
		nTupUpd: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_tup_upd"),
			"Number of rows updated",
			variableLabels,
			nil,
		),
		nTupDel: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_tup_del"),
			"Number of rows deleted",
			variableLabels,
			nil,
		),
		nTupHotUpd: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_tup_hot_upd"),
			"Number of rows HOT updated",
			variableLabels,
			nil,
		),
		nLiveTup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_live_tup"),
			"Estimated number of live rows",
			variableLabels,
			nil,
		),
		nDeadTup: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_dead_tup"),
			"Estimated number of dead rows",
			variableLabels,
			nil,
		),
		nModSinceAnalyze: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "n_mod_since_analyze"),
			"Estimated number of rows changed since last analyze",
			variableLabels,
			nil,
		),
		lastVacuum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "last_vacuum"),
			"Last time at which this table was manually vacuumed (not counting VACUUM FULL)",
			variableLabels,
			nil,
		),
		lastAutoVacuum: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "last_autovacuum"),
			"Last time at which this table was vacuumed by the autovacuum daemon",
			variableLabels,
			nil,
		),
		lastAnalyze: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "last_analyze"),
			"Last time at which this table was manually analyzed",
			variableLabels,
			nil,
		),
		lastAutoAnalyze: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "last_autoanalyze"),
			"Last time at which this table was analyzed by the autovacuum daemon",
			variableLabels,
			nil,
		),
		vacuumCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "vacuum_count"),
			"Number of times this table has been manually vacuumed (not counting VACUUM FULL)",
			variableLabels,
			nil,
		),
		autoVacuumCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "autovacuum_count"),
			"Number of times this table has been vacuumed by the autovacuum daemon",
			variableLabels,
			nil,
		),
		analyzeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "analyze_count"),
			"Number of times this table has been manually analyzed",
			variableLabels,
			nil,
		),
		autoAnalyzeCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userTablesSubSystem, "autoanalyze_count"),
			"Number of times this table has been analyzed by the autovacuum daemon",
			variableLabels,
			nil,
		),
	}
}

// Describe implements the prometheus.Collector.
func (c *PgStatUserTableCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.seqScan
	ch <- c.seqTupRead
	ch <- c.idxScan
	ch <- c.idxTupFetch
	ch <- c.nTupIns
	ch <- c.nTupUpd
	ch <- c.nTupDel
	ch <- c.nTupHotUpd
	ch <- c.nLiveTup
	ch <- c.nDeadTup
	ch <- c.nModSinceAnalyze
	ch <- c.lastVacuum
	ch <- c.lastAutoVacuum
	ch <- c.lastAnalyze
	ch <- c.lastAutoAnalyze
	ch <- c.vacuumCount
	ch <- c.autoVacuumCount
	ch <- c.analyzeCount
	ch <- c.autoAnalyzeCount
}

// Collect implements the promtheus.Collector.
func (c *PgStatUserTableCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_ = c.Scrape(ch)
}

// Scrape implements our Scraper interfacc.
func (c *PgStatUserTableCollector) Scrape(ch chan<- prometheus.Metric) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	userTableStats, err := c.db.SelectPgStatUserTables(context.Background())
	if err != nil {
		return fmt.Errorf("user table stats: %w", err)
	}

	for _, stat := range userTableStats {
		ch <- prometheus.MustNewConstMetric(c.seqScan, prometheus.CounterValue, float64(stat.SeqScan), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.seqTupRead, prometheus.CounterValue, float64(stat.SeqTupRead), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.idxScan, prometheus.CounterValue, float64(stat.IndexScan), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.idxTupFetch, prometheus.CounterValue, float64(stat.IndexTupFetch), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nTupIns, prometheus.CounterValue, float64(stat.NTupInsert), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nTupUpd, prometheus.CounterValue, float64(stat.NTupUpdate), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nTupDel, prometheus.CounterValue, float64(stat.NTupDelete), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nTupHotUpd, prometheus.CounterValue, float64(stat.NTupHotUpdate), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nLiveTup, prometheus.GaugeValue, float64(stat.NLiveTup), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nDeadTup, prometheus.GaugeValue, float64(stat.NDeadTup), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.nModSinceAnalyze, prometheus.GaugeValue, float64(stat.NModSinceAnalyze), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.lastVacuum, prometheus.GaugeValue, float64(stat.LastVacuum.Time.UnixMicro()), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.lastAutoVacuum, prometheus.GaugeValue, float64(stat.LastAutoVacuum.Time.UnixMicro()), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.lastAnalyze, prometheus.GaugeValue, float64(stat.LastAnalyze.Time.UnixMicro()), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.lastAutoAnalyze, prometheus.GaugeValue, float64(stat.LastAutoAnalyze.Time.UnixMicro()), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.vacuumCount, prometheus.CounterValue, float64(stat.VacuumCount), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.autoVacuumCount, prometheus.CounterValue, float64(stat.AutoVacuumCount), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.analyzeCount, prometheus.CounterValue, float64(stat.AnalyzeCount), stat.DatName, stat.SchemaName, stat.RelName)
		ch <- prometheus.MustNewConstMetric(c.autoAnalyzeCount, prometheus.CounterValue, float64(stat.AutoAnalyzeCount), stat.DatName, stat.SchemaName, stat.RelName)
	}
	return nil
}