# Go `Readdir` vs `ReadDir` Micro Benchmark

## Raison d'Être

Go 1.16 introduced a new API function in the `os` package called `ReadDir`.
`ReadDir` has the same functionality as the previously-existing `Readdir`,
(both return a slice of the contents of a given directory). The authors
claim that `ReadDir` is _much_ faster than `Readdir` as `ReadDir` does not
stat each file in the directory.

In order to test these claims (and see just how cool this new feature is),
this repo contains two Go benchmark tests, one for `ReadDir` and one for `Readdir`.

## How to Run the Benchmarks

Install the Go toolchain, navigate into the repo, and run

```go
go test -bench=.
```

## Results

From a run on my machine:

```console
❯ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/nschmeller/go-readdir-benchmark
cpu: Intel(R) Xeon(R) Platinum 8375C CPU @ 2.90GHz
BenchmarkReaddir-32         1242            948619 ns/op
BenchmarkReadDir-32         3039            380531 ns/op
PASS
ok      github.com/nschmeller/go-readdir-benchmark      2.477s
```

This result run can also be viewed in the GitHub Actions for this repo.
