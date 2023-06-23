package collectors

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
)

// PgStatUserIndexesCollector collects from pg_stat_user_indexes.
type PgStatUserIndexesCollector struct {
	dbClients []*db.Client
	mutex     sync.RWMutex

	idxScan          *prometheus.Desc
	idxTupRead       *prometheus.Desc
	idxTupFetch      *prometheus.Desc
}

// NewPgStatUserIndexesCollector instantiates and returns a new PgStatUserTableCollector.
func NewPgStatUserIndexesCollector(dbClients []*db.Client) *PgStatUserIndexesCollector {
	variableLabels := []string{"database", "schemaname", "relname", "indexrelname"}
	return &PgStatUserIndexesCollector{
		dbClients: dbClients,
		idxScan: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userIndexesSubSystem, "index_scan"),
			"Number of index scans initiated on this index",
			variableLabels,
			nil,
		),
		idxTupRead: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userIndexesSubSystem, "index_tup_read"),
			"Number of index entries returned by scans on this index",
			variableLabels,
			nil,
		),
		idxTupFetch: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, userIndexesSubSystem, "index_tup_fetch"),
			"Number of live table rows fetched by simple index scans using this index",
			variableLabels,
			nil,
		),
	}
}

// Describe implements the prometheus.Collector.
func (c *PgStatUserIndexesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.idxScan
	ch <- c.idxTupRead
	ch <- c.idxTupFetch
}

// Collect implements the promtheus.Collector.
func (c *PgStatUserIndexesCollector) Collect(ch chan<- prometheus.Metric) {
	_ = c.Scrape(ch)
}

// Scrape implements our Scraper interface.
func (c *PgStatUserIndexesCollector) Scrape(ch chan<- prometheus.Metric) error {
	start := time.Now()
	defer func() {
		log.Infof("user table scrape took %dms", time.Now().Sub(start).Milliseconds())
	}()
	group := errgroup.Group{}
	for _, dbClient := range c.dbClients {
		dbClient := dbClient
		group.Go(func() error { return c.scrape(dbClient, ch) })
	}
	if err := group.Wait(); err != nil {
		return fmt.Errorf("scraping: %w", err)
	}
	return nil
}


func (c *PgStatUserIndexesCollector) scrape(dbClient *db.Client, ch chan<- prometheus.Metric) error {
	userIndexStats, err := dbClient.SelectPgStatUserIndexes(context.Background())
	if err != nil {
		return fmt.Errorf("user table stats: %w", err)
	}
	for _, stat := range userIndexStats {
		ch <- prometheus.MustNewConstMetric(c.idxScan, prometheus.CounterValue, float64(stat.IndexScan), stat.Database, stat.SchemaName, stat.RelName, stat.IndexRelName)
		ch <- prometheus.MustNewConstMetric(c.idxTupRead, prometheus.CounterValue, float64(stat.IndexTupRead), stat.Database, stat.SchemaName, stat.RelName, stat.IndexRelName)
		ch <- prometheus.MustNewConstMetric(c.idxTupFetch, prometheus.CounterValue, float64(stat.IndexTupFetch), stat.Database, stat.SchemaName, stat.RelName, stat.IndexRelName)
	}
	return nil
}
