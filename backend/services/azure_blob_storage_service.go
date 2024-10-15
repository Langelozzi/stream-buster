package services

import (
	"io"

	"github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
)

type AzureBlobStorageService struct {
	dao interfaces.FileStorageDaoInterface
}

func NewAzureBlobStorageService(dao interfaces.FileStorageDaoInterface) *AzureBlobStorageService {
	return &AzureBlobStorageService{dao: dao}
}

func (s *AzureBlobStorageService) CreateFile(document *models.Document, file io.Reader) (string, error) {
	key := document.FilePath + "/" + document.FileName
	return s.dao.UploadFileDao(key, file)
}

func (s *AzureBlobStorageService) GetFile(document *models.Document) (io.Reader, error) {
	return s.dao.GetFileDao(document.FileStorageURL)
}
