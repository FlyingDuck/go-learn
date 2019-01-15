package effective

import (
	"fmt"
)

type Sequence []int

// Methods required by sort.Interface.

// Len is the number of elements in the collection.
func (s Sequence) Len() int  {
	return len(s)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s Sequence)Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap swaps the elements with indexes i and j.
func (s Sequence)Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.

func (s Sequence) String() string  {
	//sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

// conversions

func WhatIsType(arg interface{}) string {
	switch str := arg.(type) {
	case string:
		return str + "<string>"
	case int:
		return fmt.Sprint(str,  "<int>")
	case Sequence:
		return fmt.Sprint(str, "<Sequence>")
	default:
		return "Unknown"
		
	}
}









