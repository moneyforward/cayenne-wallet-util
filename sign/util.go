package sign

import "fmt"

const (
	signFormat = "%s_%s"
)

// SignWithID Sign With some identifier
func SignWithID(signature Signature, id, message string) ([]byte, error) {
	signInput := fmt.Sprintf(signFormat, id, message)
	return signature.Sign(signInput)
}

// VerifyWithID Verifyn With some identifier
func VerifyWithID(signature Signature, id, message string, sig []byte) (bool, error) {
	signInput := fmt.Sprintf(signFormat, id, message)
	return signature.Verify(signInput, sig)
}
