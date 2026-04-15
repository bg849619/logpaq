package crdt

import "testing"

// Test VClock single-node increment.
func TestIncrement(t *testing.T) {
	v := VClock{"node-a": 5, "node-b": 3}
	v.Increment("node-a")
	if v["node-a"] != 6 {
		t.Errorf("VClock intended node incremented wrongly, got: %d, want: %d", v["node-a"], 6)
	}
}

// Test merging the state of two VClocks
func TestMerge(t *testing.T) {
	vc := VClock{"node-a": 1, "node-b": 2, "node-c": 3}
	vo := VClock{"node-a": 3, "node-b": 2, "node-c": 1}
	vc.Merge(vo)

	if vc["node-a"] != 3 || vc["node-b"] != 2 || vc["node-c"] != 3 {
		t.Errorf("VClocks did not merge correctly.")
	}
}

// Test relationships between VClocks
func TestCompare(t *testing.T) {
	// Base vClock
	vc1 := VClock{"node-a": 1, "node-b": 2}
	// vc2 should be after vc1. Therefore vc1 should be before vc2.
	vc2 := VClock{"node-a": 2, "node-b": 3}
	// vc3 should be concurrent with vc1
	vc3 := VClock{"node-a": 2, "node-b": 1}
	// vc4 should be equal to vc1
	vc4 := VClock{"node-a": 1, "node-b": 2}

	if vc1.Compare(vc2) != Before {
		t.Errorf("Expected vc1 to be before vc2, got: %s", vc1.Compare(vc2))
	}
	if vc2.Compare(vc1) != After {
		t.Errorf("Expected vc2 to be after vc1, got: %s", vc2.Compare(vc1))
	}
	if vc1.Compare(vc3) != Concurrent {
		t.Errorf("Expected vc1 to be concurrent with vc3, got: %s", vc1.Compare(vc3))
	}
	if vc1.Compare(vc4) != Equal {
		t.Errorf("Expected vc1 to be equal to vc4, got: %s", vc1.Compare(vc4))
	}
}

func TestDelta(t *testing.T) {
	vc1 := VClock{"node-a": 1, "node-b": 2}
	vc2 := VClock{"node-a": 3, "node-b": 1}

	delta := vc1.Delta(vc2)

	// Shows that VC1 is up to 2 events (1 ahead) for node-b compared to VC2.
	if delta["node-a"] != 0 || delta["node-b"] != 2 {
		t.Errorf("Delta calculation is incorrect, got: %v", delta)
	}

	// Shows that VC2 is up to 3 events (2 ahead) for node-a compared to VC1.
	delta2 := vc2.Delta(vc1)
	if delta2["node-a"] != 3 || delta2["node-b"] != 0 {
		t.Errorf("Delta calculation is incorrect, got: %v", delta2)
	}
}
