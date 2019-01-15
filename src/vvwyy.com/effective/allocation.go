package effective

// new
// Let's talk about new first. It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize the memory, it only zeros it.
// That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of type *T.
// In Go terminology, it returns a pointer to a newly allocated zero value of type T.



// make
// Back to allocation. The built-in function make(T, args) serves a purpose different from new(T). It creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).


func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) { // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data)*2))
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	return slice
}

