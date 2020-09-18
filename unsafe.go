// +build !purego,!appengine,!js

package atom

import (
	"sync/atomic"
	"unsafe"
)

// Pointer is a wrapper for atomically accessed unsafe.Pointer values.
type Pointer struct {
	_     noCopy
	value unsafe.Pointer
}

// CompareAndSwap atomically sets the new value only if the current value
// matches the given old value and returns whether the new value was set.
func (p *Pointer) CompareAndSwap(old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(&p.value, old, new)
}

// Set sets the new value regardless of the previous value.
func (p *Pointer) Set(value unsafe.Pointer) {
	atomic.StorePointer(&p.value, value)
}

// Swap atomically sets the new value and returns the previous value.
func (p *Pointer) Swap(new unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(&p.value, new)
}

// Value returns the current value.
func (p *Pointer) Value() (value unsafe.Pointer) {
	return atomic.LoadPointer(&p.value)
}
