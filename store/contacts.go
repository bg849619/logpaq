package logpaq

import (
	"encoding/json"

	"bgall.dev/logpaq/crdt"
)

func (s *Store) WriteContact(contact crdt.Contact, updatedAt int64) error {
	sentJSON, err := json.Marshal(contact.SentExchange)
	if err != nil {
		return err
	}

	rcvdJSON, err := json.Marshal(contact.ReceivedExchange)
	if err != nil {
		return err
	}

	extJSON, err := json.Marshal(contact.ExtensionData)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`
		INSERT INTO contacts (callsign, band, mode, logged_at, updated_at, sent_exchange, rcvd_exchange, extension_data)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(callsign, band, mode) DO UPDATE SET
			updated_at = excluded.updated_at,
			sent_exchange = excluded.sent_exchange,
			rcvd_exchange = excluded.rcvd_exchange,
			extension_data = excluded.extension_data
	`, contact.Callsign, contact.Band, contact.Mode, contact.Timestamp, updatedAt, string(sentJSON), string(rcvdJSON), string(extJSON))
	return err
}

func (s *Store) AllContacts() ([]crdt.Contact, error) {
	rows, err := s.db.Query(`SELECT callsign, band, mode, logged_at, sent_exchange, rcvd_exchange, extension_data FROM contacts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []crdt.Contact
	for rows.Next() {
		var c crdt.Contact
		var sentJSON, rcvdJSON, extJSON string
		if err := rows.Scan(&c.Callsign, &c.Band, &c.Mode, &c.Timestamp, &sentJSON, &rcvdJSON, &extJSON); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(sentJSON), &c.SentExchange); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(rcvdJSON), &c.ReceivedExchange); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(extJSON), &c.ExtensionData); err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}
