package database

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strconv"
)

var (
	ErrInvalidPID = errors.New("id is not valid")
)

// PID Primary ID
type PID int64

// NilPID Null Primary ID
var NilPID = PID(0)

// NullPID can be used with the standard sql package to represent a
// UUID value that can be NULL in the database
type NullPID struct {
	PID   PID
	Valid bool
}

// Value - Implementation of valuer for database/sql
func (id PID) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as string, bool and ...
	return int64(id), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice is handled by UnmarshalBinary, while
// a longer byte slice or a string is handled by UnmarshalText.
func (id *PID) Scan(src interface{}) error {
	if src == nil {
		*id = PID(0)
		return nil
	}

	// ns := sql.NullInt64{}
	// if err := ns.Scan(src); err != nil {
	//     return err
	// }
	//
	// if !ns.Valid {
	//     return errors.New("scan not valid")
	// }
	//
	// nsv, _ := ns.Value()
	// *id = PID(nsv.(int64))

	*id = PID(src.(int64))

	return nil
}

func (id PID) String() string {
	return strconv.Itoa(int(id))
}

// CheckPID ...
func (id PID) CheckPID() bool {
	return true
}

func (id PID) IsValid() bool {
	return int64(id) > 0
}

// ParsePID , parses a string id to a PID one
func ParsePID(id interface{}) (pid PID, err error) {
	switch id.(type) {
	case string:
		var d int
		if d, err = strconv.Atoi(id.(string)); err != nil {
			return 0, err
		}
		pid = PID(d)
	case int:
		pid = PID(id.(int))
	case float64:
		pid = PID(id.(float64))
	case PID:
		pid = id.(PID)
	}

	if !pid.IsValid() {
		err = ErrInvalidPID
	}

	return pid, err
}

// Parse ...
func Parse(id string) PID {
	pid, _ := ParsePID(id)
	return pid
}

// Validate ...
func Validate(id string) (PID, bool) {
	pid, err := ParsePID(id)
	return pid, err == nil
}

// String ...
func String(id PID) string {
	return id.String()
}

// CheckPID ...
func CheckPID(id PID) bool {
	return id.CheckPID()
}

// Value implements the driver.Valuer interface.
func (u NullPID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	// Delegate to int64 Value function
	return u.PID.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NullPID) Scan(src interface{}) error {
	if src == nil {
		u.PID, u.Valid = NilPID, false
		return nil
	}

	// Delegate to int64 Scan function
	u.Valid = true
	return u.PID.Scan(src)
}

// MarshalJSON ...
func (u NullPID) MarshalJSON() ([]byte, error) {
	if u.Valid {
		return json.Marshal(u.PID)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (u *NullPID) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *PID
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		u.Valid = true
		u.PID = *x
	} else {
		u.Valid = false
	}
	return nil
}
