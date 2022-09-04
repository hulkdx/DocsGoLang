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

## Functions: for struct
```go
type Rectangle struct {
	Width  float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}
```

# Interfaces

No need to inherit / implements, it just works.

```go
type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
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

In macos the path directory is `~/go/bin/godoc`

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

# errcheck
```sh
# install
go install github.com/kisielk/errcheck@latest
# run
errcheck .
```

More info: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#unchecked-errors
