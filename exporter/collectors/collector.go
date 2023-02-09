package collectors

import (
	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace           = "pg_stat"
	userTablesSubSystem = "user_tables"
	activitySubSystem   = "activity"
	statementsSubSystem = "statements"
)

// Collector wraps the prometheus.Collector.
type Collector interface {
	prometheus.Collector
	// Scrape is used by our exporter to scrape data from postgres.
	Scrape(ch chan<- prometheus.Metric) error
}

// DefaultCollectors specifies the list of default collectors.
func DefaultCollectors(db *db.Client) []Collector {
	return []Collector{
		NewPgStatActivityCollector(db),
		NewPgStatStatementsCollector(db),
		NewPgStatUserTableCollector(db),
	}
}
