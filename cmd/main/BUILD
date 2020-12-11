load("@io_bazel_rules_go//go:def.bzl", "go_binary")

genrule(
    name = "resource-syso",
    outs = ["resource.syso"],
    cmd = "$(location //cmd/sysogen:sysogen) --output $@",
    tools = ["//cmd/sysogen"],
)

go_binary(
    name = "windows",
    srcs = [
        "main.go",
        "resource.syso",
    ],
    goarch = "amd64",
    goos = "windows",
)
