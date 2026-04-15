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
	// Exchange stores contest-specific exchange data as key-value pairs
	// Examples:
	//   Field Day: {"class": "2A", "section": "CA"}
	//   Sweepstakes: {"serial": "123", "precedence": "A", "check": "85", "section": "CA"}
	//   CQ WW: {"zone": "5"}
	// This exchange schema is intentionally flexible to accommodate various contest formats without needing to change the underlying data structure.
	Exchange map[string]string `json:"exchange,omitempty"`
	// ExtensionData allows for additional fields that are not vital to the contact itself. I.e. operator, notes, etc.
	ExtensionData map[string]string `json:"extension_data,omitempty"`
}

func EqualContact(c1, c2 Contact) bool {
	if c1.Callsign != c2.Callsign || c1.Band != c2.Band || c1.Mode != c2.Mode || c1.Timestamp != c2.Timestamp {
		return false
	}

	// Check Exchange maps for equality
	if len(c1.Exchange) != len(c2.Exchange) {
		return false
	}
	for key, value := range c1.Exchange {
		if c2.Exchange[key] != value {
			return false
		}
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
