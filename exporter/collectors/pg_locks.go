package collectors

import (
	"context"
	"fmt"
	"sync"

	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/prometheus/client_golang/prometheus"
)

// PgLocksCollector collects from pg_stat_user_tables.
type PgLocksCollector struct {
	db    *db.Client
	mutex sync.RWMutex

	count *prometheus.Desc
}

// NewPgLocksCollector instantiates and returns a new PgLocksCollector.
func NewPgLocksCollector(db *db.Client) *PgLocksCollector {
	variableLabels := []string{"datname", "mode"}
	return &PgLocksCollector{
		db: db,

		count: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, locksSubSystem, "count"),
			"Number of locks",
			variableLabels,
			nil,
		),
	}
}

// Describe implements the prometheus.Collector.
func (c *PgLocksCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.count
}

// Collect implements the promtheus.Collector.
func (c *PgLocksCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_ = c.Scrape(ch)
}

// Scrape implements our Scraper interfacc.
func (c *PgLocksCollector) Scrape(ch chan<- prometheus.Metric) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	locks, err := c.db.SelectPgLocks(context.Background())
	if err != nil {
		return fmt.Errorf("lock stats: %w", err)
	}

	for _, stat := range locks {
		ch <- prometheus.MustNewConstMetric(c.count, prometheus.GaugeValue, float64(stat.Count), stat.DatName, stat.Mode)
	}
	return nil
}
