package effective

import "testing"

func TestContent(t *testing.T) {
	filename := "/Users/dongsj/workspace/Bennett/goSpace/go-learn/src/resources/defer.test"
	
	result, err := Content(filename)
	if err != nil {
		t.Fail()
	}
	if result != "How to use defer" {
		t.Fail()
	}
}
