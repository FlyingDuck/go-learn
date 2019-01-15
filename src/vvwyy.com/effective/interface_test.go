package effective

import (
	"testing"
	"sort"
)

func TestSort(t *testing.T)  {
	cases := []struct{
		nums Sequence
		expected Sequence
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 7, 4}, []int{1, 2, 3, 4, 7}},
		{[]int{1, 2, 3, 7, -4}, []int{-4, 1, 2, 3, 7}},
	}
	
	for _, c := range cases {
		oldNums := make(Sequence, len(c.nums))
		copy(oldNums, c.nums)
		
		sort.Sort(c.nums)
		for i, num := range c.nums {
			if num != c.expected[i] {
				t.Errorf("sort.Sort(%q) = %q, excepted is %q", oldNums, c.nums, c.expected)
				break
			}
		}
	}
}

func TestWhatIsType(t *testing.T) {
	cases := []interface{}{
		"abc",
		123,
		Sequence{1, 2, 3},
	}
	
	for _, c := range cases {
		result := WhatIsType(c)
		t.Log(result)
	}
}