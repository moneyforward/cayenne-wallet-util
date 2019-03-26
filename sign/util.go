package sign

import (
	"encoding/base64"
	"fmt"
)

const (
	signFormat = "%s_%s"
)

// SignWithID Sign With some identifier
func SignWithID(signature Signature, id, message string) ([]byte, error) {
	signInput := fmt.Sprintf(signFormat, id, message)
	return signature.Sign(signInput)
}

// SignBase64WithID is sign with some identifier and encode by base64 string
func SignBase64WithID(signature Signature, id, message string) (string, error) {
	signedBytes, err := SignWithID(signature, id, message)
	if err != nil {
		return "", err
	}
	base64Sig := base64.StdEncoding.EncodeToString(signedBytes)
	return base64Sig, nil
}

// VerifyWithID Verifyn With some identifier
func VerifyWithID(signature Signature, id, message string, sig []byte) (bool, error) {
	signInput := fmt.Sprintf(signFormat, id, message)
	return signature.Verify(signInput, sig)
}

// VerifyBase64yWithID is to verify decoded Base64 string with some identifier
func VerifyBase64yWithID(signature Signature, id, message, base64Sig string) (bool, error) {
	sig, err := base64.StdEncoding.DecodeString(base64Sig)
	if err != nil {
		return false, err
	}
	return VerifyWithID(signature, id, message, sig)
}
