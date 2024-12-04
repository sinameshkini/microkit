package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test JsonAssertion function
func TestJsonAssertion(t *testing.T) {
	tests := []struct {
		name      string
		src       interface{}
		dest      interface{}
		expectErr bool
	}{
		{
			name:      "Valid JSON marshaling and unmarshaling",
			src:       map[string]interface{}{"key": "value"},
			dest:      &map[string]interface{}{},
			expectErr: false,
		},
		{
			name:      "Invalid destination type",
			src:       map[string]interface{}{"key": "value"},
			dest:      nil,
			expectErr: true,
		},
		{
			name:      "Empty source",
			src:       nil,
			dest:      &map[string]interface{}{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := JsonAssertion(tt.src, tt.dest)
			if tt.expectErr {
				assert.Error(t, err, "expected error but got none")
			} else {
				assert.NoError(t, err, "expected no error but got one")
			}
		})
	}
}

// Test JsonIndent function
func TestJsonIndent(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "Indent simple object",
			input:    map[string]interface{}{"key": "value"},
			expected: "{\n\t\"key\": \"value\"\n}",
		},
		{
			name:     "Indent empty object",
			input:    map[string]interface{}{},
			expected: "{}",
		},
		{
			name:     "Indent invalid object",
			input:    func() {},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := JsonIndent(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Test Indent function
func TestIndent(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Indent valid JSON",
			input:    []byte(`{"key":"value"}`),
			expected: "{\n\t\"key\": \"value\"\n}",
		},
		{
			name:     "Indent empty JSON",
			input:    []byte(`{}`),
			expected: "{}",
		},
		{
			name:     "Indent invalid JSON",
			input:    []byte(`{key:value}`), // Invalid JSON
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Indent(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
