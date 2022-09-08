package helpers

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/mail"
)

const (
	// // Context Headers ------------------------------------------------------------------

	// HeadersUser ...
	HeadersUser string = "x-user"

	// HeadersUserID ...
	HeadersUserID string = "x-user-id"

	// HeaderUserName
	HeaderUserName string = "x-username"

	// // HeadersClient ...
	// HeadersClient string = "x-client"

	// HeadersAuthenticated ...
	HeadersAuthenticated string = "x-authenticated"
	// HeadersTokenKey ...
	HeadersTokenKey string = "x-token-key"
)

var (
	// alphanumericChars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// letterChars       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	verificationChars = []byte("123456789")
	digitChars        = []byte("0123456789")
)

func JsonAssertion(src interface{}, dst interface{}) (err error) {
	var js []byte

	js, err = json.Marshal(src)
	if err != nil {
		return
	}

	err = json.Unmarshal(js, dst)
	if err != nil {
		return
	}

	return
}

func ValidateEmail(email string) (isValid bool) {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// GenerateVerificationCode ...
func GenerateVerificationCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = verificationChars[int(b[i])%len(verificationChars)]
	}
	return string(b)
}

// GenerateCode ...
func GenerateCode(max int) (string, error) {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = digitChars[int(b[i])%len(digitChars)]
	}

	return string(b), nil
}
