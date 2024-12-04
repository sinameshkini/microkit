package models

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

// IID Incremental Primary SID
type IID int64

// NilIID Null Primary SID
var NilIID = IID(0)

// NIID can be used with the standard sql package to represent a
// UUID value that can be NULL in the database
type NIID struct {
	PID   IID
	Valid bool
}

// Value - Implementation of valuer for database/sql
func (id *IID) Value() (driver.Value, error) {
	// value needs to be a base driver.Value type
	// such as string, bool and ...
	return int64(*id), nil
}

// Scan implements the sql.Scanner interface.
// A 16-byte slice is handled by UnmarshalBinary, while
// a longer byte slice or a string is handled by UnmarshalText.
func (id *IID) Scan(src interface{}) error {
	if src == nil {
		*id = IID(0)
		return nil
	}

	*id = IID(src.(int64))

	return nil
}

func (id *IID) String() string {
	return strconv.Itoa(int(*id))
}

// CheckPID ...
func (id *IID) CheckPID() bool {
	return true
}

func (id *IID) IsValid() bool {
	return int64(*id) > 0
}

// ParsePID , parses a string id to a IID one
func ParsePID(id interface{}) (pid IID, err error) {
	switch id.(type) {
	case string:
		var d int
		if d, err = strconv.Atoi(id.(string)); err != nil {
			return 0, err
		}
		pid = IID(d)
	case int:
		pid = IID(id.(int))
	case float64:
		pid = IID(id.(float64))
	case IID:
		pid = id.(IID)
	}

	if !pid.IsValid() {
		err = ErrInvalidID
	}

	return pid, err
}

func ParsePIDf(id interface{}) (pid IID) {
	var err error
	switch v := id.(type) {
	case string:
		var d int
		if d, err = strconv.Atoi(v); err != nil {
			return
		}

		pid = IID(d)
	case int:
		pid = IID(id.(int))
	case float64:
		pid = IID(id.(float64))
	case IID:
		pid = id.(IID)
	}

	return
}

// Parse ...
func Parse(id string) IID {
	pid, _ := ParsePID(id)
	return pid
}

// Validate ...
func Validate(id string) (IID, bool) {
	pid, err := ParsePID(id)
	return pid, err == nil
}

// String ...
func String(id *IID) string {
	return id.String()
}

// CheckPID ...
func CheckPID(id *IID) bool {
	return id.CheckPID()
}

// Value implements the driver.Valuer interface.
func (u *NIID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	// Delegate to int64 Value function
	return u.PID.Value()
}

// Scan implements the sql.Scanner interface.
func (u *NIID) Scan(src interface{}) error {
	if src == nil {
		u.PID, u.Valid = NilIID, false
		return nil
	}

	// Delegate to int64 Scan function
	u.Valid = true
	return u.PID.Scan(src)
}

// MarshalJSON ...
func (u *NIID) MarshalJSON() ([]byte, error) {
	if u.Valid {
		return json.Marshal(u.PID)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (u *NIID) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *IID
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
