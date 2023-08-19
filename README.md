# golang-value-pointer-benchmark
A lax benchmark about the value and pointer semantics.

From the experience of c/c++, I always think that the pointer semantics is faster than the value semantics. But in golang, it is not always true.

As the benchmark shows, the value semantics is even faster than the pointer semantics. It's probable 'values' are allocated on stack, and 'pointers' are allocated on heap.

In the case of struct method, the value semantics is a little slower than the pointer semantics, but only ~1 nanosecond per invocation.

```shell
cmd> go test -bench=.
goos: darwin
goarch: arm64
pkg: github.com/qdongxu/golang-value-pointer-benchmark/cmd
BenchmarkCalcWithValue-10       16498020                66.62 ns/op
BenchmarkCalcWithPointer-10     16337106                73.82 ns/op
BenchmarkChanWithValue-10           9410            130618 ns/op
BenchmarkChanWithPointer-10         7614            152127 ns/op
BenchmarkMethodValue-10         10743934               111.7 ns/op
BenchmarkMethodPointer-10       12711835                94.31 ns/op
PASS
ok      github.com/qdongxu/golang-value-pointer-benchmark/cmd   7.903s
```

A short summary, use value semantics prioritized and let golang to optimize. Only use pointer semantics when a pointer is required(e.g. the data should be changed).
