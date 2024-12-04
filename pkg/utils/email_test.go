package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		// Valid email tests
		{"test@example.com", true},
		{"user.name@domain.co", true},
		{"user@subdomain.example.com", true},

		// Invalid email tests
		{"invalid-email", false},
		{"@missingusername.com", false},
		{"missingat.com", false},
		{"user@domain@domain.com", false},
		{"user@domain,com", false},

		// Empty email test
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := ValidateEmail(tt.email)
			assert.Equal(t, tt.expected, result, "Expected result for %s to be %v", tt.email, tt.expected)
		})
	}
}
