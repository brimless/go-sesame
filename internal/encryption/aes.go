package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

type AES struct {
	key [32]byte
}

func NewAES(masterPassword string) *AES {
	// since we're doing AES-256, we need to ensure that the key has 32 bytes (256 bits)
	// for simplicity we'll use SHA-256
	key := sha256.Sum256([]byte(masterPassword))
	return &AES{
		key: key,
	}
}

func (a *AES) Encrypt(plaintext string) (string, error) {
	plaintextBytes := []byte(plaintext)

	// generate cipher
	ciph, err := aes.NewCipher(a.key[:])
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// generate GCM with cipher
	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// populate nonce with random numbers
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// encrypt text
	encrypted := string(gcm.Seal(nonce, nonce, plaintextBytes, nil))

	return encrypted, nil
}

func (a *AES) Decrypt(encrypted string) (string, error) {
	encryptedBytes := []byte(encrypted)

	// generate cipher
	ciph, err := aes.NewCipher(a.key[:])
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// generate GCM with cipher
	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	// get nonce size from GCM
	nonceSize := gcm.NonceSize()
	if nonceSize > len(encryptedBytes) {
		return "", fmt.Errorf("Nonce size is bigger than encrypted input")
	}

	// nonce will be the first nonceSize bytes in the slice
	nonce := encryptedBytes[:nonceSize]

	// the ciphertext will be the remaining bytes
	ciphertext := encryptedBytes[nonceSize:]

	// get the decrypted text
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf(err.Error())
	}

	plaintext := string(plaintextBytes)

	return plaintext, nil
}
