package collector

import (
	"github.com/prometheus/client_golang/prometheus"

	"infrastructure/postgres/exporter/db"
)

const (
	namespace           = "pg_stat"
	userTablesSubSystem = "user_tables"
	activitySubSystem   = "activity"
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
		NewPgStatUserTableCollector(db),
		NewPgStatActivityCollector(db),
	}
}
