package store

import (
	"bgall.dev/logpaq/crdt"
)

func (s *Store) WriteVClock(vclock crdt.VClock, updatedAt int64) error {
	for nodeID, seq := range vclock {
		_, err := s.db.Exec(`
			INSERT INTO vclock (node_id, seq)
			VALUES (?, ?)
			ON CONFLICT(node_id) DO UPDATE SET
				seq = excluded.seq
		`, nodeID, seq)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) ReadVClock() (crdt.VClock, error) {
	rows, err := s.db.Query(`SELECT node_id, seq FROM vclock`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vclock := make(crdt.VClock)
	for rows.Next() {
		var nodeID string
		var seq uint64
		if err := rows.Scan(&nodeID, &seq); err != nil {
			return nil, err
		}
		vclock[nodeID] = seq
	}
	return vclock, nil
}
