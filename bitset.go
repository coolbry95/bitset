package bitset

// things needed
// set
// get
// new
//import "fmt"

// referenced again
// https://github.com/Workiva/go-datastructures/blob/master/hashmap/fastinteger/hashmap.go#L47

// does i / 64 make an assumption of the amount of buckets in the set?
// https://github.com/chriso/bitset/blob/master/lib/bitset.c#L72

// max theoretical size is int64
// so if the machine is only 32 bit it is going to be int32 because it cannot
// reach int64 in size of memory or [Maxint64]slice is > than 4Gb memory
// if we leave the size as uint then we do not have to worry about people
// trying to exceed the maximum size because int will default to 32 or 64 on
// its own

type Bitset struct {
	data   []uint64
	// max array size * 64
	length uint64
}

func NewBitset(size uint) *Bitset {
	// this is going to be the minimum size needed to fit size
	// power of two size may increase speed
	return &Bitset{
		make([]uint64, (size+63)/64),
		uint64(size),
	}
}

// maybe return if was set to true/false
// if it was turned to 0 return false

// what does it need to take though
// does it need to be uint64 to reach every position
func (set *Bitset) Set(bit uint64) {
	if bit > set.length {
		set.miracleGrow(bit)
	}
	set.data[bit/64] |= 1 << (bit % 64)
}

/*
func (set *Bitset) growMaybe(bit uint64) {
	if bit > set.length {
		set.miracleGrow(bit)
	}
}
*/

func (set *Bitset) Get(bit uint64) bool {
	if bit > set.length {
		return false
	}
	return set.data[bit/64] & 1 << (bit % 64) != 0
}

func (set *Bitset) Clear(bit uint64) {
	if bit > set.length {
		set.miracleGrow(bit)
	}
	// this is picking the bucket based on there being 64 of them
	// change this to length of array
	// see reference above
	set.data[bit/64] &^= 1 << (bit % 64)
}

// diagnose what is wrong with this. a hash is to big for it.
func (set *Bitset) miracleGrow(bit uint64) {
	// grow double the size? reduce the chance of growing in the near future

	// find next size to reach size
	// what we want + (64-1) / 64
	// 63 becuse of zero count/ so we can reach size and / 64 for buckets
	// without +63 we are just below the right number
	// with we are just above
	a := (bit + 63) / 64
	temp := make([]uint64, a)
	copy(temp, set.data)
	set.data = temp
	set.length = uint64(a)
}
