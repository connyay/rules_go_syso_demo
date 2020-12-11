# rules_go_syso_demo

Demo showing rules_go + syso files

## Failure

As of rules_go@v0.24.9 this repo does not build.

```console
bazel build //cmd/main:windows
```

Results in:

```console
ERROR: rules_go_syso_demo/cmd/main/BUILD:10:10: in srcs attribute of go_transition_binary rule //cmd/main:windows: '//cmd/main:resource-syso' does not produce any go_transition_binary srcs files (expected .go, .s, .S, .h, .c, .cc, .cpp, .cxx, .h, .hh, .hpp, .hxx, .inc, .m or .mm). Since this rule was created by the macro 'go_binary_macro', the error might have been caused by the macro implementation
ERROR: Analysis of target '//cmd/main:windows' failed; build aborted: Analysis of target '//cmd/main:windows' failed
INFO: Elapsed time: 0.698s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (0 packages loaded, 0 targets configured)
```
