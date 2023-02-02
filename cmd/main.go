package main

import (
	"context"

	"github.com/odonate/postgres-exporter/exporter"
)

func main() {
	exporter := exporter.MustNew(context.Background(), exporter.Opts{})
	exporter.Register()
}
