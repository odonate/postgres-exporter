go_library(
    name = "db",
    srcs = [
        "db.go",
        "dsn.go",
        "model.go",
        "opts.go",
        "pg_lock.go",
        "pg_stat_activity.go",
        "pg_stat_user_table.go",
    ],
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go:pgtype",
        "//third_party/go:pgx.v4",
        "//third_party/go:scany",
    ],
)
