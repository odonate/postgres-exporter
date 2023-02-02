go_library(
    name = "exporter",
    srcs = [
        "exporter.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//exporter/db",
        "//exporter/collectors",
        "//third_party/go:prometheus-client",
        "//third_party/go:x_sync",
    ],
)
