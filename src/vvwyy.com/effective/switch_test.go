package effective

import "testing"

func TestSwitchCompare(t *testing.T) {
	cases := []struct {
		inA, inB string
		expected int
	}{
		{"abc", "abc", 0},
		{"","", 0},
		{"abc", "ab", 1},
		{"abc", "abcd", -1},
		{"cde", "abc", 1},
		{"cde", "", 1},
		{"cde", "e", -1},
	}
	
	for _, c := range cases {
		got := Compare([]byte(c.inA), []byte(c.inB))
		if got != c.expected {
			t.Errorf("Compare(%q, %q) = %d, expected is %d", c.inA, c.inB, got, c.expected)
		}
	}
}

func TestWhatsType(t *testing.T) {
	t1 := WhatsType("string")
	if "string" != t1 {
		t.Fail()
	}
	t1 = WhatsType(123)
	if "int" != t1 {
		t.Fail()
	}
}


