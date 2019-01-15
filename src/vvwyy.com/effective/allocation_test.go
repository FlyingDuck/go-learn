package effective

import "testing"

func TestAppend(t *testing.T) {
	
	cases := []struct {
		sourceSlice []byte
		appendData  []byte
		expected    string
	}{
		{[]byte("abc"), []byte("abc"), "abcabc"},
		{[]byte("abc"), []byte("123"), "abc123"},
		{[]byte(""), []byte("123"), "123"},
		{[]byte("abc"), []byte(""), "abc"},
		{[]byte("abc"), []byte("d"), "abcd"},
	}
	
	for _, c := range cases {
		result := Append(c.sourceSlice, c.appendData)
		if string(result) != c.expected {
			t.Errorf("Append(%q, %q) = %q; expected is %q", c.sourceSlice, c.appendData, result, c.expected)
		}
	}
}
