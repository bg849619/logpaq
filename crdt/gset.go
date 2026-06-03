package crdt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// GSet implements a grow-only set CRDT with vector clock causality
// It is generic over any type T and uses a configurable key function
type GSet[T any] struct {
	NodeID  string                 `json:"node_id"`
	Items   map[string]LWWField[T] `json:"items"`
	VClock  VClock                 `json:"vclock"`
	keyFunc func(T) string         `json:"-"` // Don't serialize the function
}

// NewGSet creates a new generic GSet instance
func NewGSet[T any](nodeID string, keyFunc func(T) string) *GSet[T] {
	return &GSet[T]{
		NodeID:  nodeID,
		Items:   make(map[string]LWWField[T]),
		VClock:  New(),
		keyFunc: keyFunc,
	}
}

// extractTimestamp extracts timestamp from items that have a Timestamp field
// Uses reflection to check for uint64 Timestamp field, falls back to current time
func (gs *GSet[T]) extractTimestamp(item T) uint64 {
	v := reflect.ValueOf(item)

	// Handle pointer types
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return uint64(time.Now().Unix())
		}
		v = v.Elem()
	}

	// Check if it's a struct with a Timestamp field
	if v.Kind() == reflect.Struct {
		timestampField := v.FieldByName("Timestamp")
		if timestampField.IsValid() && timestampField.CanInterface() {
			// Check if it's a uint64
			if timestampField.Kind() == reflect.Uint64 {
				return timestampField.Uint()
			}
		}
	}

	// Fallback to current time
	return uint64(time.Now().Unix())
}

// Add adds an item to the GSet if it doesn't already exist
func (gs *GSet[T]) Add(item T) bool {
	key := gs.keyFunc(item)

	// Check if item already exists
	if _, exists := gs.Items[key]; exists {
		return false // Item already exists
	}

	// Increment our local vector clock
	gs.VClock.Increment(gs.NodeID)

	// Create item entry with LWW semantics
	// Note: For Contact type, we'll extract timestamp; for other types, may need different approach
	entry := LWWField[T]{
		Value:     item,
		Timestamp: gs.extractTimestamp(item),
		NodeID:    gs.NodeID,
	}

	gs.Items[key] = entry

	return true
}

// Merge merges another GSet into this one
// This operation is commutative - Merge(A, B) == Merge(B, A)
func (gs *GSet[T]) Merge(other *GSet[T]) *GSet[T] {
	// First, compare vector clocks to determine the relationship
	relation := gs.VClock.Compare(other.VClock)

	switch relation {
	case Before:
		// Other is ahead of us, so other's state is more complete
		// Return other but with our NodeID if different
		if other.NodeID == gs.NodeID {
			return other
		}
		// Need to create copy with our NodeID
		result := &GSet[T]{
			NodeID:  gs.NodeID,
			Items:   other.Items,  // Share the map (read-only)
			VClock:  other.VClock, // Share the VClock (read-only)
			keyFunc: gs.keyFunc,
		}
		return result

	case After:
		// We are ahead of other, so our state is more complete
		// Just return ourselves
		return gs

	case Equal:
		// Same vector clock state, sets should be identical
		// Just return ourselves
		return gs

	case Concurrent:
		// Concurrent updates - need to do full merge with conflict resolution
		return gs.mergeConcurrent(other)
	}

	// Should never reach here, but fallback to concurrent merge
	return gs.mergeConcurrent(other)
}

// mergeConcurrent performs a full merge when vector clocks are concurrent
func (gs *GSet[T]) mergeConcurrent(other *GSet[T]) *GSet[T] {
	result := &GSet[T]{
		NodeID:  gs.NodeID, // Keep our node ID
		Items:   make(map[string]LWWField[T]),
		VClock:  gs.VClock.copy(),
		keyFunc: gs.keyFunc,
	}

	// Merge vector clocks first
	result.VClock.Merge(other.VClock)

	// Merge items from both sets
	allKeys := make(map[string]struct{})

	// Collect all item keys
	for k := range gs.Items {
		allKeys[k] = struct{}{}
	}
	for k := range other.Items {
		allKeys[k] = struct{}{}
	}

	// Process each item
	for key := range allKeys {
		ourEntry, ourExists := gs.Items[key]
		otherEntry, otherExists := other.Items[key]

		switch {
		case ourExists && !otherExists:
			result.Items[key] = ourEntry
		case !ourExists && otherExists:
			result.Items[key] = otherEntry
		case ourExists && otherExists:
			// Both have the item, use conflict resolution
			resolved := gs.resolveItemConflict(ourEntry, otherEntry)
			result.Items[key] = resolved
		}
	}

	return result
}

// resolveItemConflict resolves conflicts using LWW semantics
func (gs *GSet[T]) resolveItemConflict(our, other LWWField[T]) LWWField[T] {
	// LWW merge handles all conflict resolution
	return our.Merge(other)
}

// Delta returns the set of items that should be sent to bring other up to date
func (gs *GSet[T]) Delta(other *GSet[T]) *GSet[T] {
	delta := NewGSet(gs.NodeID, gs.keyFunc)
	delta.VClock = gs.VClock.copy()

	// For a grow-only set, delta is simply items we have that they don't have
	// Since items are identified by key (e.g., callsign+band+mode for contacts),
	// if they have the same key, it's the same logical item regardless of who added it
	for key, ourItem := range gs.Items {
		if _, exists := other.Items[key]; !exists {
			// Other side doesn't have this item at all
			delta.Items[key] = ourItem
		}
		// If they do have the item, no need to send it in delta
		// The merge operation will handle any LWW conflict resolution
	}

	return delta
}

// Size returns the number of items in the set
func (gs *GSet[T]) Size() int {
	return len(gs.Items)
}

// Contains checks if an item exists in the set
func (gs *GSet[T]) Contains(item T) bool {
	key := gs.keyFunc(item)
	_, exists := gs.Items[key]
	return exists
}

// GetItems returns all items as a slice
func (gs *GSet[T]) GetItems() []T {
	items := make([]T, 0, len(gs.Items))
	for _, entry := range gs.Items {
		items = append(items, entry.Value)
	}
	return items
}

// GetItemEntries returns all item entries for detailed inspection
func (gs *GSet[T]) GetItemEntries() []LWWField[T] {
	entries := make([]LWWField[T], 0, len(gs.Items))
	for _, entry := range gs.Items {
		entries = append(entries, entry)
	}
	return entries
}

// copy creates a deep copy of the VClock
func (vc VClock) copy() VClock {
	copy := make(VClock)
	for k, v := range vc {
		copy[k] = v
	}
	return copy
}

// MarshalJSON implements JSON marshaling for GSet
// Note: keyFunc is not marshaled as functions can't be serialized
func (gs *GSet[T]) MarshalJSON() ([]byte, error) {
	// Create a temporary struct without the keyFunc for marshaling
	type Alias struct {
		NodeID string                 `json:"node_id"`
		Items  map[string]LWWField[T] `json:"items"`
		VClock VClock                 `json:"vclock"`
	}
	return json.Marshal(Alias{
		NodeID: gs.NodeID,
		Items:  gs.Items,
		VClock: gs.VClock,
	})
}

// UnmarshalJSON implements JSON unmarshaling for GSet
// Note: keyFunc must be set separately after unmarshaling
func (gs *GSet[T]) UnmarshalJSON(data []byte) error {
	type Alias struct {
		NodeID string                 `json:"node_id"`
		Items  map[string]LWWField[T] `json:"items"`
		VClock VClock                 `json:"vclock"`
	}
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	gs.NodeID = aux.NodeID
	gs.Items = aux.Items
	gs.VClock = aux.VClock
	// Note: keyFunc must be set by caller after unmarshaling
	return nil
}

// String returns a string representation of the GSet
func (gs *GSet[T]) String() string {
	return fmt.Sprintf("GSet{NodeID: %s, Items: %d, VClock: %v}",
		gs.NodeID, len(gs.Items), gs.VClock)
}
