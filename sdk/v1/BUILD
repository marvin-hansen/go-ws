load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "asset_symbols.go",
        "dbg_utils.go",
        "sdk_impl.go",
        "sdk_impl_utils.go",
        "set_methods.go",
        "ws_handler.go",
        "ws_order_handler.go",
        "ws_order_req.go",
    ],
    importpath = "go-ws/sdk/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//sdk/types:go_default_library",
        "//sdk/web_socket:go_default_library",
        "@com_github_bitly_go_simplejson//:go_default_library",
    ],
)
