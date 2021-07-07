load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["sdk.go"],
    importpath = "go-ws/sdk",
    visibility = ["//visibility:public"],
    deps = [
        "//sdk/types:go_default_library",
        "//sdk/v1:go_default_library",
    ],
)
