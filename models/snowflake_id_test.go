package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnowflakeID(t *testing.T) {
	// Initialize snowflake node
	InitSnowflakeID(1)

	t.Run("Generate next ID", func(t *testing.T) {
		id := NextSID()
		assert.NotEmpty(t, id, "Generated ID should not be empty")
		assert.Greater(t, id, SID(0), "Generated ID should be greater than 0")
	})

	t.Run("ID to String", func(t *testing.T) {
		id := NextSID()
		idStr := id.String()
		assert.NotEmpty(t, idStr, "ID string representation should not be empty")
	})

	t.Run("Parse valid ID string", func(t *testing.T) {
		id := NextSID()
		idStr := id.String()
		parsedID, err := ParseSID(idStr)
		assert.NoError(t, err, "Parsing a valid ID string should not return an error")
		assert.Equal(t, id, parsedID, "Parsed ID should match the original ID")
	})

	t.Run("Parse invalid ID string", func(t *testing.T) {
		_, err := ParseSID("invalid_id")
		assert.Error(t, err, "Parsing an invalid ID string should return an error")
	})

	t.Run("ParseSIDf valid ID string", func(t *testing.T) {
		id := NextSID()
		idStr := id.String()
		parsedID := ParseSIDf(idStr)
		assert.Equal(t, id, parsedID, "Parsed IDf should match the original ID")
	})

	t.Run("ParseSIDf invalid ID string", func(t *testing.T) {
		parsedID := ParseSIDf("invalid_id")
		assert.Equal(t, SID(0), parsedID, "ParseSIDf should return SID(0) for an invalid ID string")
	})
}
