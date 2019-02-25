package encryption

// Crypter Crypto Interface
type Crypter interface {
	Encrypt(src []byte) []byte
	EncryptBase64(plainText string) string
	Decrypt(src []byte) []byte
	DecryptBase64(base64String string) (string, error)
	GenerateToEnv(importFile string) (string, error) //TODO: this func should belong to env package
}
