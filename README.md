Go language documentation

- [Variables](#variables)
- [Arrays](#arrays)
- [Slices](#slices)
- [Maps](#maps)
- [Errors](#errors)
- [Functions](#functions)
	- [Functions: assigned to objects](#functions-assigned-to-objects)
	- [Functions: for struct](#functions-for-struct)
	- [defer](#defer)
	- [variadic](#variadic)
- [Interfaces](#interfaces)
- [Pointer](#pointer)
- [Concurrency](#concurrency)
	- [race conditions detector](#race-conditions-detector)
	- [Channels](#channels)
	- [select](#select)
	- [sync](#sync)
		- [wait](#wait)
		- [mutex](#mutex)
- [Reflection](#reflection)
- [Http](#http)
	- [context](#context)
	- [Routing](#routing)
- [Property based tests](#property-based-tests)
- [String builder](#string-builder)
- [files](#files)
- [Go tools](#go-tools)
	- [godocs](#godocs)
	- [Example](#example)
	- [Benchmarking](#benchmarking)
	- [Coverage](#coverage)
	- [errcheck](#errcheck)
	- [vet](#vet)
# Variables

```go
x := ""
```
# Arrays
Fixed size, even functions that requires array needs their size.
```go
myArray := [3]int{1,2,3}
```

# Slices

Like array without fix size, it has initial size but can grow with append function. Check slices.go:14
```go
mySlice := []int{1,2,3}
// append:
var sums []int
sums = append(sums, Sum(numbers))
// every element except the first one
mySlice[1:]
```

# Maps
```go
dictionary := map[string]string{"test": "this is just a test"}
dictionary["test"]
```

# Errors
```go
// type
(error)
// create a new error:
errors.New("test")
```

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

## variadic
```go
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
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

`interface{}` means any type

# Pointer
The following code does not change amount in `Wallet`:
```go
type Wallet struct {
	amount int
}
func (r Wallet) Deposit(amount int) {
	r.amount = amount
}
```
The Deposit needs to be changed to pointer, so it works:
```go
func (r *Wallet) Deposit(amount int) ...
```
The reason is:
In Go, when you call a function or a method the arguments are copied.
When calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from.
You can find out what the address of that bit of memory with `&myVal`

Pointers let us point to some values and then let us change them. So rather than taking a copy of the whole Wallet, we instead take a pointer to that wallet so that we can change the original values within it.

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
```go
type result struct {
	string
	bool
}
// Create a channel
resultChannel := make(chan result)
// Send statement
resultChannel <- result{u, wc(u)}
// Receive expression
r := <-resultChannel

// ch is send    only channels
func a(ch chan<- int)
// ch is receive only channels
func a(ch <-chan int)
```

- Sending to a channel will block until it is read by someone.

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency#channels

## select
`myVar := <-ch` means that you wait for values to be sent to a channel, and that is a blocking call.

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

> A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
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

A Mutex must not be copied after first use. So its better to use a pointer and we can create a constructer like below:

```go
func NewCounter() *Counter {
	return &Counter{}
}
```

# Reflection
```go
func t(x interface{}) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	field.String()
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

## context
Can cancel goroutines, etc, example:

```go
func(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := make(chan string, 1)
	go func() {
		data <- Fetch()
	}()

	select {
	case d := <-data:
		fmt.Fprint(w, d)
	case <-ctx.Done():
		Cancel()
	}
}
```
`Done()` signals when it is done or cancelled.
context should be send as the first argument to every http functions.

To cancel the context use the code below:
```go
ctx, cancelCtx := context.WithCancel(parentCtx)
cancelCtx()
```

## Routing
`NewServeMux`is a builtin for routing.

[example](https://quii.gitbook.io/learn-go-with-tests/build-an-application/json#write-enough-code-to-make-it-pass)

# Property based tests

> quick.Check a function that it will run against a number of random inputs, if the function returns false it will be seen as failing the check.

```go
func TestPropertiesOfConversion(t *testing.T) {
	// arabic = 1, roman = I
	// arabic = 2, roman = II
	assertion := func(arabic int) bool {
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	// quick.Check is Property based tests sending random arabic to the test
	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
```

# String builder
A Builder is used to efficiently build a string using Write methods. It minimizes memory copying.
```go
var result strings.Builder
for i := 0; i < 10; i++ {
	result.WriteString("I")
}
return result.String()
```

# files
```go
// to test:
fstest.MapFS{"hello world.md":  {Data: []byte("hi")}}
NewPostsFromFS(fs)
// ...
func NewPostsFromFS(fileSystem fs.FS)
// read dir
dir, err := fs.ReadDir(fileSystem, ".")
// read a file / content
for _, f := range dir {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
// read line by line
scanner := bufio.NewScanner(postFile)

scanner.Scan()
titleLine := scanner.Text()

scanner.Scan()
descriptionLine := scanner.Text()
```

# Go tools

## godocs
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

## Example
```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```
## Benchmarking
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

# output is:
goos: darwin
goarch: amd64
pkg: github.com/quii/learn-go-with-tests/for/v4
10000000           136 ns/op
PASS
# What 136 ns/op means is our function takes on average 136 nanoseconds to run
```

You can reset the time, if there are some preparation for the test with this code:
```go
b.ResetTimer()
```

More info:
https://golang.org/pkg/testing/#hdr-Benchmarks
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration#benchmarking
https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency#write-a-test

## Coverage
```sh
go test -cover
```

## errcheck
```sh
# install
go install github.com/kisielk/errcheck@latest
# run
errcheck .
```

More info: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#unchecked-errors

## vet
```sh
go vet
```
Can find subtle bugs.
