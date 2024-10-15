package daos

import (
	"io"
	"os"
)

type FileSystemDao struct{}

func NewFileSystemDao() *FileSystemDao {
	return &FileSystemDao{}
}

func (fs *FileSystemDao) UploadFileDao(key string, file io.Reader) (string, error) {
	fileName := key
	outFile, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {

		}
	}(outFile)

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return fileName, err
}

func (fs *FileSystemDao) GetFileDao(key string) (io.Reader, error) {
	// Open the file at the given path
	file, err := os.Open(key)
	if err != nil {
		return nil, err
	}
	return file, nil
}
