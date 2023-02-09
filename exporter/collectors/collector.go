package collectors

import (
	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace           = "pg_stat"
	activitySubSystem   = "activity"
	locksSubSystem      = "locks"
	statementsSubSystem = "statements"
	userTablesSubSystem = "user_tables"
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
		NewPgLocksCollector(db),
		NewPgStatStatementsCollector(db),
		NewPgStatUserTableCollector(db),
	}
}
