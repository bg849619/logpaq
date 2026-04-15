package crdt

import "testing"

func TestLWWFieldMerge(t *testing.T) {
	field1 := LWWField[string]{Value: "Hello", Timestamp: 1}
	field2 := LWWField[string]{Value: "World", Timestamp: 2}

	merged := field1.Merge(field2)

	if merged.Value != "World" || merged.Timestamp != 2 {
		t.Errorf("LWWField merge did not return the field with the latest timestamp, got: %v", merged)
	}

	// Test merge in the other direction
	merged2 := field2.Merge(field1)

	if merged2.Value != "World" || merged2.Timestamp != 2 {
		t.Errorf("LWWField merge did not return the field with the latest timestamp, got: %v", merged2)
	}
}
