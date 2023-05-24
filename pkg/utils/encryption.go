package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt cookie value using AES-GCM with a random IV
// and 32 bytes encryption key
func Encrypt(key, value []byte) (string, error) {
	iv := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create a new AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt the value
	cipherValue := aesgcm.Seal(nil, iv, value, nil)

	// Concatenate all data together
	ecnryptedData := append(iv, cipherValue...)

	// Encode the encrypted data in base64 for transport/storage
	encodedData := base64.StdEncoding.EncodeToString(ecnryptedData)

	return encodedData, nil
}

// Decrypt cookie value using AES-GCM with a random IV
// and 32 bytes encryption key
func Decrypt(key []byte, value string) (string, error) {
	// Check if value length is valid
	if len(value) < 12 {
		return "", fmt.Errorf("invalid value length")
	}

	// Decode the encrypted data from base64
	decodedData, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	// Create a new AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a new AES-GCM cipher mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Get the IV from the decoded data
	iv := decodedData[:12]

	// Get the cipher value from the decoded data
	cipherValue := decodedData[12:]

	// Decrypt the value
	plainValue, err := aesgcm.Open(nil, iv, cipherValue, nil)
	if err != nil {
		return "", err
	}

	return string(plainValue), nil
}
