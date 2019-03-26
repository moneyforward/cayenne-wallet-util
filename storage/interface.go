package storage

type Uploader interface {
	Upload(path string, object []byte) error
}
