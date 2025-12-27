package util

import (
	"database/sql"
	"encoding/json"
)

type (
	// NullString is an alias for sql.NullString which implements JSON marshalling and unmarshalling.
	NullString sql.NullString
	// NullBool is an alias for sql.NullString which implements JSON marshalling and unmarshalling.
	NullBool sql.NullBool
	// NullInt32 is an alias for sql.NullString which implements JSON marshalling and unmarshalling.
	NullInt32 sql.NullInt32
)

func (s *NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

func (s *NullString) UnmarshalJSON(data []byte) error {
	var ps *string
	if err := json.Unmarshal(data, &ps); err != nil {
		return err
	}
	if ps != nil {
		s.Valid = true
		s.String = *ps
	} else {
		s.Valid = false
	}
	return nil
}

func (b *NullBool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

func (b *NullBool) UnmarshalJSON(data []byte) error {
	var pb *bool
	if err := json.Unmarshal(data, &pb); err != nil {
		return err
	}
	if pb != nil {
		b.Valid = true
		b.Bool = *pb
	} else {
		b.Valid = false
	}
	return nil
}

func (i *NullInt32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int32)
}

func (i *NullInt32) UnmarshalJSON(data []byte) error {
	var pi *int32
	if err := json.Unmarshal(data, &pi); err != nil {
		return err
	}
	if pi != nil {
		i.Valid = true
		i.Int32 = *pi
	} else {
		i.Valid = false
	}
	return nil
}
