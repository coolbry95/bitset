package bitset

import (
	"testing"
)

func TestingNewBitset(t *testing.T) {
	// this is going to be the minimum size needed to fit size
	
}

func TestSet(t *testing.T) {
	b := NewBitset(300)
	for i:= uint(0); i < 300; i++ {
		b.Set(i)	
	}

	for i := uint(0); i < 300; i++ {
		if b.Get(i) {
			continue
		}
		t.Fatal("bit was not set when it should have been")
	}
}

func TestGet(t *testing.T) {
	b := NewBitset(300)
	for i := uint(0); i < 300; i++ {
		b.Set(i)	
	}

	for i := uint(0); i < 300; i++ {
		if b.Get(i) {
			continue
		}
		t.Fatal("bit was not set when it should have been")
	}
}

func TestClear(t *testing.T) {
	b := NewBitset(300)
	for i := uint(0); i < 300; i++ {
		b.Set(i)	
	}

	for i := uint(0); i < 300; i++ {
		 b.Clear(i)
	}

	for i := uint(0); i < 300; i++ {
		if b.Get(i) {
			t.Fatal("bit was set when it should not be")
		}
	}
}
