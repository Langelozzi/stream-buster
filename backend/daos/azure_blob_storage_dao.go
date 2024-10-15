package daos

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/DOC-SWAP/Docswap-backend/utils"
)

type AzureBlobStorageDao struct {
	client *azblob.Client
	ctx    context.Context
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func generateSASToken(accountName, accountKey, containerName string) (string, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return "", err
	}

	permissions := sas.ContainerPermissions{
		Read:   true,
		Add:    true,
		Create: true,
		Write:  true,
		Delete: true,
		List:   true,
	}

	start := time.Now().UTC().Add(-10 * time.Minute)
	expiry := start.Add(1 * time.Hour)

	sasQueryParams, err := sas.BlobSignatureValues{
		Version:       sas.Version,
		Protocol:      sas.ProtocolHTTPS,
		StartTime:     start,
		ExpiryTime:    expiry,
		Permissions:   permissions.String(),
		ContainerName: containerName,
	}.SignWithSharedKey(credential)

	if err != nil {
		return "", err
	}

	return sasQueryParams.Encode(), nil
}

func NewAzureBlobStorageDao() *AzureBlobStorageDao {
	ctx := context.Background()
	accountName := utils.GetEnvVariable("AZURE_ACCOUNT_NAME")
	accountKey := utils.GetEnvVariable("AZURE_ACCOUNT_KEY")
	containerName := utils.GetEnvVariable("BLOB_CONTAINER")

	sasToken, err := generateSASToken(accountName, accountKey, containerName)
	handleError(err)
	url := "https://" + accountName + ".blob.core.windows.net/" + "?" + sasToken

	client, err := azblob.NewClientWithNoCredential(url, nil)
	handleError(err)

	return &AzureBlobStorageDao{
		client: client,
		ctx:    ctx,
	}
}

func getContainerAndBlobName(key string) (string, string, error) {
	parts := strings.SplitN(key, "/", 2) // Split into two parts at the first "/"

	err := error(nil)

	if len(parts) < 2 {
		err = errors.New("Invalid key format")
		return "", "", err
	}

	containerName := parts[0]
	blobName := parts[1]
	return containerName, blobName, err
}

func (dao *AzureBlobStorageDao) UploadFileDao(key string, file io.Reader) (string, error) {
	containerName, blobName, err := getContainerAndBlobName(key)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = dao.client.UploadBuffer(dao.ctx, containerName, blobName, data, &azblob.UploadBufferOptions{})
	return key, err
}

func (dao *AzureBlobStorageDao) CreateContainer(containerName string) error {
	_, err := dao.client.CreateContainer(dao.ctx, containerName, nil)
	return err
}

func (dao *AzureBlobStorageDao) GetFileDao(key string) (io.Reader, error) {
	containerName, blobName, err := getContainerAndBlobName(key)
	if err != nil {
		return nil, err
	}

	downloadResponse, err := dao.client.DownloadStream(dao.ctx, containerName, blobName, nil)
	if err != nil {
		return nil, err
	}

	// convert the download response to a reader
	return downloadResponse.Body, nil
}
