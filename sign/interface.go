package sign

// Signature Degital Signature Interface
type Signature interface {
	Sign(message string) ([]byte, error)
	// If true, Verify succeeded
	Verify(message string, sig []byte) (bool, error)
}
