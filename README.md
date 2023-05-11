# Go `Readdir` vs `ReadDir` Micro Benchmark

## Repository Raison d'Être

Go 1.16 [introduced](https://go.dev/doc/go1.16#:~:text=The%20package%20defines,File.ReadDir.)
a new API function in the `os` package called [`ReadDir`](https://pkg.go.dev/os#File.ReadDir).
`ReadDir` has the same functionality as the previously-existing
[`Readdir`](https://pkg.go.dev/os#File.Readdir), (both return a slice of
the contents of a given directory). The intent is that `ReadDir` is _much_
faster than `Readdir`, as `ReadDir` does not stat each file in the directory.

In order to test these claims (and measure the coolness of this new feature),
this repo contains two Go benchmark tests: one for `ReadDir` and one for `Readdir`.

The benchmarked loop for each test consists of:

* Call to `os.Open("/proc")` to get a file handle to read ( + error check)
  * I chose to read `/proc` because it was part of the inspiration for this repo
* Call to the "readdir" function under test ( + error check)

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

`ReadDir` is 2.5x faster than `Readdir` here!

This result run can also be viewed in the GitHub Actions for this repo.
From an example run:

```console
Run go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/nschmeller/go-readdir-benchmark
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkReaddir-2   	    8221	    148709 ns/op
BenchmarkReadDir-2   	   46755	     23487 ns/op
PASS
ok  	github.com/nschmeller/go-readdir-benchmark	3.179s
```

In this case, `ReadDir` is 6.3x faster than `Readdir`! 🔥
