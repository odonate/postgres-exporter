package collector

import (
	"context"
	"fmt"
	"sync"

	"github.com/odonate/postgres-exporter/db"
	"github.com/prometheus/client_golang/prometheus"
)

// PgStatActivityCollector collects from pg_stat_user_tables.
type PgStatActivityCollector struct {
	db    *db.Client
	mutex sync.RWMutex

	activityCount *prometheus.Desc
	maxTxDuration *prometheus.Desc
}

// NewPgStatActivityCollector instantiates and returns a new PgStatActivityCollector.
func NewPgStatActivityCollector(db *db.Client) *PgStatActivityCollector {
	variableLabels := []string{"datname", "state"}
	return &PgStatActivityCollector{
		db: db,

		activityCount: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, activitySubSystem, "count"),
			"Number of connections in this state",
			variableLabels,
			nil,
		),
		maxTxDuration: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, activitySubSystem, "max_tx_duration"),
			"Max duration in seconds any active transaction has been running",
			variableLabels,
			nil,
		),
	}
}

// Describe implements the prometheus.Collector.
func (c *PgStatActivityCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.activityCount
	ch <- c.maxTxDuration
}

// Collect implements the promtheus.Collector.
func (c *PgStatActivityCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_ = c.Scrape(ch)
}

// Scrape implements our Scraper interfacc.
func (c *PgStatActivityCollector) Scrape(ch chan<- prometheus.Metric) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	activityStats, err := c.db.SelectPgStatActivity(context.Background())
	if err != nil {
		return fmt.Errorf("activity stats: %w", err)
	}

	for _, stat := range activityStats {
		ch <- prometheus.MustNewConstMetric(c.activityCount, prometheus.GaugeValue, float64(stat.Count), stat.DatName, stat.State)
		ch <- prometheus.MustNewConstMetric(c.maxTxDuration, prometheus.GaugeValue, stat.MaxTxDuration, stat.DatName, stat.State)
	}
	return nil
}
