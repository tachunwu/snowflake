# Snowflake ID
## How to use
```go
workerID := uint16(1)
sf := snowflake.NewSnowflake(workerID)
id := sf.NextID()
int64(id)

```
## Test
```shell
ok  	github.com/tachunwu/snowflake	(cached)
```
## Benchmark
```shell
goos: darwin
goarch: amd64
pkg: github.com/tachunwu/snowflake
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkSnowflake-4   	 4927920	       244.5 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/tachunwu/snowflake	1.836s
```
