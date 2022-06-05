# bazel-demo

```
go run main.go
bazelisk run //:gazelle
bazelisk run //:gazelle-update-repos
bazelisk run //:bazel-demo
bazelisk query --noimplicit_deps 'deps(//:bazel-demo)'
bazelisk query --noimplicit_deps 'deps(//:bazel-demo)' --output graph > graph.in
dot -Tpng < graph.in > graph.png
```
