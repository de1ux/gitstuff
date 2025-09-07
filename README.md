# gitstuff

## Quickstart

```bash
$ bzl run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
$ bzl run //:gitstuff

# Copy the generated binary somewhere
$ bzl build
```

TODO make this work with bazel 8/9