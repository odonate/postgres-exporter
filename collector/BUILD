go_library(
    name = "collector",
    srcs = [
        "collector.go",
        "pg_stat_activity.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["//infrastructure/postgres/exporter/..."],
    deps = [
        "//infrastructure/postgres/exporter/db",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
