package storage

import (
	"fmt"
	"os"
)

// LocalUploader Uploader for local Storage
type LocalUploader struct{}

// NewLocalUploader Initialize LocalUploader
func NewLocalUploader() Uploader {
	return &LocalUploader{}
}

// Upload Upload object to local storage
func (u *LocalUploader) Upload(path string, object []byte) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintln(f, string(object)); err != nil {
		return err
	}
	return nil
}
