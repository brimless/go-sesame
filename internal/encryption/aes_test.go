package encryption

import (
	"testing"
)

func TestNewAES(t *testing.T) {
	masterPassword := "test-master-password"
	aes := NewAES(masterPassword)

	if len(aes.key) != 32 {
		t.Fatalf("Expected key length to be 32 bytes, but got %d", len(aes.key))
	}
}

// test usual flow
func TestEncryptDecrypt(t *testing.T) {
	masterPassword := "test-master-password"
	aes := NewAES(masterPassword)

	plaintext := "This is the password to encrypt!"

	// encrypt
	encrypted, err := aes.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}
	if encrypted == plaintext {
		t.Fatalf("Encrypted text should not match plaintext")
	}

	// then decrypt
	decrypted, err := aes.Decrypt(encrypted)
	if err != nil {
		t.Fatalf("Decrypt failed: %v", err)
	}

	// if they don't match, well it didn't work, did it lol
	if decrypted != plaintext {
		t.Fatalf("Expected decrypted text to match plaintext. Got: %s, Expected: %s", decrypted, plaintext)
	}
}

// test flow when wrong master password is provided
func TestDecryptWithWrongKey(t *testing.T) {
	// first encrypt with the right password
	masterPassword := "test-master-password"
	aes := NewAES(masterPassword)

	plaintext := "This is the password to encrypt!"

	encrypted, err := aes.Encrypt(plaintext)
	if err != nil {
		t.Fatalf("Encrypt failed: %v", err)
	}

	// then try to decrypt with fake/imposter password
	wrongMasterPassword := "imposter-password"
	aesImposter := NewAES(wrongMasterPassword)

	decrypted, err := aesImposter.Decrypt(encrypted)
	if err == nil {
		t.Fatalf("Expected decryption to fail with the wrong key, but it succeeded")
	}

	// if the imposter key decrypts the password, then we're screwed
	if decrypted == plaintext {
		t.Fatalf("Decrypted text should not match plaintext when using the wrong key")
	}
}
