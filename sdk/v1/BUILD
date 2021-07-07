load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dbg_utils.go",
        "sdk_impl.go",
        "sdk_impl_utils.go",
        "set_methods.go",
        "websocket.go",
        "ws_connection.go",
        "ws_handler.go",
    ],
    importpath = "go-ws/sdk/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//sdk/types:go_default_library",
        "@com_github_gorilla_websocket//:go_default_library",
    ],
)
