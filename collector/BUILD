go_library(
    name = "collector",
    srcs = [
        "collector.go",
        "pg_stat_activity.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:postgres-exporter-db",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
