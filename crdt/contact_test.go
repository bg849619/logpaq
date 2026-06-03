package crdt

import "testing"

func TestContactEqual(t *testing.T) {
	contact1 := Contact{
		Callsign: "W1AW",
		Band:     "20m",
		Mode:     "SSB",
		Exchange: map[string]string{
			"class":   "2A",
			"section": "CT",
		},
	}

	contact2 := Contact{
		Callsign: "W1AW",
		Band:     "20m",
		Mode:     "SSB",
		Exchange: map[string]string{
			"class":   "2A",
			"section": "CT",
		},
	}

	if !EqualContact(contact1, contact2) {
		t.Errorf("Expected contacts to be equal, but they are not. Got: %+v, want: %+v", contact1, contact2)
	}

	// Modify one field and check that they are no longer equal
	contact2.Mode = "CW"
	if EqualContact(contact1, contact2) {
		t.Errorf("Expected contacts to be different after modification, but they are still equal. Got: %+v, want: %+v", contact1, contact2)
	}
}

func TestContactJSON(t *testing.T) {
	// Build an example contact. Unmarshaling should produce the same contact data.
	contact := Contact{
		Callsign: "W1AW",
		Band:     "20m",
		Mode:     "SSB",
		Exchange: map[string]string{
			"class":   "2A",
			"section": "CT",
		},
		ExtensionData: map[string]string{
			"notes":    "Worked during Field Day 2024",
			"operator": "JBG",
		},
	}

	data, err := MarshalJSON(contact)
	if err != nil {
		t.Errorf("Failed to marshal contact: %v", err)
	}

	unmarshaledContact, err := UnmarshalJSON(data)
	if err != nil {
		t.Errorf("Failed to unmarshal contact: %v", err)
	}

	if !EqualContact(contact, unmarshaledContact) {
		t.Errorf("Unmarshaled contact does not match original. Got: %+v, want: %+v", unmarshaledContact, contact)
	}
}
