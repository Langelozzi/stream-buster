package interfaces

import "io"

// FileStorageService defines the interface for file storage operations.
type FileStorageDaoInterface interface {
	UploadFileDao(key string, file io.Reader) (string, error)

	GetFileDao(key string) (io.Reader, error)
}
