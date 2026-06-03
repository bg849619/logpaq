package crdt

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Contact struct {
	Callsign  string `json:"callsign"`
	Band      string `json:"band"`
	Mode      string `json:"mode"`
	Timestamp uint64 `json:"timestamp"`
	// SentExchange stores the exchange sent to the other station.
	SentExchange map[string]string `json:"sent_exchange,omitempty"`
	// ReceivedExchange stores the exchange received from the other station.
	ReceivedExchange map[string]string `json:"rcvd_exchange,omitempty"`
	// ExtensionData allows for additional fields that are not vital to the contact itself. I.e. operator, notes, etc.
	ExtensionData map[string]string `json:"extension_data,omitempty"`
}

func EqualContact(c1, c2 Contact) bool {
	if c1.Callsign != c2.Callsign || c1.Band != c2.Band || c1.Mode != c2.Mode || c1.Timestamp != c2.Timestamp {
		return false
	}

	compareMaps := func(left, right map[string]string) bool {
		if len(left) != len(right) {
			return false
		}
		for key, value := range left {
			if right[key] != value {
				return false
			}
		}
		return true
	}

	if !compareMaps(c1.SentExchange, c2.SentExchange) {
		return false
	}
	if !compareMaps(c1.ReceivedExchange, c2.ReceivedExchange) {
		return false
	}

	// Check ExtensionData maps for equality
	if len(c1.ExtensionData) != len(c2.ExtensionData) {
		return false
	}
	for key, value := range c1.ExtensionData {
		if c2.ExtensionData[key] != value {
			return false
		}
	}

	return true
}

func MarshalJSON(c Contact) ([]byte, error) {
	return json.Marshal(c)
}

func UnmarshalJSON(data []byte) (Contact, error) {
	var c Contact
	err := json.Unmarshal(data, &c)
	return c, err
}

// ContactKeyFunc generates a deterministic key for a contact
// This ensures the same contact from different nodes gets the same key
// In contest logging: callsign + band + mode should be unique (no dupes allowed)
func ContactKeyFunc(contact Contact) string {
	return fmt.Sprintf("%s:%s:%s",
		strings.ToUpper(contact.Callsign),
		strings.ToUpper(contact.Band),
		strings.ToUpper(contact.Mode))
}

// NewContactGSet creates a new GSet specifically for Contact types
// The generic GSet will automatically detect and use Contact.Timestamp field
func NewContactGSet(nodeID string) *GSet[Contact] {
	return NewGSet(nodeID, ContactKeyFunc)
}
