# PostgreSQL Exporter

Prometheus exporter Golang library for PostgreSQL server-side metrics. An alternative to [pometheus-community/postgres_exporter](https://github.com/prometheus-community/postgres_exporter), which still uses [lib/pq](https://github.com/lib/pq).

As `lib/pq` is in maintenance mode, this library leverages [pgx](https://github.com/jackc/pgx), which aims to be low-level, fast, and performant, while also enabling PostgreSQL-specific features that the standard `database/sql` package does not allow for (and is actively maintained).


## Options

| Long Flag                 | ENV Flag                 | Default           | Description                                          |
|---------------------------|--------------------------|-------------------|------------------------------------------------------|
| --postgres_host           | $POSTGRES_HOST           | localhost         | PostgreSQL Host                                      |
| --postgres_port           | $POSTGRES_PORT           | 5432              | PostgreSQL Port                                      |
| --postgres_user           | $POSTGRES_USER           | postgres          | PostgreSQL User                                      |
| --postgres_database       | $POSTGRES_DATABASE       | postgres          | PostgreSQL Database                                  |
| --postgres_password       | $POSTGRES_PASSWORD       | postgres          | PostgreSQL Password                                  |
| --auth_mechanism          | $AUTH_MECHANISM          | password          | The mechanism to use when authenticating with the DB |
| --application_name        | $APP_NAME                | postgres-exporter | The name of the application.                         |
| --default_isolation_level | $DEFAULT_ISOLATION_LEVEL | REPEATABLE_READ   | The default isolation level for DB transactions      |

## Example Usage

### Single Target
```
var opts struct {
  Exporter exporter.Opts{}
}

func main() {
  flags.MustParse(&opts)
  exporter := exporter.MustNew(context.Background(), opts.Exporter)
  exporter.Register()
  ...
```

### Multi-Target
```
var opts struct {
  ExporterA exporter.Opts{} `namespace:"a" env-namespace:"A"`
  ExporterB exporter.Opts{} `namespace:"b" env-namespace:"B"`
}

func main() {
  flags.MustParse(&opts)
  exporterA := exporter.MustNew(context.Background(), opts.ExporterA)
  exporterA.Register()
  exporterB := exporter.MustNew(context.Background(), opts.ExporterB)
  exporterB.Register()
  ...
```
