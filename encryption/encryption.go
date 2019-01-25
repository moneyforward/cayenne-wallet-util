package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mf-financial/wallet-util/env"
	"github.com/pkg/errors"
)

// Crypt is for cipher config data
type Crypt struct {
	cipher cipher.Block
	iv     []byte
}

var (
	cryptInfo  Crypt
	defaultKey string
	defaultIv  string
)

// GetDefaultKeyAndIv デバッグ利用、埋め込み変数の確認
func GetDefaultKeyAndIv() (string, string) {
	return defaultKey, defaultIv
}

// Creates a new encryption/decryption object
// with a given key of a given size
// (16, 24 or 32 for AES-128, AES-192 and AES-256 respectively,
// as per http://golang.org/pkg/crypto/aes/#NewCipher)
//
// The key will be padded to the given size if needed.
// An IV is created as a series of NULL bytes of necessary length
// when there is no iv string passed as 3rd value to function.

// NewCryptWithParam is to create crypt instance
// key size should be 16,24,32
// iv size should be 16
func NewCryptWithParam(key, iv string) (*Crypt, error) {
	if len(iv) != aes.BlockSize {
		return nil, errors.Errorf("iv size should be %d", aes.BlockSize)
	}

	padded := make([]byte, len(key))
	copy(padded, []byte(key))

	bIv := []byte(iv)
	block, err := aes.NewCipher(padded)
	if err != nil {
		return nil, err
	}

	cryptInfo = Crypt{block, bIv}

	return &cryptInfo, nil
}

// NewCryptWithEnv is setup with default settings.
func NewCryptWithEnv() (*Crypt, error) {
	key := os.Getenv("ENC_KEY")
	iv := os.Getenv("ENC_IV")

	if key == "" || iv == "" {
		return nil, errors.New("set Environment Variable: ENC_KEY, ENC_IV")
	}

	return NewCryptWithParam(key, iv)
}

// NewCrypt is setup with default settings.
func NewCrypt() (*Crypt, error) {

	if defaultKey == "" || defaultIv == "" {
		return nil, errors.New("set values for defaultKey, defaultIv when building")
	}

	return NewCryptWithParam(defaultKey, defaultIv)
}

// GetCrypt is to get crypt instance
func GetCrypt() *Crypt {
	return &cryptInfo
}

func (c *Crypt) padSlice(src []byte) []byte {
	// src must be a multiple of block size
	mult := int((len(src) / aes.BlockSize) + 1)
	leng := aes.BlockSize * mult

	srcPadded := make([]byte, leng)
	copy(srcPadded, src)
	return srcPadded
}

// Encrypt is encrypt a slice of bytes, producing a new, freshly allocated slice
// Source will be padded with null bytes if necessary
func (c *Crypt) Encrypt(src []byte) []byte {
	if len(src)%aes.BlockSize != 0 {
		src = c.padSlice(src)
	}
	dst := make([]byte, len(src))
	cipher.NewCBCEncrypter(c.cipher, c.iv).CryptBlocks(dst, src)
	return dst
}

// EncryptBase64 is encrypt and encode by base64 string
func (c *Crypt) EncryptBase64(plainText string) string {
	encryptedBytes := c.Encrypt([]byte(plainText))
	base64 := base64.StdEncoding.EncodeToString(encryptedBytes)
	return base64
}

// EncryptStream is to encrypt blocks from reader, write results into writer
func (c *Crypt) EncryptStream(reader io.Reader, writer io.Writer) error {
	for {
		buf := make([]byte, aes.BlockSize)
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.EOF {
				break
			} else if err == io.ErrUnexpectedEOF {
				// nothing
			} else {
				return err
			}
		}
		cipher.NewCBCEncrypter(c.cipher, c.iv).CryptBlocks(buf, buf)
		if _, err = writer.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

// Decrypt is to decrypt a slice of bytes, producing a new, freshly allocated slice
// Source will be padded with null bytes if necessary
func (c *Crypt) Decrypt(src []byte) []byte {
	if len(src)%aes.BlockSize != 0 {
		src = c.padSlice(src)
	}
	dst := make([]byte, len(src))
	cipher.NewCBCDecrypter(c.cipher, c.iv).CryptBlocks(dst, src)
	trimmed := bytes.Trim(dst, "\x00")
	return trimmed
}

// DecryptBase64 is to decrypt decoded Base64 string
func (c *Crypt) DecryptBase64(base64String string) (string, error) {
	unbase64, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return "", err
	}
	decryptedBytes := c.Decrypt(unbase64)
	return string(decryptedBytes[:]), nil
}

// DecryptStream is to decrypt blocks from reader, write results into writer
func (c *Crypt) DecryptStream(reader io.Reader, writer io.Writer) error {
	buf := make([]byte, aes.BlockSize)
	for {
		_, err := io.ReadFull(reader, buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		cipher.NewCBCDecrypter(c.cipher, c.iv).CryptBlocks(buf, buf)
		if _, err = writer.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

// GenerateToEnv is generate an encrypted .env file
func (c *Crypt) GenerateToEnv(importFile string) (string, error) {
	data, err := env.ImportData(importFile)
	if err != nil {
		fmt.Printf("failed to call env.ImportData(%s) error: %s\n", importFile, err)
	}

	envData := c.generateEnvData(data)
	return env.WriteEnvWithMultipleData(envData, true)
}

func (c *Crypt) generateEnvData(data []string) map[string][]string {
	var envData = map[string][]string{}

	for _, row := range data {

		rows := strings.Split(row, " ")
		envKey := rows[0]                        // envのキーを取得
		envValues := strings.Split(rows[1], ",") // envの値を取得（配列の場合はカンマ区切り）

		encryptedData := make([]string, 0, len(envValues))
		for _, envValue := range envValues {
			encryptedData = append(encryptedData, c.EncryptBase64(envValue))
		}
		envData[envKey] = encryptedData
	}
	return envData
}

// ShowAndExit main.goから呼び出される埋め込み確認用
func ShowAndExit() {
	defaultKey, defaultIv := GetDefaultKeyAndIv()
	fmt.Printf("defaultKey: %s, defaultIv: %s\n", defaultKey, defaultIv)

	os.Exit(0)
}
