// +build !purego,!appengine,!js

package atom

import (
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var p Pointer
	if p.Value() != nil {
		t.Fatal("Expected initial value to be nil")
	}

	var t1, t2 uint64
	v1 := unsafe.Pointer(&t1)
	p.Set(v1)
	if v := p.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	v2 := unsafe.Pointer(&t2)
	if p.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := p.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !p.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := p.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if p.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := p.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}
