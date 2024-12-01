package encryption

type Encryptor interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(encrypted string) (string, error)
}
