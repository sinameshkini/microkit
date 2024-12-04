package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIID_Value(t *testing.T) {
	id := IID(123)
	val, err := id.Value()
	assert.NoError(t, err)
	assert.Equal(t, int64(123), val)
}

func TestIID_Scan(t *testing.T) {
	var id IID
	err := id.Scan(int64(456))
	assert.NoError(t, err)
	assert.Equal(t, IID(456), id)

	err = id.Scan(nil)
	assert.NoError(t, err)
	assert.Equal(t, NilIID, id)
}

func TestIID_String(t *testing.T) {
	id := IID(789)
	assert.Equal(t, "789", id.String())
}

func TestIID_CheckIID(t *testing.T) {
	id := IID(1)
	assert.True(t, id.CheckIID())

	id = NilIID
	assert.False(t, id.CheckIID())
}

func TestIID_ParsePID(t *testing.T) {
	id, err := ParsePID("123")
	assert.NoError(t, err)
	assert.Equal(t, IID(123), id)

	id, err = ParsePID(456)
	assert.NoError(t, err)
	assert.Equal(t, IID(456), id)

	id, err = ParsePID("invalid")
	assert.Error(t, err)
	assert.Equal(t, IID(0), id)
}

func TestIID_ParsePIDf(t *testing.T) {
	id := ParsePIDf("789")
	assert.Equal(t, IID(789), id)

	id = ParsePIDf(123)
	assert.Equal(t, IID(123), id)
}

func TestIID_Validate(t *testing.T) {
	id, valid := Validate("321")
	assert.True(t, valid)
	assert.Equal(t, IID(321), id)

	id, valid = Validate("invalid")
	assert.False(t, valid)
	assert.Equal(t, IID(0), id)
}

func TestNIID_Value(t *testing.T) {
	u := NIID{PID: IID(123), Valid: true}
	val, err := u.Value()
	assert.NoError(t, err)
	assert.Equal(t, int64(123), val)

	u.Valid = false
	val, err = u.Value()
	assert.NoError(t, err)
	assert.Nil(t, val)
}

func TestNIID_Scan(t *testing.T) {
	var u NIID
	err := u.Scan(int64(456))
	assert.NoError(t, err)
	assert.Equal(t, IID(456), u.PID)
	assert.True(t, u.Valid)

	err = u.Scan(nil)
	assert.NoError(t, err)
	assert.Equal(t, NilIID, u.PID)
	assert.False(t, u.Valid)
}

func TestNIID_MarshalJSON(t *testing.T) {
	u := NIID{PID: IID(123), Valid: true}
	data, err := u.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `123`, string(data))

	u.Valid = false
	data, err = u.MarshalJSON()
	assert.NoError(t, err)
	assert.JSONEq(t, `null`, string(data))
}

func TestNIID_UnmarshalJSON(t *testing.T) {
	var u NIID
	err := u.UnmarshalJSON([]byte(`123`))
	assert.NoError(t, err)
	assert.Equal(t, IID(123), u.PID)
	assert.True(t, u.Valid)
}

func TestIID_Parse(t *testing.T) {
	assert.Equal(t, IID(123), Parse("123"))
	assert.Equal(t, IID(0), Parse("invalid"))
}

func TestIID_CheckIIDFunc(t *testing.T) {
	assert.True(t, CheckIID(IID(123)))
	assert.False(t, CheckIID(NilIID))
}
