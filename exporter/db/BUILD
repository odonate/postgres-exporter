go_library(
    name = "db",
    srcs = [
        "db.go",
        "dsn.go",
        "opts.go",
        "pg_lock.go",
        "pg_stat_activity.go",
        "pg_stat_statements.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//exporter/db/model",
        "//third_party/go:pgx.v4",
        "//third_party/go:scany",
    ],
)
