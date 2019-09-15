# atom
[![Build Status](https://travis-ci.com/julienschmidt/atom.svg?branch=master)](https://travis-ci.com/julienschmidt/atom) [![Coverage Status](https://coveralls.io/repos/github/julienschmidt/atom/badge.svg?branch=master)](https://coveralls.io/github/julienschmidt/atom?branch=master) [![GoDoc](https://godoc.org/github.com/julienschmidt/atom?status.svg)](https://godoc.org/github.com/julienschmidt/atom)

Intuitive wrapper types enforcing atomic access for lock-free concurrency.

A safe and convenient alternative to sync/atomic.

- Prevents unsafe non-atomic access
- Prevents unsafe copying (which is a non-atomic read)
- No size overhead. The wrappers have the same size as the wrapped type

## Usage

A simple counter can be implemented as follows:

```go
var counter atom.Uint64

// concurrently:

// optionally set a specific value
counter.Set(42)

// increase counter
counter.Add(1)

fmt.Println("Counter value:", counter.Value())
```

Instead of using a costly [`sync.Mutex`](https://golang.org/pkg/sync/#Mutex) to guard the counter:

```go
var (
    counter   uint64
    counterMu sync.Mutex
)

// concurrently:

// optionally set a specific value
counterMu.Lock()
counter = 42
counterMu.Unlock()

// increase counter
counterMu.Lock()
counter++
counterMu.Unlock()

counterMu.Lock()
value := counter
counterMu.Unlock()
fmt.Println("Counter value:", value)
```

Or using [`sync/atomic`](https://golang.org/pkg/sync/atomic/) directly, using an unintuitive interface and not guarding against non-atomic access to the counter:

```go
var counter uint64

// concurrently:

// optionally set a specific value
atomic.StoreUint64(&counter, 42)

// increase counter
atomic.AddUint64(&counter, 1)

fmt.Println("Counter value:", atomic.LoadUint64(&counter))

// beware of direct access!
counter = 0
```
