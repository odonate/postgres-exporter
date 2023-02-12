# PostgreSQL Exporter

Prometheus exporter for PostgreSQL server-side metrics. 

Tested PostgreSQL Versions: `11`

This library is intended as an alternative to the prometheus-community's [postgres_exporter](https://github.com/prometheus-community/postgres_exporter), which still uses [lib/pq](https://github.com/lib/pq); a library that is in maintenance mode.

Instead, we leverage [pgx](https://github.com/jackc/pgx), which aims to be low-level, fast, and performant, while also enabling PostgreSQL-specific features that the standard `database/sql` package does not allow for (and also it is actively maintained).

We also offer support for multi-target scraping (see below).


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
#### Go Binary
```
import (
	"github.com/odonate/postgres-exporter/exporter"
	"github.com/odonate/postgres-exporter/exporter/db"
)

var opts struct {
  DB db.Opts `group:"Postgres"`
}

func main() {
  flags.MustParse(&opts)
  exporterOpts := exporter.Opts{DBOpts: []db.Opts{opts.DB}}
  exporter := exporter.MustNew(context.Background(), exporterOpts)
  exporter.Register()
  ...
```
#### Kubernetes Deployment
```
---
apiVersion: apps/v1
kind: Deployment
metadata: 
  ...
spec:
  ...
    spec:
      containers:
      - name: main
        image: <Go-Binary-Docker-Image>
        args:
          - --application_name=postgres-exporter
        resources:
          requests:
            memory: 20Mi
            cpu: 5m
          limits:
            memory: 10Mi
            cpu: 10m
        ports:
        - containerPort: 13434
          name: prometheus

        envFrom:
          - secretRef:
              name: your-db-secret

```

### Multi-Target
#### Go Binary
```
import (
	"github.com/odonate/postgres-exporter/exporter"
	"github.com/odonate/postgres-exporter/exporter/db"
)

var opts struct {
  FirstDB db.Opts{} `namespace:"first" env-namespace:"FIRST"`
  SecondDB db.Opts{} `namespace:"second" env-namespace:"SECOND"`
}

func main() {
  flags.MustParse(&opts)
  exporterOpts := exporter.Opts{
    DBOpts: []db.Opts{
      opts.FirstDB, 
      opts.SecondDB,
    },
  }
  exporter := exporter.MustNew(context.Background(), exporterOpts)
  exporter.Register()
  ...
```
#### Kubernetes Deployment
```
---
apiVersion: apps/v1
kind: Deployment
metadata: 
  ...
spec:
  ...
    spec:
      containers:
      - name: main
        image: <Go-Binary-Docker-Image>
        args:
          - --first.application_name=postgres-exporter
          - --second.application_name=postgres-exporter
        resources:
          requests:
            memory: 20Mi
            cpu: 5m
          limits:
            memory: 10Mi
            cpu: 10m
        ports:
        - containerPort: 13434
          name: prometheus
        envFrom:
          - secretRef:
              name: your-first-db-secret
            prefix: FIRST_
          - secretRef:
              name: your-second-db-secret
            prefix: SECOND_

```

