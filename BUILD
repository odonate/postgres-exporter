go_library(
    name = "exporter",
    srcs = [
        "collector.go",
        "exporter.go",
        "pg_stat_activity.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//db:db",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
