package storage

import (
	"context"

	gstorage "cloud.google.com/go/storage"
)

// GCSUploader Uploader for GCS
type GCSUploader struct {
	ctx        context.Context
	client     *gstorage.Client
	bucket     *gstorage.BucketHandle
	bucketName string
}

// NewGCSUploader Initialize GCSUploader
func NewGCSUploader(ctx context.Context, bucketName string) (Uploader, error) {
	u := &GCSUploader{}
	var err error
	u.client, err = gstorage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	u.bucketName = bucketName
	u.bucket = u.client.Bucket(bucketName)
	return u, nil
}

// Upload Upload object to GCS
func (u *GCSUploader) Upload(path string, object []byte) error {

	w := u.bucket.Object(path).NewWriter(u.ctx)
	defer w.Close()

	_, err := w.Write(object)
	return err
}
