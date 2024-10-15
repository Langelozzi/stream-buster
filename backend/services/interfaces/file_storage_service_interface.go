package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"io"
)

// FileStorageServiceInterface defines the interface for file storage operations.
type FileStorageServiceInterface interface {
	CreateFile(document *models.Document, file io.Reader) (string, error)
	GetFile(document *models.Document) (io.Reader, error)
}
