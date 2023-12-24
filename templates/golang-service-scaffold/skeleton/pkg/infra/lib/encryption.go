package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

const PADDING_CHAR = '_'

var errCipherTextLength = errors.New("ciphertext too short")

func GetRandomBytes(size int) ([]byte, error) {
	key := make([]byte, size)

	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("error on GetRandomBytes: %w", err)
	}

	return key, nil
}

func Encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error on encrypt: %w", err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	initVector := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, initVector); err != nil {
		return nil, fmt.Errorf("error on encrypt: %w", err)
	}

	stream := cipher.NewCFBEncrypter(block, initVector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error on encrypt: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errCipherTextLength
	}

	initVector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, initVector)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

func EncodeString(input []byte) string {
	result := base64.URLEncoding.
		// WithPadding(PADDING_CHAR).
		EncodeToString(input)

	return result
}

func DecodeString(input string) ([]byte, error) {
	dst, err := base64.URLEncoding.
		// WithPadding(PADDING_CHAR).
		DecodeString(input)
	if err != nil {
		return nil, fmt.Errorf("Error decoding string: %w", err)
	}

	return dst, nil
}
