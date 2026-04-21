package blob

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type BlobClient struct {
	client    *azblob.Client
	container string
	account   string
}

func NewBlobClient() (*BlobClient, error) {

	account := os.Getenv("AZURE_STORAGE_ACCOUNT")
	container := os.Getenv("AZURE_STORAGE_CONTAINER")

	if account == "" || container == "" {
		return nil, fmt.Errorf("missing blob env vars")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/", account)

	client, err := azblob.NewClient(url, cred, nil)
	if err != nil {
		return nil, err
	}

	return &BlobClient{
		client:    client,
		container: container,
		account:   account,
	}, nil
}

func (b *BlobClient) Upload(ctx context.Context, fileName string, data []byte) (string, error) {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var err error

	for i := 0; i < 3; i++ {
		_, err = b.client.UploadBuffer(ctx, b.container, fileName, data, nil)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.blob.core.windows.net/%s/%s",
		b.account,
		b.container,
		fileName,
	)

	return url, nil
}
