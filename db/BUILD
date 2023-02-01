go_library(
    name = "db",
    srcs = [
        "db.go",
        "dsn.go",
        "opts.go",
        "pg_lock.go",
        "pg_stat_activity.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["//infrastructure/postgres/exporter/..."],
    deps = [
        "//infrastructure/postgres/exporter/db/model",
        "//third_party/go:pgx.v4",
        "//third_party/go:scany",
    ],
)
