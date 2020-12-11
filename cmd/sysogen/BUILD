load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "sysogen",
    srcs = ["main.go"],
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_josephspurrier_goversioninfo//:goversioninfo",
    ],
)
