package crdt

import "fmt"

type LWWField[T any] struct {
	Value     T      `json:"value"`
	Timestamp uint64 `json:"timestamp"`
	NodeID    string `json:"node_id,omitempty"` // For tie-breaking
}

func (f LWWField[T]) Merge(other LWWField[T]) LWWField[T] {
	// Prefer higher timestamp (more recent)
	if other.Timestamp > f.Timestamp {
		return other
	}
	if f.Timestamp > other.Timestamp {
		return f
	}
	
	// Same timestamp, use node ID for deterministic tie-breaking
	if f.NodeID <= other.NodeID {
		return f
	}
	return other
}

func (f LWWField[T]) String() string {
	if f.NodeID != "" {
		return fmt.Sprintf("%v@%d[%s]", f.Value, f.Timestamp, f.NodeID)
	}
	return fmt.Sprintf("%v@%d", f.Value, f.Timestamp)
}
