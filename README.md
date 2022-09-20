Learning GO from https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

# Arrays
Fixed size, even functions that requires array needs their size.

# Slices

Like array without fix size, it has initial size but can grow with append function. Check slices.go:14

# Functions

private functions start with lower-case and public functions start with capital.

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

## defer
By prefixing a function call with defer it will now call that function at the end of the containing function. Example: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select#refactor


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
# Concurrency

`go doSomething()` will run doSomething concurrently (similar to `suspend fun` in kotlin). The example below will run in anonymous function:
```go
go func() {
	results[url] = wc(url)
}()
```

## race conditions detector
```sh
go test -race
```

More info:
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency

## Channels

Writing to the map can cause race condition, we can use channels in go:

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency#channels

## select
`myVar := <-ch` means that you wait for values to be sent to a channel, and that is blocking call.

`select` wait on multiple channels, the first one wins and the code underneath the `case` is executed.

```go
// 1. struct{} is the smallest datatype avaialble. But it can be any other type like bool
// 2. always use make for channels
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}
```

## sync

### wait

```go
var wg sync.WaitGroup
wg.Add(wantedCount)

for i := 0; i < wantedCount; i++ {
	go func() {
		counter.Inc()
		wg.Done()
	}()
}
wg.Wait()
```

### mutex

```go
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
```

# Http
```go
http.Get("http://www.facebook.com")

// in test:
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}))
url := server.URL
server.Close()
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

You can reset the time, if there are some preparation for the test with this code:
```go
b.ResetTimer()
```

## More info
https://golang.org/pkg/testing/#hdr-Benchmarks
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration#benchmarking
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency#write-a-test

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

# go vet
Remember to use go vet in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.

# Others

- The `var` keyword allows us to define values global to the package.

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor-3
