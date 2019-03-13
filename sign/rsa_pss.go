package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
)

const (
	saltLength = 100
)

var (
	rsaHash          = sha256.New()
	rsaCryptoHash    = crypto.SHA256
	defaultPssOption = &rsa.PSSOptions{
		SaltLength: saltLength,
	}
)

// RSASignature RSA Implementation of Signature
type RSASignature struct {
	publicKey  *rsa.PublicKey
	privatekey *rsa.PrivateKey
}

// NewRSASignatureWithPublicKey Initialize RSASignature With PublicKey
func NewRSASignatureWithPublicKey(publicKey *rsa.PublicKey) Signature {
	return &RSASignature{publicKey: publicKey}
}

// NewRSASignatureWithPrivatekey Initialize RSASignature With PrivateKey
func NewRSASignatureWithPrivatekey(privateKey *rsa.PrivateKey) (Signature, error) {
	if privateKey == nil {
		return nil, errors.New("private key is nil")
	}
	return &RSASignature{
		privatekey: privateKey,
		publicKey:  &privateKey.PublicKey,
	}, nil
}

// Sign Sign Signature with RSAPSS
func (s *RSASignature) Sign(message string) ([]byte, error) {
	if s.privatekey == nil {
		return nil, errors.New("private key is nil")
	}
	byteMessage := []byte(message)
	hashed := rsaHash.Sum(byteMessage)
	return rsa.SignPSS(rand.Reader, s.privatekey, rsaCryptoHash, hashed, defaultPssOption)
}

// Verify Verify Signature with RSAPSS
func (s *RSASignature) Verify(message string, sig []byte) (bool, error) {
	if s.publicKey == nil {
		return false, errors.New("public key is nil")
	}
	byteMessage := []byte(message)
	hashed := rsaHash.Sum(byteMessage)
	if err := rsa.VerifyPSS(s.publicKey, rsaCryptoHash, hashed, sig, defaultPssOption); err != nil {
		return false, err
	}
	return true, nil
}
