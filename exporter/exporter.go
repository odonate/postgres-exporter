package exporter

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/odonate/postgres-exporter/exporter/collectors"
	"github.com/odonate/postgres-exporter/exporter/db"
	"github.com/odonate/postgres-exporter/exporter/logging"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"
)

var log = logging.NewLogger()

const namespace = "pg_stat"

// Opts for the exporter.
type Opts struct {
	DBOpts []db.Opts
}

// Exporter collects PostgreSQL metrics and exports them via prometheus.
type Exporter struct {
	dbClients  []*db.Client
	collectors []collectors.Collector

	// Internal metrics.
	up           prometheus.Gauge
	totalScrapes prometheus.Counter

	mutex sync.RWMutex
}

// MustNew instantiates and returns a new Exporter or panics.
func MustNew(ctx context.Context, opts Opts) *Exporter {
	exporter, err := New(ctx, opts)
	if err != nil {
		panic(err)
	}
	return exporter
}

// New instaniates and returns a new Exporter.
func New(ctx context.Context, opts Opts) (*Exporter, error) {
	if len(opts.DBOpts) < 1 {
		return nil, fmt.Errorf("missing db opts")
	}
	dbClients := make([]*db.Client, 0, len(opts.DBOpts))
	for _, dbOpt := range opts.DBOpts {
		dbClient, err := db.New(ctx, dbOpt)
		if err != nil {
			return nil, fmt.Errorf("creating exporter: %w", err)
		}
		dbClients = append(dbClients, dbClient)
	}
	return &Exporter{
		dbClients:  dbClients,
		collectors: collectors.DefaultCollectors(dbClients),

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
func (e *Exporter) WithCustomCollectors(collectors ...collectors.Collector) *Exporter {
	e.collectors = append(e.collectors, collectors...)
	return e
}

// Register the exporter.
func (e *Exporter) Register() { prometheus.MustRegister(e) }

// HealthCheck pings PostgreSQL.
func (e *Exporter) HealthCheck(ctx context.Context) error {
	group := errgroup.Group{}
	for _, dbClient := range e.dbClients {
		ctx := ctx
		dbClient := dbClient
		group.Go(func() error { return dbClient.CheckConnection(ctx) })
	}
	return group.Wait()
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
	start := time.Now()
	defer log.Infof("exporter collect took %dms", time.Now().Sub(start).Milliseconds())
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
		fmt.Println(fmt.Sprintf("collecting: %w", err))
	}
	ch <- prometheus.MustNewConstMetric(e.up.Desc(), prometheus.GaugeValue, float64(up))
	ch <- e.totalScrapes
}
