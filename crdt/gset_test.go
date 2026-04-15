package crdt

import (
	"fmt"
	"testing"
	"time"
)

func TestNewGSet(t *testing.T) {
	gset := NewContactGSet("node1")

	if gset.NodeID != "node1" {
		t.Errorf("Expected NodeID 'node1', got '%s'", gset.NodeID)
	}

	if gset.Size() != 0 {
		t.Errorf("Expected empty GSet, got size %d", gset.Size())
	}

	if len(gset.VClock) != 0 {
		t.Errorf("Expected empty VClock, got %v", gset.VClock)
	}
}

func TestGSetAdd(t *testing.T) {
	gset := NewContactGSet("node1")

	contact := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: uint64(time.Now().Unix()),
		Exchange:  map[string]string{"rst": "599"},
	}

	// First add should succeed
	added := gset.Add(contact)
	if !added {
		t.Error("Expected Add to return true for new contact")
	}

	if gset.Size() != 1 {
		t.Errorf("Expected size 1, got %d", gset.Size())
	}

	if gset.VClock["node1"] != 1 {
		t.Errorf("Expected VClock[node1] = 1, got %d", gset.VClock["node1"])
	}

	// Second add of same contact should fail
	added = gset.Add(contact)
	if added {
		t.Error("Expected Add to return false for duplicate contact")
	}

	if gset.Size() != 1 {
		t.Errorf("Expected size still 1, got %d", gset.Size())
	}
}

func TestGSetContains(t *testing.T) {
	gset := NewContactGSet("node1")

	contact := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: uint64(time.Now().Unix()),
	}

	if gset.Contains(contact) {
		t.Error("Expected Contains to return false for non-existent contact")
	}

	gset.Add(contact)

	if !gset.Contains(contact) {
		t.Error("Expected Contains to return true for added contact")
	}
}

func TestGSetMerge(t *testing.T) {
	gset1 := NewContactGSet("node1")
	gset2 := NewContactGSet("node2")

	// Add different contacts to each set
	contact1 := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: 1000,
	}

	contact2 := Contact{
		Callsign:  "W2DEF",
		Band:      "40M",
		Mode:      "SSB",
		Timestamp: 2000,
	}

	gset1.Add(contact1)
	gset2.Add(contact2)

	// Merge gset2 into gset1
	merged := gset1.Merge(gset2)

	if merged.Size() != 2 {
		t.Errorf("Expected merged size 2, got %d", merged.Size())
	}

	if !merged.Contains(contact1) {
		t.Error("Expected merged set to contain contact1")
	}

	if !merged.Contains(contact2) {
		t.Error("Expected merged set to contain contact2")
	}

	// Check that merge is commutative
	merged2 := gset2.Merge(gset1)

	if merged.Size() != merged2.Size() {
		t.Errorf("Merge not commutative: sizes %d vs %d", merged.Size(), merged2.Size())
	}
}

func TestGSetConflictResolution(t *testing.T) {
	gset1 := NewContactGSet("node1")
	gset2 := NewContactGSet("node2")

	// Same contact added to both sets
	contact := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: 1000,
	}

	gset1.Add(contact)
	gset2.Add(contact)

	// Merge should resolve conflict
	merged := gset1.Merge(gset2)

	if merged.Size() != 1 {
		t.Errorf("Expected merged size 1 after conflict resolution, got %d", merged.Size())
	}

	if !merged.Contains(contact) {
		t.Error("Expected merged set to contain the contact")
	}
}

func TestGSetDelta(t *testing.T) {
	gset1 := NewContactGSet("node1")
	gset2 := NewContactGSet("node2")

	// Add contacts to gset1
	contact1 := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: 1000,
	}

	contact2 := Contact{
		Callsign:  "W2DEF",
		Band:      "40M",
		Mode:      "SSB",
		Timestamp: 2000,
	}

	gset1.Add(contact1)
	gset1.Add(contact2)

	// gset2 only knows about contact1
	gset2.Add(contact1)

	// Delta should contain only contact2
	delta := gset1.Delta(gset2)

	if delta.Size() != 1 {
		t.Errorf("Expected delta size 1, got %d", delta.Size())
	}

	if !delta.Contains(contact2) {
		t.Error("Expected delta to contain contact2")
	}

	if delta.Contains(contact1) {
		t.Error("Expected delta to NOT contain contact1")
	}
}

func TestGSetVectorClockSync(t *testing.T) {
	gset1 := NewContactGSet("node1")
	gset2 := NewContactGSet("node2")

	// Add contacts and verify vector clock advancement
	contact1 := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: 1000,
	}

	gset1.Add(contact1)
	if gset1.VClock["node1"] != 1 {
		t.Errorf("Expected node1 clock = 1, got %d", gset1.VClock["node1"])
	}

	contact2 := Contact{
		Callsign:  "W2DEF",
		Band:      "40M",
		Mode:      "SSB",
		Timestamp: 2000,
	}

	gset2.Add(contact2)
	if gset2.VClock["node2"] != 1 {
		t.Errorf("Expected node2 clock = 1, got %d", gset2.VClock["node2"])
	}

	// After merge, both clocks should be present
	merged := gset1.Merge(gset2)

	if merged.VClock["node1"] != 1 {
		t.Errorf("Expected merged node1 clock = 1, got %d", merged.VClock["node1"])
	}

	if merged.VClock["node2"] != 1 {
		t.Errorf("Expected merged node2 clock = 1, got %d", merged.VClock["node2"])
	}
}

func TestContactKeysConsistent(t *testing.T) {
	// Same contact should generate same key regardless of node
	contact := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "cw", // lowercase to test case normalization
		Timestamp: 1000,
		Exchange:  map[string]string{"rst": "599", "serial": "001"},
	}

	key1 := ContactKeyFunc(contact)

	// Modify case - should still generate same key
	contact.Callsign = "w1abc"
	contact.Band = "20m"
	contact.Mode = "CW"
	key2 := ContactKeyFunc(contact)

	if key1 != key2 {
		t.Errorf("Contact keys should be case-insensitive: %s vs %s", key1, key2)
	}
}

func TestGSetIdempotence(t *testing.T) {
	gset1 := NewContactGSet("node1")

	contact := Contact{
		Callsign:  "W1ABC",
		Band:      "20M",
		Mode:      "CW",
		Timestamp: 1000,
	}

	gset1.Add(contact)

	// Merging with self should be idempotent
	merged := gset1.Merge(gset1)

	if merged.Size() != 1 {
		t.Errorf("Expected idempotent merge size 1, got %d", merged.Size())
	}

	if merged.VClock["node1"] != gset1.VClock["node1"] {
		t.Error("Idempotent merge should preserve vector clock")
	}
}

func TestGenericGSet(t *testing.T) {
	// Test that GSet works with types other than Contact
	type TestItem struct {
		ID        int
		Name      string
		Timestamp uint64
	}

	keyFunc := func(item TestItem) string {
		return fmt.Sprintf("item:%d:%s", item.ID, item.Name)
	}

	gset := NewGSet("node1", keyFunc)

	item1 := TestItem{ID: 1, Name: "test", Timestamp: 1000}
	item2 := TestItem{ID: 2, Name: "test2", Timestamp: 2000}

	// Test adding items
	if !gset.Add(item1) {
		t.Error("Expected Add to return true for new item")
	}

	if gset.Add(item1) {
		t.Error("Expected Add to return false for duplicate item")
	}

	if !gset.Add(item2) {
		t.Error("Expected Add to return true for second item")
	}

	// Test size and contains
	if gset.Size() != 2 {
		t.Errorf("Expected size 2, got %d", gset.Size())
	}

	if !gset.Contains(item1) {
		t.Error("Expected Contains to return true for item1")
	}

	if !gset.Contains(item2) {
		t.Error("Expected Contains to return true for item2")
	}

	// Test GetItems
	items := gset.GetItems()
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
}

func TestTimestampExtraction(t *testing.T) {
	// Test that timestamp extraction works with reflection
	type WithTimestamp struct {
		Name      string
		Timestamp uint64
	}

	type WithoutTimestamp struct {
		Name string
		ID   int
	}

	keyFunc1 := func(item WithTimestamp) string { return item.Name }
	keyFunc2 := func(item WithoutTimestamp) string { return item.Name }

	gset1 := NewGSet("node1", keyFunc1)
	gset2 := NewGSet("node1", keyFunc2)

	// Item with timestamp should use its timestamp
	itemWithTS := WithTimestamp{Name: "test", Timestamp: 12345}
	gset1.Add(itemWithTS)

	entries1 := gset1.GetItemEntries()
	if len(entries1) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries1))
	}
	if entries1[0].Timestamp != 12345 {
		t.Errorf("Expected timestamp 12345, got %d", entries1[0].Timestamp)
	}

	// Item without timestamp should use current time (approximately)
	itemWithoutTS := WithoutTimestamp{Name: "test", ID: 1}
	beforeTime := uint64(time.Now().Unix())
	gset2.Add(itemWithoutTS)
	afterTime := uint64(time.Now().Unix())

	entries2 := gset2.GetItemEntries()
	if len(entries2) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(entries2))
	}

	timestamp := entries2[0].Timestamp
	if timestamp < beforeTime || timestamp > afterTime+1 {
		t.Errorf("Expected timestamp between %d and %d, got %d", beforeTime, afterTime+1, timestamp)
	}
}
