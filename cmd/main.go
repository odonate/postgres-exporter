package main

import (
	"context"

	"github.com/odonate/postgres-exporter/exporter"
	"github.com/odonate/postgres-exporter/exporter/db"
)

func main() {
	exporter := exporter.MustNew(context.Background(), exporter.Opts{DBOpts: []db.Opts{{}}})
	exporter.Register()
}
