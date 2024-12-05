package logger

import (
	"context"
	"errors"
	"github.com/sinameshkini/microkit/pkg/utils"
	"go.uber.org/zap/zapcore"
	"testing"
)

// Test New Logger
func TestNewLogger(t *testing.T) {
	logger, err := NewStdOut()
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	if logger == nil {
		t.Fatalf("Expected logger to be non-nil")
	}
}

// Test Adding Fields
func TestLogs_AddField(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), InfoLevel)
	log.addField(FieldUserID, "123")

	if val, ok := log.fields[FieldUserID]; !ok || val != "123" {
		t.Errorf("Expected field %v to be added with value '123'", FieldUserID)
	}
}

// Test With Fields
func TestLogs_With(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), InfoLevel)
	log.With(FieldRequestID, "req-456")

	if val, ok := log.fields[FieldRequestID]; !ok || val != "req-456" {
		t.Errorf("Expected field %v to be added with value 'req-456'", FieldRequestID)
	}
}

// Test Commit Log
func TestLogs_Commit(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), InfoLevel)
	log.With(FieldUserID, "123")
	log.Commit("Test log message")

	// No actual assertion here since it prints logs to stdout.
	// Ensure no panics/errors during execution.
}

// Test Logging With Context
func TestLogs_WithContext(t *testing.T) {
	ctx := utils.SetUserIDToCtx(context.Background(), "123")
	logger, _ := NewStdOut()

	log := logger.newLogs(ctx, InfoLevel)
	log.logContext(ctx)

	if val, ok := log.fields[FieldUserID]; !ok || val != "123" {
		t.Errorf("Expected context field %v to be set with value '123'", FieldUserID)
	}
}

// Test Error Level Log
func TestLogger_Error(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.Error(context.Background())
	log.With(FieldError, "Something went wrong").Commit("Error occurred")

	// No assertion. Validate output manually or mock zap.Logger for assertions.
}

// Test Adding Database Fields
func TestLogs_WithDBFields(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), InfoLevel)
	log.WithDBQuery("SELECT * FROM users")
	log.WithDBRowsAffected(10)
	log.WithDBDelay(120)

	if log.fields[FieldDBQuery] != "SELECT * FROM users" {
		t.Errorf("Expected DBQuery field to be 'SELECT * FROM users'")
	}
	if log.fields[FieldDBRowAffected] != 10 {
		t.Errorf("Expected DBRowAffected field to be 10")
	}
	if log.fields[FieldDBDelay] != int64(120) {
		t.Errorf("Expected DBDelay field to be 120")
	}
}

// Test Logger Level Conversion
func TestLevelToZapCoreLevel(t *testing.T) {
	tests := []struct {
		level        Level
		expectedCore zapcore.Level
	}{
		{DebugLevel, zapcore.DebugLevel},
		{InfoLevel, zapcore.InfoLevel},
		{WarnLevel, zapcore.WarnLevel},
		{ErrorLevel, zapcore.ErrorLevel},
		{PanicLevel, zapcore.DPanicLevel},
		{FatalLevel, zapcore.FatalLevel},
		{"unknown", zapcore.InvalidLevel},
	}

	for _, tt := range tests {
		t.Run(string(tt.level), func(t *testing.T) {
			if got := tt.level.toZapCoreLevel(); got != tt.expectedCore {
				t.Errorf("Expected %v, got %v", tt.expectedCore, got)
			}
		})
	}
}

// Test Logger With Error
func TestLogs_WithError(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), ErrorLevel)

	log.WithError(errors.New("test error"))

	if val, ok := log.fields[FieldError]; !ok || val != "test error" {
		t.Errorf("Expected field %v to be set with value 'test error'", FieldError)
	}
}

// Test Logger With Input
func TestLogs_WithInput(t *testing.T) {
	logger, _ := NewStdOut()
	log := logger.newLogs(context.Background(), InfoLevel)

	log.WithInput(map[string]string{"key": "value"})

	if val, ok := log.fields[FieldMethodInput]; !ok || val == nil {
		t.Errorf("Expected field %v to be set with JSON input", FieldMethodInput)
	}
}
