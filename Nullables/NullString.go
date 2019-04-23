package Nullables

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("\"\""), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// MakeNew .. creates s nullstring
func (ns *NullString) MakeNew(s string) {
	bytes := []byte(fmt.Sprintf("\"%s\"", s))
	err := json.Unmarshal(bytes, &ns.String)
	ns.Valid = (err == nil)
}
