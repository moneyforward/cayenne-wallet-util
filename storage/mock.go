package storage

// MockUploader Uploader for Mock Storage
type MockUploader struct{}

// NewMockUploader Initialize MockUploader
func NewMockUploader() Uploader {
	return &MockUploader{}
}

// Upload Mock Implementation of Upload
func (u *MockUploader) Upload(path string, object []byte) error {
	return nil
}
