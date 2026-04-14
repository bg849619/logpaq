package crdt

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// VClock - Map of Node IDs to a count of their logs
// Tracks casual consistency between nodes.
// - Increment your own clock when logging
// - When syncing, merge by taking the maximum clock for each element.
type VClock map[string]uint64

// Empty vector clock.
func New() VClock {
	return make(VClock)
}

// Increment
func (vc VClock) Increment(nodeID string) {
	vc[nodeID]++
}

// Merge with the maximum clock value for each
func (vc VClock) Merge(other VClock) {
	for nodeID, otherSeq := range other {
		if otherSeq > vc[nodeID] {
			vc[nodeID] = otherSeq
		}
	}
}

type Relation int

const (
	// This clock's events are already recorded on the other side.
	Before Relation = iota

	// This clock only has additional information compared to the other.
	After

	// Identical histories
	Equal

	// Diverging/partition scenarios.
	Concurrent
)

// Because you can never have enough string functions.
func (r Relation) String() string {
	switch r {
	case Before:
		return "before"
	case After:
		return "after"
	case Equal:
		return "equal"
	case Concurrent:
		return "concurrent"
	default:
		return "Unknown"
	}
}

// Compare this VC to another.
func (vc VClock) Compare(other VClock) Relation {
	allNodes := make(map[string]struct{})
	// Start by filling with our own clock's known nodes. Then fill with the other clock in case there are unknown nodes.
	for k := range vc {
		allNodes[k] = struct{}{}
	}
	for k := range other {
		allNodes[k] = struct{}{}
	}

	vcGreater := false    // Ours has at least one greater than other
	otherGreater := false // Other has at least one greater than ours.

	for nodeID := range allNodes {
		vcVal := vc[nodeID]
		otherVal := other[nodeID]
		if vcVal > otherVal {
			vcGreater = true
		} else if otherVal > vcVal {
			otherGreater = true
		}
		if vcGreater && otherGreater {
			// Not *necessarily* needed here, but is technically a shortcut.
			return Concurrent
		}
	}

	switch {
	case !vcGreater && !otherGreater:
		return Equal
	case vcGreater && !otherGreater:
		return After
	case !vcGreater && otherGreater:
		return Before
	default:
		return Concurrent // Theoretically shouldn't be met?
	}
}

// Returns which Node IDs which *we've* seen more events than the other clock.
// This would be used by another mechanism to determine which events to send back to the other clock.
func (vc VClock) Delta(other VClock) map[string]uint64 {
	delta := make(map[string]uint64)
	for nodeID, seq := range vc {
		if seq > other[nodeID] {
			delta[nodeID] = seq
		}
	}
	return delta
}

// Has this clock seen an event (seq) reported from (nodeID)
// Check this before trying to merge a contact to the local datastore
func (vc VClock) HasSeen(nodeID string, seq uint64) bool {
	return vc[nodeID] >= seq
}

// Serialize the clock.
func (vc VClock) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]uint64(vc))
}

func (vc *VClock) UnmarshalJSON(data []byte) error {
	m := make(map[string]uint64)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*vc = VClock(m)
	return nil
}

// String rep sorted by node ID
// Wouldn't JSON.Marshal already do this?
func (vc VClock) String() string {
	if len(vc) == 0 {
		return "{}" // fast.
	}

	// Collect keys and sort
	keys := make([]string, 0, len(vc))
	for k := range vc {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Claude wanted to put an insertion sort in here lol.

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s:%d", k, vc[k]))
	}

	return "{" + strings.Join(parts, ", ") + "}"
}
