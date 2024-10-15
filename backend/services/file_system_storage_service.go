package services

import (
	daoInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
	"io"
)

type FileSystemStorageService struct {
	fileSystemDao daoInterfaces.FileStorageDaoInterface
}

func NewFileStorageService(fileSystemDao daoInterfaces.FileStorageDaoInterface) *FileSystemStorageService {
	return &FileSystemStorageService{fileSystemDao: fileSystemDao}
}

func (service *FileSystemStorageService) CreateFile(document models.Document, file io.Reader) (string, error) {
	key := document.FileName
	return service.fileSystemDao.UploadFileDao(key, file)
}

func (service *FileSystemStorageService) GetFile(document models.Document) (io.Reader, error) {
	key := document.FileName
	return service.fileSystemDao.GetFileDao(key)
}
