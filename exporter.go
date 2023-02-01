package exporter

import (
	"context"
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"collector"
	"db"
)

const namespace = "pg_stat"

// Exporter collects PostgreSQL metrics and exports them via prometheus.
type Exporter struct {
	db         *db.Client
	collectors []collector.Collector

	// Internal metrics.
	up           prometheus.Gauge
	totalScrapes prometheus.Counter

	mutex sync.RWMutex
}

// MustNew instantiates and returns a new Exporter or panics.
func MustNew(ctx context.Context, dbOpts db.Opts) *Exporter {
	exporter, err := New(ctx, dbOpts)
	if err != nil {
		panic(err)
	}
	return exporter
}

// New instaniates and returns a new Exporter.
func New(ctx context.Context, dbOpts db.Opts) (*Exporter, error) {
	db, err := db.New(ctx, dbOpts)
	if err != nil {
		return nil, fmt.Errorf("creating exporter: %w", err)
	}
	return &Exporter{
		db:         db,
		collectors: collector.DefaultCollectors(db),

		// Internal metrics.
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "up",
			Help:      "Was the last scrape of PostgreSQL successful.",
		}),
		totalScrapes: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "exporter_scrapes_total",
			Help:      "Current total PostgreSQL scrapes",
		}),
	}, nil
}

// WithCustomCollectors lets the exporter scrape custom metrics.
func (e *Exporter) WithCustomCollectors(collectors ...collector.Collector) *Exporter {
	e.collectors = append(e.collectors, collectors...)
	return e
}

// Register the exporter.
func (e *Exporter) Register() { prometheus.MustRegister(e) }

// HealthCheck pings PostgreSQL.
func (e *Exporter) HealthCheck(ctx context.Context) error {
	return e.db.CheckConnection(ctx)
}

// Describe implements the prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// Internal metrics.
	ch <- e.up.Desc()
	ch <- e.totalScrapes.Desc()
	for _, collector := range e.collectors {
		collector.Describe(ch)
	}
}

// Collect implements the promtheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.totalScrapes.Inc()
	up := 1
	group := errgroup.Group{}
	for _, collector := range e.collectors {
		collector := collector
		group.Go(func() error { return collector.Scrape(ch) })
	}
	if err := group.Wait(); err != nil {
		up = 0
	}
	ch <- prometheus.MustNewConstMetric(e.up.Desc(), prometheus.GaugeValue, float64(up))
	ch <- e.totalScrapes
}