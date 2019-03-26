package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// ReadRSAPrivateKeyFromBytes PrivateKeyの読み込み
func ReadRSAPrivateKeyFromBytes(bytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid private key data")
	}

	var key *rsa.PrivateKey
	var err error
	if block.Type == "RSA PRIVATE KEY" {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("invalid private key type : %s", block.Type)
	}

	key.Precompute()

	if err := key.Validate(); err != nil {
		return nil, err
	}

	return key, nil
}

// ReadRSAPrivateKeyFromBase64String base64StringからのPrivateKeyの読み込み
func ReadRSAPrivateKeyFromBase64String(base64String string) (*rsa.PrivateKey, error) {
	pemBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return ReadRSAPrivateKeyFromBytes(pemBytes)
}

// ReadRSAPublicKeyFromBytes PublicKeyの読み込み
func ReadRSAPublicKeyFromBytes(bytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(bytes)
	if block == nil {
		return nil, errors.New("invalid public key data")
	}
	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid public key type : %s", block.Type)
	}

	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// ReadRSAPublicKeyFromBase64String Base64StringからのPublicKeyの読み込み
func ReadRSAPublicKeyFromBase64String(base64String string) (*rsa.PublicKey, error) {
	pemBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}
	return ReadRSAPublicKeyFromBytes(pemBytes)
}
