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
func TestRelation(t *testing.T) {

}
