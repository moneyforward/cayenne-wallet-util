package sign

// MockSignature Mock Implementation of Signature
type MockSignature struct{}

// NewMockSignature Initialize MockSignature
func NewMockSignature() Signature {
	return &MockSignature{}
}

// Sign Sign Signature with Mock
func (s *MockSignature) Sign(message string) ([]byte, error) {
	return []byte(message), nil
}

// Verify Verify Signature with Mock
func (s *MockSignature) Verify(message string, sig []byte) (bool, error) {
	return true, nil
}
