Learning GO from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

# Arrays
Fixed size, even functions that requires array needs their size.

# Slices

Like array without fix size, it has initial size but can grow with append function. Check slices.go:14

# Functions

## Functions: assigned to objects
functions can be assigned to objects:

```go
checkSums := func(t testing.TB, got, want []int) {}
// Usage:
checkSums(t, got, want)
```

# godocs
- install
```sh
go install golang.org/x/tools/cmd/godoc@latest
```
- run
```sh
godoc -http=:6060
```

Should show the docs: http://localhost:6060/pkg/

# Example
```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
# Benchmarking
```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
```
Running the benchmark with
```sh
go test -bench=.
```
## More info
https://golang.org/pkg/testing/#hdr-Benchmarks
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration#benchmarking

# Coverage
```sh
go test -cover
```