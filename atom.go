// Package atom provides simple wrappers around types enforcing atomic usage
//
// The wrapper types do not introduce any size overhead and have the same size
// as the wrapped type.
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
	_     noCopy
	value uint32
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
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

// Set sets the new value regardless of the previous value.
func (b *Bool) Set(value bool) {
	if value {
		atomic.StoreUint32(&b.value, 1)
	} else {
		atomic.StoreUint32(&b.value, 0)
	}
}

// Swap atomically sets the new value and returns the previous value.
func (b *Bool) Swap(new bool) (old bool) {
	if new {
		return atomic.SwapUint32(&b.value, 1) > 0
	}
	return atomic.SwapUint32(&b.value, 0) > 0
}

// Value returns the current value.
func (b *Bool) Value() (value bool) {
	return atomic.LoadUint32(&b.value) > 0
}

// Duration is a wrapper for atomically accessed time.Duration values.
type Duration struct {
	_     noCopy
	value int64
}

// Add atomically adds delta to the current value and returns the new value.
// No arithmetic overflow checks are applied.
func (d *Duration) Add(delta time.Duration) (new time.Duration) {
	return time.Duration(atomic.AddInt64(&d.value, int64(delta)))
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (d *Duration) CompareAndSwap(old, new time.Duration) (swapped bool) {
	return atomic.CompareAndSwapInt64(&d.value, int64(old), int64(new))
}

// Set sets the new value regardless of the previous value.
func (d *Duration) Set(value time.Duration) {
	atomic.StoreInt64(&d.value, int64(value))
}

// Sub atomically subtracts delta to the current value and returns the new value.
// No arithmetic underflow checks are applied.
func (d *Duration) Sub(delta time.Duration) (new time.Duration) {
	return time.Duration(atomic.AddInt64(&d.value, -int64(delta)))
}

// Swap atomically sets the new value and returns the previous value.
func (d *Duration) Swap(new time.Duration) (old time.Duration) {
	return time.Duration(atomic.SwapInt64(&d.value, int64(new)))
}

// Value returns the current value.
func (d *Duration) Value() (value time.Duration) {
	return time.Duration(atomic.LoadInt64(&d.value))
}

// errNil is a special error signaling a nil value.
var errNil = errors.New("nil")

// Error is a wrapper for atomically accessed error values
type Error struct {
	_     noCopy
	value atomic.Value
}

// Set sets the new value regardless of the previous value.
// The value may be nil.
func (e *Error) Set(value error) {
	// Setting atomic.Value to nil is not allowed.
	// Use the special error errNil instead to signal a nil value.
	if value == nil {
		value = errNil
	}
	e.value.Store(value)
}

// Value returns the current error value.
func (e *Error) Value() (value error) {
	v := e.value.Load()
	if v == nil || v == errNil {
		return nil
	}
	return v.(error)
}

// Float32 is a wrapper for atomically accessed float32 values.
type Float32 struct {
	_     noCopy
	value uint32
}

// Add adds delta to the current value and returns the new value.
// Note: Internally this performs a CompareAndSwap operation within a loop.
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
// matches the given old value and returns whether the new value was set.
func (f *Float32) CompareAndSwap(old, new float32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&f.value, math.Float32bits(old), math.Float32bits(new))
}

// Set sets the new value regardless of the previous value.
func (f *Float32) Set(value float32) {
	atomic.StoreUint32(&f.value, math.Float32bits(value))
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (f *Float32) Sub(delta float32) (new float32) {
	return f.Add(-delta)
}

// Swap atomically sets the new value and returns the previous value.
func (f *Float32) Swap(new float32) (old float32) {
	return math.Float32frombits(atomic.SwapUint32(&f.value, math.Float32bits(new)))
}

// Value returns the current value.
func (f *Float32) Value() (value float32) {
	return math.Float32frombits(atomic.LoadUint32(&f.value))
}

// Float64 is a wrapper for atomically accessed float64 values.
type Float64 struct {
	_     noCopy
	value uint64
}

// Add adds delta to the current value and returns the new value.
// Note: Internally this performs a CompareAndSwap operation within a loop.
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
// matches the given old value and returns whether the new value was set.
func (f *Float64) CompareAndSwap(old, new float64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&f.value, math.Float64bits(old), math.Float64bits(new))
}

// Set sets the new value regardless of the previous value.
func (f *Float64) Set(value float64) {
	atomic.StoreUint64(&f.value, math.Float64bits(value))
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (f *Float64) Sub(delta float64) (new float64) {
	return f.Add(-delta)
}

// Swap atomically sets the new value and returns the previous value.
func (f *Float64) Swap(new float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&f.value, math.Float64bits(new)))
}

// Value returns the current value.
func (f *Float64) Value() (value float64) {
	return math.Float64frombits(atomic.LoadUint64(&f.value))
}

// Int is a wrapper for atomically accessed int values.
type Int struct {
	_     noCopy
	value uintptr
}

// Add atomically adds delta to the current value and returns the new value.
func (i *Int) Add(delta int) (new int) {
	return int(atomic.AddUintptr(&i.value, uintptr(delta)))
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (i *Int) CompareAndSwap(old, new int) (swapped bool) {
	return atomic.CompareAndSwapUintptr(&i.value, uintptr(old), uintptr(new))
}

// Set sets the new value regardless of the previous value.
func (i *Int) Set(value int) {
	atomic.StoreUintptr(&i.value, uintptr(value))
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (i *Int) Sub(delta int) (new int) {
	return i.Add(-delta)
}

// Swap atomically sets the new value and returns the previous value.
func (i *Int) Swap(new int) (old int) {
	return int(atomic.SwapUintptr(&i.value, uintptr(new)))
}

// Value returns the current value.
func (i *Int) Value() (value int) {
	return int(atomic.LoadUintptr(&i.value))
}

// Int32 is a wrapper for atomically accessed int32 values.
type Int32 struct {
	_     noCopy
	value int32
}

// Add atomically adds delta to the current value and returns the new value.
func (i *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&i.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (i *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (i *Int32) Set(value int32) {
	atomic.StoreInt32(&i.value, value)
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (i *Int32) Sub(delta int32) (new int32) {
	return i.Add(-delta)
}

// Swap atomically sets the new value and returns the previous value.
func (i *Int32) Swap(new int32) (old int32) {
	return atomic.SwapInt32(&i.value, new)
}

// Value returns the current value.
func (i *Int32) Value() (value int32) {
	return atomic.LoadInt32(&i.value)
}

// Int64 is a wrapper for atomically accessed int64 values.
type Int64 struct {
	_     noCopy
	value int64
}

// Add atomically adds delta to the current value and returns the new value.
func (i *Int64) Add(delta int64) (new int64) {
	return atomic.AddInt64(&i.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (i *Int64) CompareAndSwap(old, new int64) (swapped bool) {
	return atomic.CompareAndSwapInt64(&i.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (i *Int64) Set(value int64) {
	atomic.StoreInt64(&i.value, value)
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (i *Int64) Sub(delta int64) (new int64) {
	return i.Add(-delta)
}

// Swap atomically sets the new value and returns the previous value.
func (i *Int64) Swap(new int64) (old int64) {
	return atomic.SwapInt64(&i.value, new)
}

// Value returns the current value.
func (i *Int64) Value() (value int64) {
	return atomic.LoadInt64(&i.value)
}

// String is a wrapper for atomically accessed string values.
// Note: The string value is wrapped in an interface. Thus, this wrapper has
// a memory overhead.
type String struct {
	_     noCopy
	value atomic.Value
}

// Set sets the new value regardless of the previous value.
// Note: Set requires an allocation as the value is wrapped in an interface.
func (s *String) Set(value string) {
	s.value.Store(value)
}

// Value returns the current error value.
func (s *String) Value() (value string) {
	v := s.value.Load()
	if v == nil {
		return ""
	}
	return v.(string)
}

// Uint is a wrapper for atomically accessed uint values.
type Uint struct {
	_     noCopy
	value uintptr
}

// Add atomically adds delta to the current value and returns the new value.
func (u *Uint) Add(delta uint) (new uint) {
	return uint(atomic.AddUintptr(&u.value, uintptr(delta)))
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (u *Uint) CompareAndSwap(old, new uint) (swapped bool) {
	return atomic.CompareAndSwapUintptr(&u.value, uintptr(old), uintptr(new))
}

// Set sets the new value regardless of the previous value.
func (u *Uint) Set(value uint) {
	atomic.StoreUintptr(&u.value, uintptr(value))
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (u *Uint) Sub(delta uint) (new uint) {
	return u.Add(^uint(delta - 1))
}

// Swap atomically sets the new value and returns the previous value.
func (u *Uint) Swap(new uint) (old uint) {
	return uint(atomic.SwapUintptr(&u.value, uintptr(new)))
}

// Value returns the current value.
func (u *Uint) Value() (value uint) {
	return uint(atomic.LoadUintptr(&u.value))
}

// Uint32 is a wrapper for atomically accessed uint32 values.
type Uint32 struct {
	_     noCopy
	value uint32
}

// Add atomically adds delta to the current value and returns the new value.
func (u *Uint32) Add(delta uint32) (new uint32) {
	return atomic.AddUint32(&u.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (u *Uint32) CompareAndSwap(old, new uint32) (swapped bool) {
	return atomic.CompareAndSwapUint32(&u.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (u *Uint32) Set(value uint32) {
	atomic.StoreUint32(&u.value, value)
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (u *Uint32) Sub(delta uint32) (new uint32) {
	return u.Add(^uint32(delta - 1))
}

// Swap atomically sets the new value and returns the previous value.
func (u *Uint32) Swap(new uint32) (old uint32) {
	return atomic.SwapUint32(&u.value, new)
}

// Value returns the current value.
func (u *Uint32) Value() (value uint32) {
	return atomic.LoadUint32(&u.value)
}

// Uint64 is a wrapper for atomically accessed uint64 values.
type Uint64 struct {
	_     noCopy
	value uint64
}

// Add atomically adds delta to the current value and returns the new value.
func (u *Uint64) Add(delta uint64) (new uint64) {
	return atomic.AddUint64(&u.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (u *Uint64) CompareAndSwap(old, new uint64) (swapped bool) {
	return atomic.CompareAndSwapUint64(&u.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (u *Uint64) Set(value uint64) {
	atomic.StoreUint64(&u.value, value)
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (u *Uint64) Sub(delta uint64) (new uint64) {
	return u.Add(^uint64(delta - 1))
}

// Swap atomically sets the new value and returns the previous value.
func (u *Uint64) Swap(new uint64) (old uint64) {
	return atomic.SwapUint64(&u.value, new)
}

// Value returns the current value.
func (u *Uint64) Value() (value uint64) {
	return atomic.LoadUint64(&u.value)
}

// Uintptr is a wrapper for atomically accessed uintptr values.
type Uintptr struct {
	_     noCopy
	value uintptr
}

// Add atomically adds delta to the current value and returns the new value.
func (u *Uintptr) Add(delta uintptr) (new uintptr) {
	return atomic.AddUintptr(&u.value, delta)
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (u *Uintptr) CompareAndSwap(old, new uintptr) (swapped bool) {
	return atomic.CompareAndSwapUintptr(&u.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (u *Uintptr) Set(value uintptr) {
	atomic.StoreUintptr(&u.value, value)
}

// Sub atomically subtracts delta to the current value and returns the new value.
func (u *Uintptr) Sub(delta uintptr) (new uintptr) {
	return u.Add(^uintptr(delta - 1))
}

// Swap atomically sets the new value and returns the previous value.
func (u *Uintptr) Swap(new uintptr) (old uintptr) {
	return atomic.SwapUintptr(&u.value, new)
}

// Value returns the current value.
func (u *Uintptr) Value() (value uintptr) {
	return atomic.LoadUintptr(&u.value)
}

// Value is a wrapper for atomically accessed consistently typed values.
type Value struct {
	_     noCopy
	value atomic.Value
}

// Set sets the new value regardless of the previous value.
// All calls to Set for a given Value must use values of the same concrete type.
// Set of an inconsistent type panics, as does Set(nil).
func (v *Value) Set(value interface{}) {
	v.value.Store(value)
}

// Value returns the current value.
// It returns nil if there has been no call to Set for this Value.
func (v *Value) Value() (value interface{}) {
	return v.value.Load()
}
