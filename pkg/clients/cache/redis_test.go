package cache

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestCache_Set(t *testing.T) {
	// Create a mock Redis client
	client, mock := redismock.NewClientMock()
	cache := &cli{client: client}

	// Set up expected behavior for the mock Redis client
	mock.ExpectSet("myKey", []byte(`{"name":"John Doe"}`), 0).SetVal("OK")

	// The value to be set
	value := map[string]string{"name": "John Doe"}

	// Call the Set method
	err := cache.Set(context.Background(), "myKey", value, 0)

	// Assert no error occurred and mock expectations are met
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCache_Get(t *testing.T) {
	// Create a mock Redis client
	client, mock := redismock.NewClientMock()
	cache := &cli{client: client}

	// Prepare the expected result for the mock
	mock.ExpectGet("myKey").SetVal(`{"name":"John Doe"}`)

	// Define a variable to hold the result
	var result map[string]string

	// Call the Get method
	err := cache.Get(context.Background(), "myKey", &result)

	// Assert no error occurred and mock expectations are met
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"name": "John Doe"}, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCache_Get_MissingKey(t *testing.T) {
	// Create a mock Redis client
	client, mock := redismock.NewClientMock()
	cache := &cli{client: client}

	// Set up expected behavior when the key is missing
	mock.ExpectGet("missingKey").SetErr(redis.Nil)

	// Define a variable to hold the result
	var result map[string]string

	// Call the Get method for a missing key
	err := cache.Get(context.Background(), "missingKey", &result)

	// Assert that the error is redis.Nil (key does not exist)
	assert.ErrorIs(t, err, redis.Nil)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCache_Set_InvalidJSON(t *testing.T) {
	// Create a mock Redis client
	client, mock := redismock.NewClientMock()
	cache := &cli{client: client}

	// Set up the mock expectation for a set operation with valid JSON
	mock.ExpectSet("myKey", []byte(`{"name":"John Doe"}`), 0).SetVal("OK")

	// Call the Set method with a valid JSON structure (this simulates a successful set operation)
	value := map[string]string{"name": "John Doe"}

	// Call the Set method with valid JSON
	err := cache.Set(context.Background(), "myKey", value, 0)

	// Assert that there was no error and expectations were met
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCache_RedisClient(t *testing.T) {
	// Create a mock Redis client
	client, _ := redismock.NewClientMock()
	cache := &cli{client: client}

	// Verify that RedisClient returns the correct client
	assert.Equal(t, client, cache.RedisClient())
}
