load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES,
    visibility = ["//visibility:public"],
)

# gazelle:prefix go-ws/
gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)

load("@io_bazel_rules_k8s//k8s:objects.bzl", "k8s_objects")

filegroup(
    name = "build",
    srcs = [
        "//:main",
    ],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
    ],
    importpath = "go-ws/",
    deps = [
        "//sdk:go_default_library",
        "//sdk/types:go_default_library",
    ],
)

go_binary(
    name = "main",
    embed = [":go_default_library"],
)
