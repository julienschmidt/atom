package atom

import (
	"errors"
	"testing"
	"time"
)

func TestBool(t *testing.T) {
	// make go cover happy
	var nc noCopy
	nc.Lock()

	var b Bool
	if b.Value() {
		t.Fatal("Expected initial value to be false")
	}

	b.Set(true)
	if v := b.Value(); !v {
		t.Fatal("Value is still false")
	}
	b.Set(false)
	if v := b.Value(); v {
		t.Fatal("Value unchanged")
	}

	if b.CompareAndSwap(true, false) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := b.Value(); v {
		t.Fatal("Value changed")
	}

	if !b.CompareAndSwap(false, true) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := b.Value(); !v {
		t.Fatal("Value unchanged")
	}

	if !b.Swap(true) {
		t.Fatal("Old value does not match")
	}
	if v := b.Value(); !v {
		t.Fatal("Value unchanged")
	}
	if !b.Swap(false) {
		t.Fatal("Old value does not match")
	}
	if v := b.Value(); v {
		t.Fatal("Value unchanged")
	}
}

func TestDuration(t *testing.T) {
	var d Duration
	if d.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	v1 := time.Duration(1337)
	d.Set(v1)
	if v := d.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if d.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if d.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	v2 := time.Duration(987654321)
	if d.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := d.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !d.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := d.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if d.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := d.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestError(t *testing.T) {
	var e Error
	if e.Value() != nil {
		t.Fatal("Expected initial value to be nil")
	}

	a := errors.New("a")

	e.Set(a)
	if v := e.Value(); v != a {
		if v == nil {
			t.Fatal("Value is still nil")
		}
		t.Fatal("Value did not match")
	}
	e.Set(nil)
	if v := e.Value(); v == a {
		t.Fatal("Value still matches initial value")
	} else if v != nil {
		t.Fatal("Value did not match")
	}
}

func TestFloat32(t *testing.T) {
	var f Float32
	if f.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	var v1 float32 = 13.37
	f.Set(v1)
	if v := f.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if f.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if f.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	var v2 float32 = 98765.4321
	if f.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := f.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !f.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := f.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if f.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := f.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestFloat64(t *testing.T) {
	var f Float64
	if f.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	v1 := 13.37
	f.Set(v1)
	if v := f.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if f.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if f.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	v2 := 98765.4321
	if f.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := f.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !f.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := f.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if f.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := f.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestInt32(t *testing.T) {
	var i Int32
	if i.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	var v1 int32 = 1337
	i.Set(v1)
	if v := i.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if i.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if i.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	var v2 int32 = 987654321
	if i.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := i.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !i.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := i.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if i.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := i.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestInt64(t *testing.T) {
	var i Int64
	if i.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	var v1 int64 = 1337
	i.Set(v1)
	if v := i.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if i.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if i.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	var v2 int64 = 987654321
	if i.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := i.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !i.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := i.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if i.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := i.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestString(t *testing.T) {
	var s String
	if s.Value() != "" {
		t.Fatal("Expected initial value to be an empty string")
	}

	s.Set("a")
	if v := s.Value(); v != "a" {
		if v == "" {
			t.Fatal("Value is still an empty string")
		}
		t.Fatal("Value did not match")
	}
	s.Set("")
	if v := s.Value(); v == "a" {
		t.Fatal("Value still matches initial value")
	} else if v != "" {
		t.Fatal("Value did not match")
	}
}

func TestUint32(t *testing.T) {
	var u Uint32
	if u.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	var v1 uint32 = 1337
	u.Set(v1)
	if v := u.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if u.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if u.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	var v2 uint32 = 987654321
	if u.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := u.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !u.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := u.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if u.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := u.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}

func TestUint64(t *testing.T) {
	var u Uint64
	if u.Value() != 0 {
		t.Fatal("Expected initial value to be 0")
	}

	var v1 uint64 = 1337
	u.Set(v1)
	if v := u.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}

	if u.Add(-v1) != 0 {
		t.Fatal("New value does not match")
	}
	if u.Add(v1) != v1 {
		t.Fatal("New value does not match")
	}

	var v2 uint64 = 987654321
	if u.CompareAndSwap(v2, v2) {
		t.Fatal("CompareAndSwap reported swap when the old value did not match")
	}
	if v := u.Value(); v != v1 {
		t.Fatal("Value changed")
	}

	if !u.CompareAndSwap(v1, v2) {
		t.Fatal("CompareAndSwap did not report a swap")
	}
	if v := u.Value(); v != v2 {
		t.Fatal("Value unchanged")
	}

	if u.Swap(v1) != v2 {
		t.Fatal("Old value does not match")
	}
	if v := u.Value(); v != v1 {
		t.Fatal("Value unchanged")
	}
}
