go_library(
    name = "exporter",
    srcs = ["exporter.go"],
    visibility = ["//infrastructure/postgres/exporter/..."],
    deps = [
        "//infrastructure/postgres/exporter/collector",
        "//infrastructure/postgres/exporter/db",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
