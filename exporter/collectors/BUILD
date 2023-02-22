go_library(
    name = "collectors",
    srcs = [
        "collector.go",
        "pg_locks.go",
        "pg_stat_activity.go",
        "pg_stat_statements.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//exporter/db",
        "//exporter/logging",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
