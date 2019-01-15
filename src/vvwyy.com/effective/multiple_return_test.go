package effective

import "testing"

func TestNextNum(t *testing.T) {
	cases := []struct{
		input                    string
		pos                      int
		expectedNum, expectedPos int
	} {
		{"abc123d456", 0, 123, 6},
		{"abc123d456", 3, 123, 6},
		{"abc123d456", 4, 23, 6},
		{"abc123d456", 6, 456, 10},
	}
	
	for _, c := range cases{
		num, pos := NextNum([]byte(c.input), c.pos)
		if c.expectedNum != num || c.expectedPos != pos {
			t.Errorf("NextNum(%q, %d) = %d, %d ; expected is %d, %d", c.input, c.pos, num, pos, c.expectedNum, c.expectedPos)
		}
	}
	
}