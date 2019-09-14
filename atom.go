// Package atom provides simple wrappers around types enforcing atomic usage
package atom

import (
	"errors"
	"math"
	"sync/atomic"
	"time"
)

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock() {}

// Bool is a wrapper around uint32 for usage as a boolean value with
// atomic access.
type Bool struct {
	_noCopy noCopy
	value   uint32
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (b *Bool) CompareAndSwap(old, new bool) (swapped bool) {
	var uold, unew uint32
	if old {
		uold = 1
	}
	if new {
		unew = 1
	}
	return atomic.CompareAndSwapUint32(&b.value, uold, unew)
}

// Set sets the new value regardless of the previous value
func (b *Bool) Set(value bool) {
	if value {
		atomic.StoreUint32(&b.value, 1)
	} else {
		atomic.StoreUint32(&b.value, 0)
	}
}

// Swap atomically sets the new value and returns the previous value
func (b *Bool) Swap(new bool) (old bool) {
	if new {
		return atomic.SwapUint32(&b.value, 1) > 0
	}
	return atomic.SwapUint32(&b.value, 0) > 0
}

// Value returns the current value
func (b *Bool) Value() (value bool) {
	return atomic.LoadUint32(&b.value) > 0
}

// Duration is a wrapper for atomically accessed time.Duration values
type Duration struct {
	_noCopy noCopy
	value   int64
}

// Add atomically adds delta to the current value and returns the new value
func (d *Duration) Add(delta time.Duration) (new time.Duration) {
	return time.Duration(atomic.AddInt64(&d.value, int64(delta)))
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (d *Duration) CompareAndSwap(old, new time.Duration) (swapped bool) {
	return atomic.CompareAndSwapInt64(&d.value, int64(old), int64(new))
}

// Set sets the new value regardless of the previous value
func (d *Duration) Set(value time.Duration) {
	atomic.StoreInt64(&d.value, int64(value))
}

// Swap atomically sets the new value and returns the previous value
func (d *Duration) Swap(new time.Duration) (old time.Duration) {
	return time.Duration(atomic.SwapInt64(&d.value, int64(new)))
}

// Value returns the current value
func (d *Duration) Value() (value time.Duration) {
	return time.Duration(atomic.LoadInt64(&d.value))
}

var errNil = errors.New("nil")

// Error is a wrapper for atomically accessed error values
type Error struct {
	_noCopy noCopy
	value   atomic.Value
}

// Set sets the new value regardless of the previous value.
// The value may be nil
func (e *Error) Set(value error) {
	if value == nil {
		value = errNil
	}
	e.value.Store(value)
}

// Value returns the current error value
func (e *Error) Value() (value error) {
	v := e.value.Load()
	if v == nil || v == errNil {
		return nil
	}
	return v.(error)
}

// Float32 is a wrapper for atomically accessed float32 values
type Float32 struct {
	_noCopy noCopy
	value   uint32
}

// Add adds delta to the current value and returns the new value.
// Note: Internally this performs a CompareAndSwap operation within a loop
func (f *Float32) Add(delta float32) (new float32) {
	for {
		old := f.Value()
		new := old + delta
		if f.CompareAndSwap(old, new) {
			return new
		}
	}
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (f *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&f.value, math.Float32bits(old), math.Float32bits(new))
}

// Set sets the new value regardless of the previous value
func (f *Float32) Set(value float32) {
	atomic.StoreUint32(&f.value, math.Float32bits(value))
}

// Swap atomically sets the new value and returns the previous value
func (f *Float32) Swap(new float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(&f.value, math.Float32bits(new)))
}

// Value returns the current value
func (f *Float32) Value() (value float32) {
	return math.Float32frombits(atomic.LoadUint32(&f.value))
}

// Float64 is a wrapper for atomically accessed float64 values
type Float64 struct {
	_noCopy noCopy
	value   uint64
}

// Add adds delta to the current value and returns the new value.
// Note: Internally this performs a CompareAndSwap operation within a loop
func (f *Float64) Add(delta float64) (new float64) {
	for {
		old := f.Value()
		new := old + delta
		if f.CompareAndSwap(old, new) {
			return new
		}
	}
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (f *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&f.value, math.Float64bits(old), math.Float64bits(new))
}

// Set sets the new value regardless of the previous value
func (f *Float64) Set(value float64) {
	atomic.StoreUint64(&f.value, math.Float64bits(value))
}

// Swap atomically sets the new value and returns the previous value
func (f *Float64) Swap(new float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&f.value, math.Float64bits(new)))
}

// Value returns the current value
func (f *Float64) Value() (value float64) {
	return math.Float64frombits(atomic.LoadUint64(&f.value))
}

// Int32 is a wrapper for atomically accessed int32 values
type Int32 struct {
	_noCopy noCopy
	value   int32
}

// Add atomically adds delta to the current value and returns the new value
func (i *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&i.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (i *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}

// Set sets the new value regardless of the previous value
func (i *Int32) Set(value int32) {
	atomic.StoreInt32(&i.value, value)
}

// Swap atomically sets the new value and returns the previous value
func (i *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&i.value, new)
}

// Value returns the current value
func (i *Int32) Value() (value int32) {
	return atomic.LoadInt32(&i.value)
}

// Int64 is a wrapper for atomically accessed int64 values
type Int64 struct {
	_noCopy noCopy
	value   int64
}

// Add atomically adds delta to the current value and returns the new value
func (i *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&i.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (i *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}

// Set sets the new value regardless of the previous value
func (i *Int64) Set(value int64) {
	atomic.StoreInt64(&i.value, value)
}

// Swap atomically sets the new value and returns the previous value
func (i *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(&i.value, new)
}

// Value returns the current value
func (i *Int64) Value() (value int64) {
	return atomic.LoadInt64(&i.value)
}

// String is a wrapper for atomically accessed string values
type String struct {
	_noCopy noCopy
	value   atomic.Value
}

// Set sets the new value regardless of the previous value.
// The value may be nil
func (s *String) Set(value string) {
	s.value.Store(value)
}

// Value returns the current error value
func (s *String) Value() (value string) {
	v := s.value.Load()
	if v == nil {
		return ""
	}
	return v.(string)
}

// Uint32 is a wrapper for atomically accessed uint32 values
type Uint32 struct {
	_noCopy noCopy
	value   uint32
}

// Add atomically adds delta to the current value and returns the new value
func (u *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&u.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (u *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&u.value, old, new)
}

// Set sets the new value regardless of the previous value
func (u *Uint32) Set(value uint32) {
	atomic.StoreUint32(&u.value, value)
}

// Swap atomically sets the new value and returns the previous value
func (u *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&u.value, new)
}

// Value returns the current value
func (u *Uint32) Value() (value uint32) {
	return atomic.LoadUint32(&u.value)
}

// Uint64 is a wrapper for atomically accessed uint64 values
type Uint64 struct {
	_noCopy noCopy
	value   uint64
}

// Add atomically adds delta to the current value and returns the new value
func (u *Uint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&u.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set
func (u *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&u.value, old, new)
}

// Set sets the new value regardless of the previous value
func (u *Uint64) Set(value uint64) {
	atomic.StoreUint64(&u.value, value)
}

// Swap atomically sets the new value and returns the previous value
func (u *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(&u.value, new)
}

// Value returns the current value
func (u *Uint64) Value() (value uint64) {
	return atomic.LoadUint64(&u.value)
}
