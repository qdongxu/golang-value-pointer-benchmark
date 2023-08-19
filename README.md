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
BenchmarkCalcWithValue-10      	  156490	      6839 ns/op
BenchmarkCalcWithPointer-10    	  139938	      8560 ns/op
BenchmarkChanWithValue-10      	     168	   7033463 ns/op
BenchmarkChanWithPointer-10    	     130	   9103360 ns/op
BenchmarkMethodValue-10        	10595004	       113.3 ns/op
BenchmarkMethodPointer-10      	12543979	        94.60 ns/op
PASS
ok  	github.com/qdongxu/golang-value-pointer-benchmark/cmd	9.533s
```

A short summary, use value semantics prioritized and let golang to optimize. Only use pointer semantics when a pointer is required(e.g. the data should be changed).
