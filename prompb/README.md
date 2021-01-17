
## Can't import prompb
```
$ bazel build //cmd
ERROR: /home/justin/.cache/bazel/_bazel_justin/1693fa99f572d436b2940bbaf4d9bf38/external/com_github_prometheus_prometheus/prompb/BUILD.bazel:34:11: no such target '@com_github_grpc_ecosystem_grpc_gateway//runtime:runtime': target 'runtime' not declared in package 'runtime' defined by /home/justin/.cache/bazel/_bazel_justin/1693fa99f572d436b2940bbaf4d9bf38/external/com_github_grpc_ecosystem_grpc_gateway/runtime/BUILD.bazel and referenced by '@com_github_prometheus_prometheus//prompb:prompb'
ERROR: /home/justin/.cache/bazel/_bazel_justin/1693fa99f572d436b2940bbaf4d9bf38/external/com_github_prometheus_prometheus/prompb/BUILD.bazel:34:11: no such target '@com_github_grpc_ecosystem_grpc_gateway//utilities:utilities': target 'utilities' not declared in package 'utilities' defined by /home/justin/.cache/bazel/_bazel_justin/1693fa99f572d436b2940bbaf4d9bf38/external/com_github_grpc_ecosystem_grpc_gateway/utilities/BUILD.bazel and referenced by '@com_github_prometheus_prometheus//prompb:prompb'
ERROR: Analysis of target '//cmd:cmd' failed; build aborted: Analysis failed
INFO: Elapsed time: 0.077s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (0 packages loaded, 0 targets configured)
```
