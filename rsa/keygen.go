package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

const (
	keySize = 2048
)

// GeneratePrivateKey Generate RSA PrivateKey
func GeneratePrivateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keySize)
}

// GeneratePublickey Generate RSA PublicKey from PrivateKey
func GeneratePublickey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	return &privateKey.PublicKey
}

// EncodePrivateKeyToPem Encode PrivateKey To Pem as []byte
func EncodePrivateKeyToPem(privateKey *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
}

// EncodePrivateKeyToPemBase64String Encode PrivateKey To Pem as base64 encoded string
func EncodePrivateKeyToPemBase64String(privateKey *rsa.PrivateKey) string {
	pemBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	return base64.StdEncoding.EncodeToString(pemBytes)
}

// EncodePublicKeyToPem Encode PublicKey To Pem as []byte
func EncodePublicKeyToPem(publicKey *rsa.PublicKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
}

// EncodePublicKeyToPemBase64String Encode PublicKey To Pem as base64 encoded string
func EncodePublicKeyToPemBase64String(publicKey *rsa.PublicKey) string {
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
	return base64.StdEncoding.EncodeToString(pemBytes)
}
